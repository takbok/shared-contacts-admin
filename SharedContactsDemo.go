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
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"appengine"

	newappengine "google.golang.org/appengine"
)

const feedUrl = `https://www.google.com/m8/feeds/contacts/%s/full?v=3.0`

var inp_file multipart.File

var exampleEntry = `
  <atom:entry xmlns:atom='http://www.w3.org/2005/Atom'
    xmlns:gd='http://schemas.google.com/g/2005'>
  <atom:category scheme='http://schemas.google.com/g/2005#kind'
    term='http://schemas.google.com/contact/2008#contact' />
  <gd:name>
     <gd:givenName>My External Contact 1</gd:givenName>
     <gd:familyName>My External Contact 1 family name</gd:familyName>
     <gd:fullName>My External Contact 1</gd:fullName>
  </gd:name>
  <atom:content type='text'>Notes</atom:content>
  <gd:email rel='http://schemas.google.com/g/2005#work'
    primary='true'
    address='myexternalcontact1@testmail.com' displayName='E. Bennet' />
  <gd:email rel='http://schemas.google.com/g/2005#home'
    address='myexternalcontact1@testmail.com' />
  <gd:phoneNumber rel='http://schemas.google.com/g/2005#work'
    primary='true'>
    (206)555-1212
  </gd:phoneNumber>
  <gd:phoneNumber rel='http://schemas.google.com/g/2005#home'>
    (111)111-1111
  </gd:phoneNumber>
  <gd:im address='myexternalcontact1@testmail.com'
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
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	ETag    string    `xml:"http://schemas.google.com/g/2005 etag,attr"`
	Link    []Link    `xml:"link"`
	Content string    `xml:"content"`
	Updated time.Time `xml:"updated" datastore:",noindex"`
	//	Author                  Person                      `xml:"author"`
	//	Summary                 Text                        `xml:"summary"`
	Name                    GDName                      `xml:"http://schemas.google.com/g/2005 name"`
	Im                      []GDIm                      `xml:"http://schemas.google.com/g/2005 im"`
	Email                   []GDEmail                   `xml:"http://schemas.google.com/g/2005 email"`
	PhoneNumber             []GDPhoneNumber             `xml:"http://schemas.google.com/g/2005 phoneNumber"`
	Organization            GDOrganization              `xml:"http://schemas.google.com/g/2005 organization"`
	StructuredPostalAddress []GDStructuredPostalAddress `xml:"http://schemas.google.com/g/2005 structuredPostalAddress"`
	ExtendedProperty        []GDExtendedProperty        `xml:"http://schemas.google.com/g/2005 extendedProperty"`
	ContactUDField          []GContactUDField           `xml:"http://schemas.google.com/contact/2008 userDefinedField"`
	ContactWebsite          []GContactWebsite           `xml:"http://schemas.google.com/contact/2008 website"`
	Birthday                GContactBirthday            `xml:"http://schemas.google.com/contact/2008 birthday"`
	Nickname                string                      `xml:"http://schemas.google.com/contact/2008 nickname"`
	ExternalId              []GContactExternalId        `xml:"http://schemas.google.com/contact/2008 externalId"`
	Occupation              string                      `xml:"http://schemas.google.com/contact/2008 occupation"`
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

type GContactBirthday struct {
	When string `xml:"when,attr"`
}

type GContactExternalId struct {
	Label string `xml:"label,attr"`
	Rel   string `xml:"rel,attr"`
	Value string `xml:"value,attr"`
}

type Text struct {
	Type string `xml:"type,attr,omitempty"`
	Body string `xml:",chardata"`
}

type GDName struct {
	FullName       string `xml:"http://schemas.google.com/g/2005 fullName"`
	GivenName      string `xml:"http://schemas.google.com/g/2005 givenName"`
	FamilyName     string `xml:"http://schemas.google.com/g/2005 familyName"`
	AdditionalName string `xml:"http://schemas.google.com/g/2005 additionalName"`
	NamePrefix     string `xml:"http://schemas.google.com/g/2005 namePrefix"`
	NameSuffix     string `xml:"http://schemas.google.com/g/2005 nameSuffix"`
}

type GDOrganization struct {
	Label             string `xml:"label,attr"`
	Rel               string `xml:"rel,attr"`
	Primary           bool   `xml:"primary,attr"`
	OrgDepartment     string `xml:"http://schemas.google.com/g/2005 orgDepartment"`
	OrgJobDescription string `xml:"http://schemas.google.com/g/2005 orgJobDescription"`
	OrgName           string `xml:"http://schemas.google.com/g/2005 orgName"`
	OrgSymbol         string `xml:"http://schemas.google.com/g/2005 orgSymbol"`
	OrgTitle          string `xml:"http://schemas.google.com/g/2005 orgTitle"`
	Where             string `xml:"http://schemas.google.com/g/2005 where"`
}

type GDIm struct {
	Address  string `xml:"address,attr"`
	Protocol string `xml:"protocol,attr"`
	Primary  bool   `xml:"primary,attr"`
	Label    string `xml:"label,attr"`
	Rel      string `xml:"rel,attr"`
}

type GDEmail struct {
	Address     string `xml:"address,attr"`
	Primary     bool   `xml:"primary,attr"`
	Label       string `xml:"label,attr"`
	Rel         string `xml:"rel,attr"`
	DisplayName string `xml:"displayName,attr"`
}

type GDPhoneNumber struct {
	PhoneNumber string `xml:",chardata"`
	Primary     bool   `xml:"primary,attr"`
	Label       string `xml:"label,attr"`
	Rel         string `xml:"rel,attr"`
	Uri         string `xml:"uri,attr"`
}

type GDStructuredPostalAddress struct {
	Primary          bool   `xml:"primary,attr"`
	City             string `xml:"city"`
	Street           string `xml:"street"`
	Region           string `xml:"region"`
	Postcode         string `xml:"postcode"`
	Country          string `xml:"country"`
	FormattedAddress string `xml:"formattedAddress"`
	Label            string `xml:"label,attr"`
	Rel              string `xml:"rel,attr"`
	MailClass        string `xml:"mailClass,attr"`
	Usage            string `xml:"usage,attr"`
}

type GDExtendedProperty struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type GContactUDField struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

type GContactWebsite struct {
	Webref  string `xml:"href,attr"`
	Label   string `xml:"label,attr"`
	Primary bool   `xml:"primary,attr"`
	Rel     string `xml:"rel,attr"`
}
type ImportData struct {
	Data []byte
}

func writeCSV(ctx appengine.Context, w http.ResponseWriter, data []byte) {
	var feed Feed
	if err := xml.Unmarshal(data, &feed); err != nil {
		ctx.Errorf("unmarshal feed: %v", err)
		return
	}
	buildColumnMap()

	contacts := [][]string{}

	contacts = append(contacts, column_names)

	for _, entry := range feed.Entry {
		values := []string{}
		for n := 0; n < NUM_COLUMNS; n++ {
			values = append(values, "")
		}
		values[ACTION_COL_IDX] = "update"
		str := entry.Id
		str = strings.Replace(str, "http://www.google.com/m8/feeds/contacts/", "", -1)
		str = strings.Replace(str, "base/", "", -1)
		strs := strings.Split(str, "/")
		values[DOMAIN_COL_IDX] = strs[0]
		values[ID_COL_IDX] = strs[1]
		values[NAME_COL_IDX] = entry.Name.FullName
		values[NAMELOWER_COL_IDX] = strings.ToLower(entry.Name.FullName)

		values[COMPANY_COL_IDX] = entry.Organization.OrgName
		values[JOBTITLE_COL_IDX] = entry.Organization.OrgTitle
		values[DEPARTMENT_COL_IDX] = entry.Organization.OrgDepartment
		values[JOBDESCRIPTION_COL_IDX] = entry.Organization.OrgJobDescription
		values[NOTES_COL_IDX] = entry.Content
		values[CONTACTTYPE_COL_IDX] = "External"
		values[APIID_COL_IDX] = entry.Id

		numEmails := len(entry.Email)
		for n := 0; n < numEmails; n++ {
			if entry.Email[n].Label != "" {
				colIdx, ok := column_name_map[entry.Email[n].Label]
				if ok {
					values[colIdx] = entry.Email[n].Address
				}
			} else if entry.Email[n].Rel == "http://schemas.google.com/g/2005#work" {
				colIdx, ok := column_name_map["E-mail Address"]
				if ok {
					values[colIdx] = entry.Email[n].Address
				}
			} else {
				colName := fmt.Sprintf("E-mail %v Address", n+1)
				colIdx, ok := column_name_map[colName]
				if ok {
					values[colIdx] = entry.Email[n].Address
				}
			}
		}

		numIms := len(entry.Im)
		imColnames := [2]string{"IM", "IM Rel"}
		for n := 0; n < numIms; n++ {
			if entry.Im[n].Label != "" {
				colIdx, ok := column_name_map[entry.Im[n].Label]
				if ok {
					values[colIdx] = entry.Im[n].Address
				}
			} else if n < 2 {
				colIdx, ok := column_name_map[imColnames[n]]
				if ok {
					values[colIdx] = entry.Im[n].Address
				}
			}
		}

		numPhNos := len(entry.PhoneNumber)
		for n := 0; n < numPhNos; n++ {
			if entry.PhoneNumber[n].Label != "" {
				colIdx, ok := column_name_map[entry.PhoneNumber[n].Label]
				if ok {
					values[colIdx] = entry.PhoneNumber[n].PhoneNumber
				}
			} else {
				phStr := strings.Replace(entry.PhoneNumber[n].Rel, "http://schemas.google.com/g/2005#", "", -1)
				var colName string
				switch phStr {
				case "work":
					colName = "Business Phone"
				case "work_fax":
					colName = "Business Fax"
				case "mobile":
					colName = "Mobile Phone"
				case "home":
					colName = "Home Phone"
				case "home_fax":
					colName = "Home Fax"
				case "other":
					colName = "Other Phone"
				case "pager":
					colName = "Pager"
				}
				colIdx, ok := column_name_map[colName]
				if ok {
					values[colIdx] = entry.PhoneNumber[n].PhoneNumber
				}
			}
		}

		numPostAdds := len(entry.StructuredPostalAddress)
		for n := 0; n < numPostAdds; n++ {
			if entry.StructuredPostalAddress[n].Label != "" {
				colIdx, ok := column_name_map[entry.StructuredPostalAddress[n].Label]
				if ok {
					values[colIdx] = entry.StructuredPostalAddress[n].FormattedAddress
				}
			} else {
				addStr := strings.Replace(entry.StructuredPostalAddress[n].Rel, "http://schemas.google.com/g/2005#", "", -1)
				var colName string
				switch addStr {
				case "work":
					colName = "Business Address"
				case "home":
					colName = "Home Address"
				case "other":
					colName = "Other Address"
				}
				colIdx, ok := column_name_map[colName]
				if ok {
					values[colIdx] = entry.StructuredPostalAddress[n].FormattedAddress
				}
			}
		}

		numExtProps := len(entry.ExtendedProperty)
		for n := 0; n < numExtProps; n++ {
			if entry.ExtendedProperty[n].Name != "" {
				colIdx, ok := column_name_map[entry.ExtendedProperty[n].Name]
				if ok {
					values[colIdx] = entry.ExtendedProperty[n].Value
				}
			}
		}

		numUDs := len(entry.ContactUDField)
		for n := 0; n < numUDs; n++ {
			colName1 := fmt.Sprintf("Custom Key%v", n+1)
			colName2 := fmt.Sprintf("Custom Value%v", n+1)
			colIdx, ok := column_name_map[colName1]
			if ok {
				values[colIdx] = entry.ContactUDField[n].Key
			}
			colIdx, ok = column_name_map[colName2]
			if ok {
				values[colIdx] = entry.ContactUDField[n].Value
			}
		}

		numWebs := len(entry.ContactWebsite)
		for n := 0; n < numWebs; n++ {
			var colName string
			webStr := entry.ContactWebsite[n].Rel
			switch webStr {
			case "home-page":
				colName = "Website Home-Page"
			case "blog":
				colName = "Website Blog"
			case "profile":
				colName = "Website Profile"
			case "home":
				colName = "Website Home"
			case "work":
				colName = "Website Work"
			case "ftp":
				colName = "Website FTP"
			}
			colIdx, ok := column_name_map[colName]
			if ok {
				values[colIdx] = entry.ContactWebsite[n].Webref
			}
		}
		values[BIRTHDAY_COL_IDX] = entry.Birthday.When
		values[NICKNAME_COL_IDX] = entry.Nickname
		if len(entry.ExternalId) > 0 {
			values[EXTERNALID_COL_IDX] = entry.ExternalId[0].Value
		}
		values[OCCUPATION_COL_IDX] = entry.Occupation

		contacts = append(contacts, values)
	}

	out := csv.NewWriter(w)
	out.WriteAll(contacts)
	if err := out.Error(); err != nil {
		ctx.Errorf("error writing csv:", err)
	}
}

func loadFullFeed(domain string, ctx appengine.Context, r *http.Request) (buf *bytes.Buffer) {
	newctx := newappengine.NewContext(r)

	tok, err := config.Exchange(newctx, r.FormValue("code"))
	if err != nil {
		ctx.Errorf("exchange error: %v", err)
		return
	}

	ctx.Infof("tok: %v", tok)

	client := config.Client(newctx, tok)
	expFeedUrl := fmt.Sprintf(feedUrl, domain)
	expFeedUrl = expFeedUrl + "&max-results=50000"
	res, err := client.Get(expFeedUrl)
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
