// Package constants ... This package exports various constant values used by the app.
package constants

const (
	// SolrPort ... The port on which solr runs.
	SolrPort = 8983

	// ApplicationPort ... The port on which application will run.
	ApplicationPort = ":8080"

	// ApplicationHost ... Application Host
	ApplicationHost = "127.0.0.1"

	// DocsCore ... Docs Core name in solr
	DocsCore = "docs"

	// UserCore ... Users Core name in solr
	UserCore = "user"

	// SolrHost .... Solr Host
	SolrHost = "127.0.0.1"

	//MinDocNumLen ... Min document number length
	MinDocNumLen = 3
)
