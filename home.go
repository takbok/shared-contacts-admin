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
	"fmt"
	"net/http"
)

var html = `<!DOCTYPE html>
<html>
  <head>
    <title>Shared Contacts Exporter</title>
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
	%s
	<form action="/contacts" method="post">
	  <span> Domain hosted with Google Apps for Business </span>
	  <label for="app_url"></label> <input id="app_url" type="url" name="url" placeholder="http://www.example.com" />
	  <input type="submit" value="Set Domain & Export CSV" />
	</form>
	<br/><hr/>
	<form enctype="multipart/form-data" action="/import" method="post">
      <input type="file" name="inputfile" /><br/>
	  <span> Domain hosted with Google Apps for Business </span>
	  <label for="app_url"></label> <input id="app_url" type="url" name="url" placeholder="http://www.example.com" />
      <input type="submit" value="Import CSV" />
	</form>
	<br/><hr/>
	<form enctype="multipart/form-data" action="/set-action" method="post">
	  <span> Domain hosted with Google Apps for Business </span>
	  <label for="app_url"></label> <input id="app_url" type="url" name="url" placeholder="http://www.example.com" />
	  <button type="submit" name="what" value="delete">Delete All Contacts</button>
	</form>
	<br/><hr/>
  </body>
</html>
`

func init() {
	http.HandleFunc("/", handleHomePage)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	err := r.FormValue("error")
	message := ""

	switch err {
	case "notOnGoogleApps":
		message = `<h4> This URL is not hosted on Google Apps </h4>`
	case "badUrl":
		message = `<h4> An invalid URL was entered </h4>`
	}

	fmt.Fprintf(w, html, message)
}
