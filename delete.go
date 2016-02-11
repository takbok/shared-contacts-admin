package demo

import (
	"fmt"
	"net/http"

	"appengine"
)

func initiateContactsDeletion(w http.ResponseWriter, r *http.Request, url string) {
	ctx := appengine.NewContext(r)
	config.RedirectURL = fmt.Sprintf(`http://%s/contacts/delete`, r.Host)

	x := AppState{url}
	url = config.AuthCodeURL(x.encodeState())
	ctx.Infof("Auth: %v", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
