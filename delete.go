package demo

import (
	"fmt"
	"net/http"

	"appengine"
)

func init() {
	http.HandleFunc("/contacts/delete", deleteAllContacts)
}

func initiateContactsDeletion(w http.ResponseWriter, r *http.Request, url string) {
	ctx := appengine.NewContext(r)
	config.RedirectURL = fmt.Sprintf(`http://%s/contacts/delete`, r.Host)

	x := AppState{url}
	url = config.AuthCodeURL(x.encodeState())
	ctx.Infof("Auth: %v", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func deleteAllContacts(w http.ResponseWriter, r *http.Request) {
	y := r.FormValue("state")

	state := new(AppState)
	state.decodeState(y)

	ctx := appengine.NewContext(r)
	client := getOAuthClient(ctx, r)

	contactsFeed := loadAllContacts(state.Domain, client, ctx)
}
