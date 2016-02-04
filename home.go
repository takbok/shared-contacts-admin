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
    <a class="button" href="/contacts">Get Shared Contacts</a>
	<br/><hr/>
	<form enctype="multipart/form-data" action="/import" method="post">
      <input type="file" name="inputfile" />
      <input type="submit" value="Import" />
	</form>
	<br/><hr/>
    <a class="button" href="/export">Export</a><br/>
  </body>
</html>
`

func init() {
	http.HandleFunc("/", handleHomePage)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html)
}
