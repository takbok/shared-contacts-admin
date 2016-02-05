package demo

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	config = &oauth2.Config{
		ClientID:     `client-id-goes-here`,
		ClientSecret: `secret-goes-here`,
		RedirectURL:  `ideally-should-be-set-later`,
		Scopes:       []string{`http://www.google.com/m8/feeds/contacts/`},
		Endpoint:     google.Endpoint,
	}

	yeah = "yeah"
)
