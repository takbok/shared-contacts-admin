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
	//"fmt"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/set-action", setAction)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	//	err := r.FormValue("error")
	//	message := ""
	//
	//	switch err {
	//	case "notOnGoogleApps":
	//		message = `<h4> This URL is not hosted on Google Apps </h4>`
	//	case "badUrl":
	//		message = `<h4> An invalid URL was entered </h4>`
	//	}

	//fmt.Fprintf(w, html, message)
	//fmt.Fprintf(w, html, message)
	if err := htmlBody1.Execute(w, ""); err != nil {
		panic(err)
	}

}

func setAction(w http.ResponseWriter, r *http.Request) {
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

	switch r.FormValue("what") {
	case "delete":
		initiateContactsDeletion(w, r, url)
	case "xmlExport":
		initiateContactsXmlExport(w, r, url)
	}
}

var htmlBody1 = template.Must(template.New("User").Parse(`<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Google Apps - Shared Contacts Admin API</title>
  <meta name="title" content="Google Apps - Shared Contacts Admin API | Takbok" />
  <meta name="description" content="Manage Domain Shared Contacts API" />
  <meta name="viewport" content="initial-scale=1, width=device-width, maximum-scale=1, minimum-scale=1, user-scalable=no">
  <meta name="author" content="Takbok">
  <link rel="shortcut icon" id="favicon" href="favicon.png">
  <link href='http://fonts.googleapis.com/css?family=Open+Sans:300,400,700' rel='stylesheet' type='text/css'>
  <link href='http://fonts.googleapis.com/css?family=Pacifico:400' rel='stylesheet' type='text/css'>
  <link href='/css/onepage-scroll.css' rel='stylesheet' type='text/css'>
  <link href='/css/onepage-scroll-demo.css' rel='stylesheet' type='text/css'>
  <link rel="stylesheet" media="screen,projection,tv" href="/css/modalWindow.css" />
  <script type="text/javascript" src="http://code.jquery.com/jquery-1.11.0.js"></script>
  <script type="text/javascript" src="/js/jquery.onepage-scroll.js"></script>
  <script>
    $(document).ready(function(){
      $(".main").onepage_scroll({
        sectionContainer: "section",
        loop: true,
        responsiveFallback: false
      });
    });
  </script>
</head>
<body>

    <div class="main">
      <section class="page1">
        <div class="page_container">
          <h1>Shared Contacts Admin</h1>
          <h2>A web based application for administering Google Apps Domain Shared Contacts</h2>
          <p class="credit">Created by <a href="http://www.takbok.com">Takbok</a>, a company that provides speedy software solutions for your business</p>

          <div class="btns">
            <a class="reload btn" href="https://github.com/takbok/shared-contacts-admin">Download on Github</a>
            <p class="text-btn"><a href="https://github.com/takbok/shared-contacts-admin/blob/master/README.md">README</a> | <a href="http://golang-programming.appspot.com/slides?TYPE=SLIDE&DOC_ID=87&SID=TDSSLIDE-87">SLIDES</a></p>
          </div>
        </div>
        <img src="/images/gapps.png" alt="gapps">
      </section>

      <section class="page2">
        <div class="page_container">
          <h1>Import Contacts</h1>
          <h2>All you need is a CSV file with contacts.</h2>
          <div class="btns">
            <a class="reload btn" href="#importMod" >Import Contacts</a>
            <p class="text-btn"><a href="https://github.com/takbok/shared-contacts-admin/tree/master/testcases/test-data">Sample CSV</a></p>
          </div>
        </div>
      </section>

      <section class="page3">
        <div class="page_container">
          <h1>Export Contacts</h1>
          <h2>You can export contacts via CSV ad XML formats.</h2>
          <div class="btns">
            <a class="reload btn" href="#exportModCSV">Export to CSV</a> <a class="reload btn" href="#exportModXML">Export to XML</a>
          </div>
        </div>
      </section>

      <section class="page4">
        <div class="page_container">
          <h1>Delete Contacts</h1>
          <h2>You can delete contacts as well.</h2>
          <div class="btns">
            <a class="reload btn" href="#deleteMod">Delete Contacts</a>
          </div>
        </div>
      </section>
    <a class="back">Shared Contacts Admin</a>
    <a class="rehide" href="https://github.com/takbok/shared-contacts-admin"><img style="position: absolute; top: 0; right: 0; border: 0;" src="https://s3.amazonaws.com/github/ribbons/forkme_right_darkblue_121621.png" alt="Fork me on GitHub"></a>
    </div>

<!--MODAL DIALOGS-->

  <div id="importMod" class="modalDialog">
    <a href="#close" title="Close" class="close">X</a>
    <div>
      <form enctype="multipart/form-data" action="/import" method="post">
        <span>Select file & enter domain</span><br>
           <input type="file" name="inputfile" /><br/>
        <label for="app_url"></label> <input id="app_url" type="url" name="url" placeholder="http://www.example.com" />
           <input type="submit" value="Import CSV" />
      </form>
    </div>
  </div>

  <div id="exportModCSV" class="modalDialog">
    <a href="#close" title="Close" class="close">X</a>
    <div>
      <form action="/contacts" method="post">
        <span>Enter Google Apps Domain</span><br>
        <label for="app_url"></label> <input id="app_url" type="url" name="url" placeholder="http://www.example.com" />
        <input type="submit" value="Export CSV" />
      </form>
    </div>
  </div>

  <div id="exportModXML" class="modalDialog">
    <a href="#close" title="Close" class="close">X</a>
    <div>
      <form enctype="multipart/form-data" action="/set-action" method="post">
        <span>Enter Google Apps Domain</span><br>
        <label for="app_url"></label> <input id="app_url" type="url" name="url" placeholder="http://www.example.com" />
        <button type="submit" name="what" value="xmlExport">Export XML</button>
      </form>
    </div>
  </div>

  <div id="deleteMod" class="modalDialog">
    <a href="#close" title="Close" class="close">X</a>
    <div>
      <form enctype="multipart/form-data" action="/set-action" method="post">
        <span>Enter Google Apps Domain</span><br>
        <label for="app_url"></label> <input id="app_url" type="url" name="url" placeholder="http://www.example.com" />
        <button type="submit" name="what" value="delete">Delete All Contacts</button>
      </form>
    </div>
  </div>

</body>
</html>

`))
