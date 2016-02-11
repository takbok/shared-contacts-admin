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
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"

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

	var contactsXml Feed
	if err := xml.Unmarshal(contactsFeed.Bytes(), &contactsXml); err != nil {
		ctx.Errorf("unmarshal feed: %v", err)
		return
	}

	var buffer bytes.Buffer

	for _, entry := range contactsXml.Entry {
		for _, link := range entry.Link {
			if link.Rel == `edit` {
				buffer.WriteString(fmt.Sprintf(deleteEntryTemplate, entry.ETag, link.Href))
			}
		}
	}

	batchData := fmt.Sprintf(batchFeedTemplate, buffer.String())

	res, _ := client.Post(getContactsBatchUrl(contactsXml.Link), `application/atom+xml`, strings.NewReader(batchData))
	defer res.Body.Close()

	fmt.Fprintf(w, "Result: %v<br/>", res.Status)

	// TODO Add a more verbose response to the resultsof whether the contacts are deleted or not
}

const deleteEntryTemplate = `<entry gd:etag='%s'> <batch:operation type='delete'/> <id>%s</id> </entry>`
