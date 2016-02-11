package demo

import (
	"net/http"

	"appengine"

	newappengine "google.golang.org/appengine"
)

func getOAuthClient(context appengine.Context, r *http.Request) *http.Client {
	newctx := newappengine.NewContext(r)

	tok, err := config.Exchange(newctx, r.FormValue("code"))
	if err != nil {
		context.Errorf("exchange error: %v", err)
		return nil
	}

	context.Infof("tok: %v", tok)

	return config.Client(newctx, tok)
}
