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

	"appengine"
)

func init() {
	http.HandleFunc("/contacts/exportxml", exportXML)
}

func initiateContactsXmlExport(w http.ResponseWriter, r *http.Request, url string) {
	ctx := appengine.NewContext(r)
	config.RedirectURL = fmt.Sprintf(`http://%s/contacts/exportxml`, r.Host)

	x := AppState{url}
	url = config.AuthCodeURL(x.encodeState())
	ctx.Infof("Auth: %v", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func exportXML(w http.ResponseWriter, r *http.Request) {
	y := r.FormValue("state")

	state := new(AppState)
	state.decodeState(y)

	w.Header().Set(`Content-Type`, `application/xml`)
	w.Header().Set(`Content-Disposition`, `attachment; filename="`+state.Domain+`-contacts-export.xml"`)

	ctx := appengine.NewContext(r)

	buf := loadFullFeed(state.Domain, ctx, r)
	
	fmt.Fprintf(w, buf.String())
}

