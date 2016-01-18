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

/* import packages we're using */
import (
	"bytes"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"appengine"
	"appengine/datastore"
	//"appengine/user"

	//"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	newappengine "google.golang.org/appengine"
	//newurlfetch "google.golang.org/appengine/urlfetch"
)

const feedUrl = `https://www.google.com/m8/feeds/contacts/hugotest.org/full`

var exampleEntry = `<atom:entry xmlns:atom='http://www.w3.org/2005/Atom'
    xmlns:gd='http://schemas.google.com/g/2005'>
  <atom:category scheme='http://schemas.google.com/g/2005#kind'
    term='http://schemas.google.com/contact/2008#contact' />
  <gd:name>
     <gd:givenName>Testo 1</gd:givenName>
     <gd:familyName>Testo User 1 family name</gd:familyName>
     <gd:fullName>Testo User 1</gd:fullName>
  </gd:name>
  <atom:content type='text'>Notes</atom:content>
  <gd:email rel='http://schemas.google.com/g/2005#work'
    primary='true'
    address='testouser1@gmail.com' displayName='E. Bennet' />
  <gd:email rel='http://schemas.google.com/g/2005#home'
    address='testouser@example.org' />
  <gd:phoneNumber rel='http://schemas.google.com/g/2005#work'
    primary='true'>
    (206)555-1212
  </gd:phoneNumber>
  <gd:phoneNumber rel='http://schemas.google.com/g/2005#home'>
    (111)111-1111
  </gd:phoneNumber>
  <gd:im address='testouser@gmail.com'
    protocol='http://schemas.google.com/g/2005#GOOGLE_TALK'
    primary='true'
    rel='http://schemas.google.com/g/2005#home' />
  <gd:structuredPostalAddress
      rel='http://schemas.google.com/g/2005#work'
      primary='true'>
    <gd:city>testouser</gd:city>
    <gd:street>1600 Amphitheatre Pkwy</gd:street>
    <gd:region>CA</gd:region>
    <gd:postcode>94043</gd:postcode>
    <gd:country>United States</gd:country>
    <gd:formattedAddress>
      1600 Amphitheatre Pkwy Mountain View
    </gd:formattedAddress>
  </gd:structuredPostalAddress>
</atom:entry>`

// xmlns:gd=http://schemas.google.com/g/2005

type Feed struct {
	XMLName xml.Name  `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated,attr"`
	Author  Person    `xml:"author"`
	Entry   []Entry   `xml:"entry"`
}

type Entry struct {
	Title                   string                      `xml:"title"`
	Id                      string                      `xml:"id"`
	Link                    []Link                      `xml:"link"`
	Updated                 time.Time                   `xml:"updated" datastore:",noindex"`
	Author                  Person                      `xml:"author"`
	Summary                 Text                        `xml:"summary"`
	Name                    GDName                      `xml:"http://schemas.google.com/g/2005 name"`
	Im                      []GDIm                      `xml:"http://schemas.google.com/g/2005 im"`
	Email                   []GDEmail                   `xml:"http://schemas.google.com/g/2005 email"`
	PhoneNumber             []GDPhoneNumber             `xml:"http://schemas.google.com/g/2005 phoneNumber"`
	StructuredPostalAddress []GDStructuredPostalAddress `xml:"http://schemas.google.com/g/2005 formattedAddress"`
	ExtendedProperty        []GDExtendedProperty        `xml:"http://schemas.google.com/g/2005 extendedProperty"`
}

type EntryArb struct {
	Field []*Field
}

type Field struct {
	Name  string
	Value string
}

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Href string `xml:"href,attr"`
}

type Person struct {
	Name     string `xml:"name"`
	URI      string `xml:"uri"`
	Email    string `xml:"email"`
	InnerXML string `xml:",innerxml"`
}

type Text struct {
	Type string `xml:"type,attr,omitempty"`
	Body string `xml:",chardata"`
}

type GDName struct {
	FullName   string `xml:"http://schemas.google.com/g/2005 fullName"`
	GivenName  string `xml:"http://schemas.google.com/g/2005 givenName"`
	FamilyName string `xml:"http://schemas.google.com/g/2005 familyName"`
}

type GDIm struct {
	Address  string `xml:"address,attr"`
	Protocol string `xml:"protocol,attr"`
	Primary  bool   `xml:"primary,attr"`
}

type GDEmail struct {
	Address string `xml:"address,attr"`
	Primary bool   `xml:"primary,attr"`
}

type GDPhoneNumber struct {
	PhoneNumber string `xml:",chardata"`
	Primary     bool   `xml:"primary,attr"`
}

type GDStructuredPostalAddress struct {
	Primary          bool   `xml:"primary,attr"`
	City             string `xml:"city"`
	Street           string `xml:"street"`
	Region           string `xml:"region"`
	Postcode         string `xml:"postcode"`
	Country          string `xml:"country"`
	FormattedAddress string `xml:"formattedAddress"`
}

type GDExtendedProperty struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type ImportData struct {
	Data []byte
}

var (
	config = &oauth2.Config{
		ClientID:     `1078325845115-7lsb0oi0k539jrpqdpahun0v69684ceo.apps.googleusercontent.com`,
		ClientSecret: `I-0bww6IDmGk6IWLSQsq8v7T`,
		RedirectURL:  `http://hugotest.org/contacts/export`, //`http://hugotest.org/import/do`,
		Scopes:       []string{`http://www.google.com/m8/feeds/contacts/`},
		Endpoint:     google.Endpoint,
	}

	yeah = "yeah"
)

/* init - this is somewhat the start point */
func init() {
	http.HandleFunc("/import", handleImport)
	http.HandleFunc("/import/do", handleImportDo)
	//http.HandleFunc("/export", handleExport)
	http.HandleFunc("/export", handleContacts)
	http.HandleFunc("/contacts", handleContacts)
	http.HandleFunc("/contacts/export", handleContactsExport)
	http.HandleFunc("/", handleHome)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
  <head>
    <title>Shared Contacts</title>
    <style type="text/css">
a.button {
    -webkit-appearance: button;
    -moz-appearance: button;
    appearance: button;
    text-decoration: none;
    color: initial;
    padding: 8px;
    margin: 8px;
}
    </style>
  </head>
  <body>
    <a class="button" href="/contacts">Get Shared Contacts</a><br/><hr/>
    <a class="button" href="/import">Import</a><br/>
    <a class="button" href="/export">Export</a><br/>
  </body>
</html>
`)
}

func handleContacts(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	config.RedirectURL = `http://hugotest.org/contacts/export`
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

	w.Header().Set("Content-Type", "application/csv" /*"text/csv"*/)

	buf := loadFullFeed(ctx, r)

	ctx.Infof("%v", buf.String())

	writeCSV(ctx, w, buf.Bytes())
}

func handleImport(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	//u := user.Current(ctx)

	config.RedirectURL = `http://hugotest.org/import/do`
	url := config.AuthCodeURL(yeah)

	ctx.Infof("Auth: %v", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleImportDo(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	state := r.FormValue("state")
	if state != yeah {
		ctx.Errorf("invalid state '%v'", state)
		return
	}

	newctx := newappengine.NewContext(r)

	tok, err := config.Exchange(newctx /*oauth2.NoContext*/, r.FormValue("code"))
	if err != nil {
		ctx.Errorf("exchange error: %v", err)
		return
	}

	client := config.Client(newctx, tok)

	cr := csv.NewReader(strings.NewReader(sample))
	records, err := cr.ReadAll()
	if err != nil {
		ctx.Errorf("%v", err)
		return
	}

	buf := new(bytes.Buffer)

	names, rec := records[0], records[8]
	fmt.Fprintf(buf, `<atom:entry xmlns:atom='http://www.w3.org/2005/Atom' xmlns:gd='http://schemas.google.com/g/2005'>
<atom:category scheme='http://schemas.google.com/g/2005#kind' term='http://schemas.google.com/contact/2008#contact' />
<atom:content type='text'>Notes</atom:content>
`)

	numExtended, M := 0, 10
	for i, v := range rec {
		s := names[i]
		if s == "Name" {
			fmt.Fprintf(buf, "<gd:name><gd:fullName>%v</gd:fullName></gd:name>\n", v)
		} else if s == "E-mail Address" {
			fmt.Fprintf(buf, `<gd:email rel='http://schemas.google.com/g/2005#home' address='%v' />
`, v)
		} else if strings.HasPrefix(s, "E-mail ") && strings.HasSuffix(s, " Address") {
			var num uint
			fmt.Sscanf(s, "E-mail %v Address", &num)
			if numExtended < M && ((0 < num && num < 6) || s == "E-mail Address") {
				fmt.Fprintf(buf, `<gd:extendedProperty name="%v" value="%v" />`+"\n", s, v)
				numExtended++
			}
		} else if numExtended < M {
			fmt.Fprintf(buf, `<gd:extendedProperty name="%v" value="%v" />`+"\n", s, v)
			numExtended++
		}
	}
	fmt.Fprintf(buf, `</atom:entry>`)

	res, _ := client.Post(feedUrl, `application/atom+xml`, strings.NewReader(exampleEntry))

	fmt.Fprintf(w, `Result: %v<br/>`, res.Status)
	fmt.Fprintf(w, `<textarea>%v</textarea>`, buf.String())
}

func handleExport(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	w.Header().Set("Content-Type", "application/csv" /*"text/csv"*/)

	d := &ImportData{}
	k := datastore.NewKey(ctx, "ImportData", "ImportedFeed", 0, nil)
	if err := datastore.Get(ctx, k, d); err != nil {
		ctx.Errorf("get: %v", err)
	} else {
		writeCSV(ctx, w, d.Data)
	}
}

func writeCSV(ctx appengine.Context, w http.ResponseWriter, data []byte) {
	var feed Feed
	if err := xml.Unmarshal(data, &feed); err != nil {
		ctx.Errorf("unmarshal feed: %v", err)
		return
	}

	contacts, names := [][]string{}, []string{}

	upboundEmail, upboundPhone, upboundIm := 0, 0, 0
	propNames := []string{}
	names = append(names, "Id", "FullName", "GivenName", "FamilyName")
	for _, entry := range feed.Entry {
		if n := len(entry.Email); upboundEmail < n {
			upboundEmail = n
		}
		if n := len(entry.PhoneNumber); upboundPhone < n {
			upboundPhone = n
		}
		if n := len(entry.Im); upboundIm < n {
			upboundIm = n
		}
		for _, p := range entry.ExtendedProperty {
			existed := false
			for _, s := range propNames {
				if s == p.Name {
					existed = true
					break
				}
			}
			if !existed {
				propNames = append(propNames, p.Name)
			}
		}
	}
	for n := 1; n <= upboundEmail; n++ {
		names = append(names, fmt.Sprintf("Email %v", n))
	}
	for n := 1; n <= upboundPhone; n++ {
		names = append(names, fmt.Sprintf("Phone %v", n))
	}
	for n := 1; n <= upboundIm; n++ {
		names = append(names, fmt.Sprintf("Im %v", n))
	}
	propStart := len(names)
	sort.Strings(propNames[0:])
	names = append(names, propNames...)
	contacts = append(contacts, names)

	for _, entry := range feed.Entry {
		values := []string{
			entry.Id,
			entry.Name.FullName,
			entry.Name.GivenName,
			entry.Name.FamilyName,
		}
		for n, m := 0, len(entry.Email); n < upboundEmail; n++ {
			s := ""
			if n < m {
				s = entry.Email[n].Address
			}
			values = append(values, strings.TrimSpace(s))
		}
		for n, m := 0, len(entry.PhoneNumber); n < upboundPhone; n++ {
			s := ""
			if n < m {
				s = entry.PhoneNumber[n].PhoneNumber
			}
			values = append(values, strings.TrimSpace(s))
		}
		for n, m := 0, len(entry.Im); n < upboundIm; n++ {
			s := ""
			if n < m {
				s = entry.Im[n].Address
			}
			values = append(values, strings.TrimSpace(s))
		}
		for n, m := 0, len(propNames); n < m; n++ {
			values = append(values, "")
		}
		for n, m, x := 0, len(entry.ExtendedProperty), len(propNames); n < m; n++ {
			ep := &entry.ExtendedProperty[n]
			for i := 0; i < x; i++ {
				if ss := propNames[i]; ss == ep.Name {
					values[propStart+i] = strings.TrimSpace(ep.Value)
					break
				}
			}
		}
		contacts = append(contacts, values)
	}

	out := csv.NewWriter(w)
	out.WriteAll(contacts)
	if err := out.Error(); err != nil {
		ctx.Errorf("error writing csv:", err)
	}
}

func loadFullFeed(ctx appengine.Context, r *http.Request) (buf *bytes.Buffer) {
	newctx := newappengine.NewContext(r)

	tok, err := config.Exchange(newctx /*oauth2.NoContext*/, r.FormValue("code"))
	if err != nil {
		ctx.Errorf("exchange error: %v", err)
		return
	}

	ctx.Infof("tok: %v", tok)

	client := config.Client(newctx, tok)

	res, err := client.Get(feedUrl)
	if err != nil {
		ctx.Errorf("get: %v", err)
		return
	}

	defer res.Body.Close()

	buf = new(bytes.Buffer)
	io.Copy(buf, res.Body)
	return
}

// https://developers.google.com/admin-sdk/domain-shared-contacts
/*
<feed xmlns='http://www.w3.org/2005/Atom'
    xmlns:openSearch='http://a9.com/-/spec/opensearchrss/1.0/'
    xmlns:gd='http://schemas.google.com/g/2005'
    xmlns:gContact='http://schemas.google.com/contact/2008'
    xmlns:batch='http://schemas.google.com/gdata/batch'>
  <id>https://www.google.com/m8/feeds/contacts/example.com/base</id>
  <updated>2008-03-05T12:36:38.836Z</updated>
  <category scheme='http://schemas.google.com/g/2005#kind'
    term='http://schemas.google.com/contact/2008#contact' />
  <title type='text'>example.com's Contacts</title>
  <link rel='http://schemas.google.com/g/2005#feed'
    type='application/atom+xml'
    href='https://www.google.com/m8/feeds/contacts/example.com/full' />
  <link rel='http://schemas.google.com/g/2005#post'
    type='application/atom+xml'
    href='https://www.google.com/m8/feeds/contacts/example.com/full' />
  <link rel='http://schemas.google.com/g/2005#batch'
    type='application/atom+xml'
    href='https://www.google.com/m8/feeds/contacts/example.com/full/batch' />
  <link rel='self' type='application/atom+xml'
    href='https://www.google.com/m8/feeds/contacts/example.com/full?max-results=25' />
  <author>
    <name>example.com</name>
    <email>example.com</email>
  </author>
  <generator version='1.0' uri='https://www.google.com/m8/feeds/contacts'>
    Contacts
  </generator>
  <openSearch:totalResults>1</openSearch:totalResults>
  <openSearch:startIndex>1</openSearch:startIndex>
  <openSearch:itemsPerPage>25</openSearch:itemsPerPage>
  <entry>
    <id>
      https://www.google.com/m8/feeds/contacts/example.com/base/c9012de
    </id>
    <updated>2008-03-05T12:36:38.835Z</updated>
    <category scheme='http://schemas.google.com/g/2005#kind'
      term='http://schemas.google.com/contact/2008#contact' />
    <title type='text'>Fitzgerald</title>
    <gd:name>
      <gd:fullName>Fitzgerald</gd:fullName>
    </gd:name>
    <link rel="http://schemas.google.com/contacts/2008/rel#photo" type="image/*"
      href="http://google.com/m8/feeds/photos/media/example.com/c9012de"/>
    <link rel='self' type='application/atom+xml'
      href='https://www.google.com/m8/feeds/contacts/example.com/full/c9012de' />
    <link rel='edit' type='application/atom+xml'
      href='https://www.google.com/m8/feeds/contacts/example.com/full/c9012de/1204720598835000' />
    <gd:phoneNumber rel='http://schemas.google.com/g/2005#home'
      primary='true'>
      456
    </gd:phoneNumber>
    <gd:extendedProperty name="pet" value="hamster" />
  </entry>
</feed>
*/
