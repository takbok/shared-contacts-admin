package demo

import (
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	newappengine "google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"appengine"
	"appengine/datastore"
)

func init() {
	http.HandleFunc("/export", handleContacts)

	http.HandleFunc("/contacts", handleContacts)
	http.HandleFunc("/contacts/export", handleContactsExport)
}

func isUrlOnGoogleApp(writer http.ResponseWriter, request *http.Request, url string) bool {
	ctx := newappengine.NewContext(request)
	client := urlfetch.Client(ctx)

	uri := fmt.Sprintf("https://www.google.com/a/%s/ServiceLogin", url)
	resp, err := client.Get(uri)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return false
	}

	body, err := ioutil.ReadAll(resp.Body)
	responseBody := string(body[:])

	// Check if the Google Apps login URL is valid
	return strings.Contains(responseBody, "https://www.google.com/accounts/AccountChooser")
}

func handleContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/?error=noDirectAccess", http.StatusTemporaryRedirect)
		return
	}

	url, err := getProperDomainNameFromUrl(r.FormValue("url"))
	if err != nil {
		http.Redirect(w, r, "/?error=badUrl", http.StatusTemporaryRedirect)
		return
	}

	if !isUrlOnGoogleApp(w, r, url) {
		http.Redirect(w, r, "/?error=notOnGoogleApps", http.StatusTemporaryRedirect)
		return
	}

	ctx := appengine.NewContext(r)
	config.RedirectURL = fmt.Sprintf(`http://%s/contacts/export`, r.Host)

	x := AppState{url}
	url = config.AuthCodeURL(x.encodeState())
	ctx.Infof("Auth: %v", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleContactsExport(w http.ResponseWriter, r *http.Request) {
	y := r.FormValue("state")

	state := new(AppState)
	state.decodeState(y)

	w.Header().Set(`Content-Type`, `application/csv`)
	w.Header().Set(`Content-Disposition`, `attachment; filename="`+state.Domain+`-contacts-export.csv"`)

	ctx := appengine.NewContext(r)
	buf := loadFullFeed(state.Domain, ctx, r)

	ctx.Infof("%v", buf.String())
	writeCSV(ctx, w, buf.Bytes())
}

func handleExport(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	d := &ImportData{}
	k := datastore.NewKey(ctx, "ImportData", "ImportedFeed", 0, nil)

	w.Header().Set("Content-Type", "application/csv")
	if err := datastore.Get(ctx, k, d); err != nil {
		ctx.Errorf("get: %v", err)
	} else {
		writeCSV(ctx, w, d.Data)
	}
}

func getProperDomainNameFromUrl(u string) (string, error) {
	uri, err := url.Parse(u)
	if err != nil {
		return ``, fmt.Errorf("The URL seems to be invalid")
	}

	p := strings.Split(uri.Host, ".")

	if len(p) < 2 {
		return ``, fmt.Errorf("The URL seems to be invalid")
	}

	p = p[len(p)-2 : len(p)]

	return strings.Join(p, "."), nil
}
