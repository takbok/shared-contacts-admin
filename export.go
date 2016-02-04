package demo

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/datastore"
)

func init() {
	http.HandleFunc("/export", handleContacts)

	http.HandleFunc("/contacts", handleContacts)
	http.HandleFunc("/contacts/export", handleContactsExport)
}

func handleContacts(w http.ResponseWriter, r *http.Request) {
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
