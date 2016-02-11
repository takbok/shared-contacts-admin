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
}

const deleteEntryTemplate = `<entry gd:etag='%s'> <batch:operation type='delete'/> <id>%s</id> </entry>`
