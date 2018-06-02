// Package constants ... This package exports various constant values used by the app.
package constants

import "github.com/gorilla/securecookie"

const (
	// SolrPort ... The port on which solr runs.
	SolrPort = 8983

	// ApplicationPort ... The port on which application will run.
	ApplicationPort = ":8080"

	// ApplicationHost ... Application Host
	ApplicationHost = "192.168.136.120"

	// DocsCore ... Docs Core name in solr
	DocsCore = "testdocs"

	// UserCore ... Users Core name in solr
	UserCore = "testuser"

	// SolrHost .... Solr Host
	SolrHost = "192.168.136.120"

	//MinDocNumLen ... Min document number length
	MinDocNumLen = 3

	InitFlow = 0

	QaFlow = 1

	CreateFlow = 2

	ReviewFlow = 3

	ApproveFlow = 4

	AuthFlow = 5

	ActiveFlow = 6

	OneByOne = "OneByOne"

	Everyone = "Everyone"

	Anyone = "Anyone"
)

// CookieHandler ... Random key generator for securecookie
var CookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))
