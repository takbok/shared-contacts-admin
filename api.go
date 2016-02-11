package demo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"appengine"

	newappengine "google.golang.org/appengine"
)

const contactsFeedUrl = `https://www.google.com/m8/feeds/contacts/%s/%s?v=3.0&start-index=%d&max-results=%d&alt=%s`

func getContactsFeedUrl(feedtype, domain string, full bool, start, number int) string {
	var projection = `full`
	if !full {
		projection = `thin`
	}

	if feedtype != `rss` && feedtype != `json` {
		feedtype = `atom`
	}

	return fmt.Sprintf(contactsFeedUrl, domain, projection, start, number, feedtype)
}

func loadAllContacts(domain string, client *http.Client, context appengine.Context) (buf *bytes.Buffer) {
	res, err := client.Get(getContactsFeedUrl(`atom`, domain, false, 1, 100))
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
