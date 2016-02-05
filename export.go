package demo

import (
	"fmt"
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

func isUrlOnGoogleApp(writer http.ResponseWriter, request *http.Request) bool {
	u, err := url.Parse(request.FormValue("url"))
	if err != nil {
		return false
	}

	p := strings.Split(u.Host, ".")
	p = p[len(p)-2 : len(p)]

	ctx := newappengine.NewContext(request)
	client := urlfetch.Client(ctx)

	uri := fmt.Sprintf("https://www.google.com/a/%s/ServiceLogin", strings.Join(p, "."))
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
	if r.Method == "POST" && !isUrlOnGoogleApp(w, r) {
		http.Redirect(w, r, "/?error=notOnGoogleApps", http.StatusTemporaryRedirect)
		return
	}

	ctx := appengine.NewContext(r)
	config.RedirectURL = fmt.Sprintf(`http://www.cloudtest1.com/contacts/export`, r.Host)
	url := config.AuthCodeURL(yeah)
	ctx.Infof("Auth: %v", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleContactsExport(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	state := r.FormValue("state")
	if state != yeah {
		ctx.Errorf("invalid state '%v'", state)
		return
	}

	w.Header().Set(`Content-Type`, `application/csv`)
	w.Header().Set(`Content-Disposition`, `attachment; filename="export.csv"`)

	buf := loadFullFeed(ctx, r)

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
