package demo

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
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
		log.Print("\n returning bcoz of error 1")
		log.Print(err)
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
		log.Print("\n CSV file error")
		ctx.Errorf("%v", err)
		return
	}

	names := records[0]
	datalen := len(records)
	log.Print("\n Loop started")

	for i := 1; i < datalen; i++ {
		rec := records[i]
		buf := new(bytes.Buffer)
		fmt.Fprintf(buf, `<atom:entry xmlns:atom='http://www.w3.org/2005/Atom' xmlns:gd='http://schemas.google.com/g/2005'>
<atom:category scheme='http://schemas.google.com/g/2005#kind' term='http://schemas.google.com/contact/2008#contact' />
<atom:content type='text'>Notes</atom:content>
`)
		numExtended, M := 0, 10
		for j, s := range names {
			if s == "Name" {
				fmt.Fprintf(buf, "<gd:name><gd:fullName>%v</gd:fullName></gd:name>\n", rec[j])
			} else if s == "E-mail Address" {
				fmt.Fprintf(buf, "<gd:email rel='http://schemas.google.com/g/2005#home' address='%v'/>", rec[j])
			} else if strings.HasPrefix(s, "E-mail ") && strings.HasSuffix(s, " Address") {
				var num uint
				fmt.Sscanf(s, "E-mail %v Address", &num)
				if numExtended < M && ((0 < num && num < 6) || s == "E-mail Address") {
					fmt.Fprintf(buf, `<gd:extendedProperty name="%v" value="%v" />`+"\n", s, rec[j])
					numExtended++
				}
			} else if numExtended < M {
				fmt.Fprintf(buf, `<gd:extendedProperty name="%v" value="%v" />`+"\n", s, rec[j])
				numExtended++
			}
		}

		fmt.Fprintf(buf, `</atom:entry>`)

		res, _ := client.Post(fmt.Sprintf(feedUrl, state.Domain), `application/atom+xml`, strings.NewReader(buf.String()))

		fmt.Fprintf(w, "Result: %v<br/>", res.Status)
	}
}
