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
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"

	"appengine"

	newappengine "google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/import", handleImport)
	http.HandleFunc("/import/do", handleImportDo)
}

func handleImport(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/?error=noDirectAccess", http.StatusTemporaryRedirect)
		return
	}

	url, err := getProperDomainNameFromUrl(r.FormValue("url"))
	if err != nil {
		http.Redirect(w, r, "/?error=badUrl", http.StatusTemporaryRedirect)
		return
	}

	if !isUrlOnGoogleApp(w, r, url) {
		http.Redirect(w, r, "/?error=notOnGoogleApps", http.StatusTemporaryRedirect)
		return
	}

	r.ParseMultipartForm(32 << 20)
	inp_file, _, err = r.FormFile("inputfile")
	if err != nil {
		return
	}
	defer inp_file.Close()

	x := AppState{url}
	ctx := appengine.NewContext(r)
	config.RedirectURL = fmt.Sprintf(`http://%s/import/do`, r.Host)

	url = config.AuthCodeURL(x.encodeState())
	ctx.Infof("Auth: %v", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleImportDo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IMPORT:<br>")
	ictr := 0
	y := r.FormValue("state")

	state := new(AppState)
	state.decodeState(y)

	ctx := appengine.NewContext(r)
	newctx := newappengine.NewContext(r)

	tok, err := config.Exchange(newctx, r.FormValue("code"))
	if err != nil {
		ctx.Errorf("exchange error: %v", err)
		return
	}

	client := config.Client(newctx, tok)

	cr := csv.NewReader(bufio.NewReader(inp_file))
	records, err := cr.ReadAll()
	if err != nil {
		ctx.Errorf("%v", err)
		return
	}

	names := records[0]
	datalen := len(records)

	for i := 1; i < datalen; i++ {
		rec := records[i]
		buf := new(bytes.Buffer)
		fmt.Fprintf(buf, `<atom:entry xmlns:atom='http://www.w3.org/2005/Atom'
    	xmlns:gd='http://schemas.google.com/g/2005' 
		xmlns:gContact='http://schemas.google.com/contact/2008'>
  		<atom:category scheme='http://schemas.google.com/g/2005#kind'
    	term='http://schemas.google.com/contact/2008#contact' />`)
		fmt.Fprintf(buf, "\n")

		var nameBuf, emailBuf, imBuf, orgBuf, phoneBuf, extendedBuf, postalAdress, webBuf, notesBuf string
		var birthBuf, nickBuf, externalIdBuf, occBuf string
		numExtendedProp, maxExtendedProp := 0, 10
		udKeys := []string{}
		udValues := []string{}

		orgBuf = `<gd:organization label="Company" primary="true">` + "\n"
		for j, s := range names {
			if strings.Trim(rec[j], " ") == "" {
				continue
			}
			//skip unmatched entries
			if s == "Action" {
				continue
			}
			if s == "ID" {
				continue
			}
			if s == "contactType" {
				continue
			}
			if s == "domain" {
				continue
			}
			if s == "apiId" {
				continue
			}
			if s == "NameLower" {
				continue
			}
			if s == "E-mail Address" {
				emailBuf += fmt.Sprintf(`<gd:email rel="http://schemas.google.com/g/2005#work" address="%v"/>`+"\n", rec[j])
				continue
			}
			if strings.Contains(s, "E-mail") {
				emailBuf += fmt.Sprintf(`<gd:email rel="http://schemas.google.com/g/2005#other" address="%v"/>`+"\n", rec[j])
				continue
			}
			if strings.Contains(s, "IM") {
				imBuf += fmt.Sprintf(`<gd:im label="%v" address="%v"/>`+"\n", s, rec[j])
				continue
			}
			if strings.Contains(s, "Custom Key") {
				udKeys = append(udKeys, rec[j])
				continue
			}
			if strings.Contains(s, "Custom Value") {
				udValues = append(udValues, rec[j])
				continue
			}
			switch s {
			case "Name":
				nameBuf += fmt.Sprintf("\n"+`<gd:fullName>%v</gd:fullName>`, rec[j])
			case "GivenName":
				nameBuf += fmt.Sprintf("\n"+`<gd:givenName>%v</gd:givenName>`, rec[j])
			case "FamilyName":
				nameBuf += fmt.Sprintf("\n"+`<gd:familyName>%v</gd:familyName>`, rec[j])
			case "Company":
				orgBuf += fmt.Sprintf(`<gd:orgName>%v</gd:orgName>`+"\n", rec[j])
			case "Job Title":
				orgBuf += fmt.Sprintf(`<gd:orgTitle>%v</gd:orgTitle>`+"\n", rec[j])
			case "Department":
				orgBuf += fmt.Sprintf(`<gd:orgDepartment>%v</gd:orgDepartment>`+"\n", rec[j])
			case "Job Description":
				orgBuf += fmt.Sprintf(`<gd:orgJobDescription>%v</gd:orgJobDescription>`+"\n", rec[j])
			case "Business Fax":
				phoneBuf += fmt.Sprintf(`<gd:phoneNumber rel="http://schemas.google.com/g/2005#work_fax" >%v</gd:phoneNumber>`+"\n", rec[j])
			case "Business Phone":
				phoneBuf += fmt.Sprintf(`<gd:phoneNumber rel="http://schemas.google.com/g/2005#work" >%v</gd:phoneNumber>`+"\n", rec[j])
			case "Business Phone 2":
				phoneBuf += fmt.Sprintf(`<gd:phoneNumber rel="http://schemas.google.com/g/2005#other" >%v</gd:phoneNumber>`+"\n", rec[j])
			case "Home Fax":
				phoneBuf += fmt.Sprintf(`<gd:phoneNumber rel="http://schemas.google.com/g/2005#home_fax" >%v</gd:phoneNumber>`+"\n", rec[j])
			case "Home Phone":
				phoneBuf += fmt.Sprintf(`<gd:phoneNumber rel="http://schemas.google.com/g/2005#home" >%v</gd:phoneNumber>`+"\n", rec[j])
			case "Home Phone 2":
				phoneBuf += fmt.Sprintf(`<gd:phoneNumber rel="http://schemas.google.com/g/2005#other" >%v</gd:phoneNumber>`+"\n", rec[j])
			case "Other Phone":
				phoneBuf += fmt.Sprintf(`<gd:phoneNumber rel="http://schemas.google.com/g/2005#other" >%v</gd:phoneNumber>`+"\n", rec[j])
			case "Mobile Phone":
				phoneBuf += fmt.Sprintf(`<gd:phoneNumber rel="http://schemas.google.com/g/2005#mobile" >%v</gd:phoneNumber>`+"\n", rec[j])
			case "Pager":
				phoneBuf += fmt.Sprintf(`<gd:phoneNumber rel="http://schemas.google.com/g/2005#pager" >%v</gd:phoneNumber>`+"\n", rec[j])
			case "Home Address":
				postalAdress += fmt.Sprintf(`<gd:structuredPostalAddress rel='http://schemas.google.com/g/2005#home' >` + "\n")
				postalAdress += fmt.Sprintf(`<gd:formattedAddress>%v</gd:formattedAddress>`+"\n", rec[j])
				postalAdress += "</gd:structuredPostalAddress>\n"
			case "Other Address":
				postalAdress += fmt.Sprintf(`<gd:structuredPostalAddress rel='http://schemas.google.com/g/2005#other '>` + "\n")
				postalAdress += fmt.Sprintf(`<gd:formattedAddress>%v</gd:formattedAddress>`+"\n", rec[j])
				postalAdress += "</gd:structuredPostalAddress>\n"
			case "Business Address":
				postalAdress += fmt.Sprintf(`<gd:structuredPostalAddress rel='http://schemas.google.com/g/2005#work' >` + "\n")
				postalAdress += fmt.Sprintf(`<gd:formattedAddress>%v</gd:formattedAddress>`+"\n", rec[j])
				postalAdress += "</gd:structuredPostalAddress>\n"
			case "Website Home-Page":
				webBuf += fmt.Sprintf(`<gContact:website href="%v" rel="home-page"/>`+"\n", rec[j])
			case "Web Page":
				webBuf += fmt.Sprintf(`<gContact:website href="%v" rel="other"/>`+"\n", rec[j])
			case "Website Blog":
				webBuf += fmt.Sprintf(`<gContact:website href="%v" rel="blog"/>`+"\n", rec[j])
			case "Website Profile":
				webBuf += fmt.Sprintf(`<gContact:website href="%v" rel="profile"/>`+"\n", rec[j])
			case "Website Home":
				webBuf += fmt.Sprintf(`<gContact:website href="%v" rel="home"/>`+"\n", rec[j])
			case "Website Work":
				webBuf += fmt.Sprintf(`<gContact:website href="%v" rel="work"/>`+"\n", rec[j])
			case "Website Other":
				webBuf += fmt.Sprintf(`<gContact:website href="%v" rel="other"/>`+"\n", rec[j])
			case "Website FTP":
				webBuf += fmt.Sprintf(`<gContact:website href="%v" rel="ftp"/>`+"\n", rec[j])
			case "Notes":
				notesBuf += fmt.Sprintf("<atom:content type='text'>%v</atom:content>\n", rec[j])
			case "birthday":
				birthBuf += fmt.Sprintf("<gContact:birthday when='%v'/>\n", rec[j])
			case "NickName":
				nickBuf += fmt.Sprintf("<gContact:nickname>%v</gContact:nickname>\n", rec[j])
			case "ExternalId":
				externalIdBuf += fmt.Sprintf("<gContact:externalId value='%v'/>\n", rec[j])
			case "Occupation":
				occBuf += fmt.Sprintf("<gContact:occupation>%v</gContact:occupation>\n", rec[j])
			default:
				if numExtendedProp >= maxExtendedProp {
					fmt.Fprintf(w, "Skipped Property Name='%v' value='%v'\n", s, rec[j])
					continue
				}
				//fmt.Fprintf(w, "Mapped Property Name='%v' value='%v'\n", s, rec[j])
				extendedBuf += fmt.Sprintf(`<gd:extendedProperty name='%v' value="%v" />`+"\n", s, rec[j])
				numExtendedProp += 1
				break
			}
		}
		if len(notesBuf) > 0 {
			fmt.Fprintf(buf, notesBuf)
		}
		nameBuf = "<gd:name>" + nameBuf + "\n</gd:name>\n"

		fmt.Fprintf(buf, nameBuf)
		fmt.Fprintf(buf, emailBuf)
		fmt.Fprintf(buf, imBuf)
		orgBuf += "</gd:organization>\n"
		fmt.Fprintf(buf, orgBuf)
		fmt.Fprintf(buf, phoneBuf)
		fmt.Fprintf(buf, postalAdress)
		fmt.Fprintf(buf, extendedBuf)
		fmt.Fprintf(buf, webBuf)
		if len(birthBuf) > 0 {
			fmt.Fprintf(buf, birthBuf)
		}
		if len(nickBuf) > 0 {
			fmt.Fprintf(buf, nickBuf)
		}
		if len(externalIdBuf) > 0 {
			fmt.Fprintf(buf, externalIdBuf)
		}
		if len(occBuf) > 0 {
			fmt.Fprintf(buf, occBuf)
		}

		for j, key := range udKeys {
			fmt.Fprintf(buf, `<gContact:userDefinedField key="%v" value="%v"/>`+"\n", key, udValues[j])
		}

		fmt.Fprintf(buf, `</atom:entry>`)

		res, _ := client.Post(fmt.Sprintf(feedUrl, state.Domain), `application/atom+xml`, strings.NewReader(buf.String()))

		ictr++
		fmt.Fprintf(w, "Result[%v]: %v<br/>", ictr, res.Status)
		if res.StatusCode != 201 {
			fmt.Fprintf(w, buf.String())
		}

	}
}
