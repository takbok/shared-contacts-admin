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
