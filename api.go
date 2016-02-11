package demo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"appengine"

	newappengine "google.golang.org/appengine"
)

func loadAllContacts(domain string, client *http.Client, context appengine.Context) (buf *bytes.Buffer) {
	res, err := client.Get(fmt.Sprintf(`https://www.google.com/m8/feeds/contacts/%s/full?v=3.0&max-results=100`, domain))
	if err != nil {
		context.Errorf("get: %v", err)
		return
	}

	defer res.Body.Close()

	buf = new(bytes.Buffer)
	io.Copy(buf, res.Body)
	return
}

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
