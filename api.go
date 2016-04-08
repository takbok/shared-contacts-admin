// Copyright 2016 Takbok, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func getContactsBatchUrl(links []Link) string {
	for _, link := range links {
		if link.Rel == `http://schemas.google.com/g/2005#batch` {
			return link.Href
		}
	}

	return ``
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

const batchFeedTemplate = `<?xml version='1.0' encoding='UTF-8'?>
<feed xmlns='http://www.w3.org/2005/Atom'
      xmlns:gContact='http://schemas.google.com/contact/2008'
      xmlns:gd='http://schemas.google.com/g/2005'
      xmlns:batch='http://schemas.google.com/gdata/batch'>
  <category scheme='http://schemas.google.com/g/2005#kind'
      term='http://schemas.google.com/g/2008#contact' />
  %s
</feed>`
