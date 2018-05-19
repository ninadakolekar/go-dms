package models

// InactiveDoc ... Inactive Document
type InactiveDoc struct {
	DocNo        string
	Title        string
	DocType      string
	DocProcess   string
	DocEffDate   string
	DocExpDate   string
	DocStatus    bool
	Initiator    string
	Creator      string
	Reviewer     []string
	Approver     []string
	Authorizer   []string
	DocDept      string
	FlowStatus   uint32
	DocTemplate  uint32
	InitTS       string
	DocumentBody string
}

/*
	InactiveDoc{
	DocNo  : "doc1"
	Title  : "doc2"
	DocType      string
	DocProcess   string
	DocEffDate   string
	DocExpDate   string
	DocStatus    bool
	Initiator    string
	Creator      string
	Reviewer     []string
	Approver     []string
	Authorizer   []string
	DocDept      string
	FlowStatus   uint32
	DocTemplate  uint32
	InitTS       string
	DocumentBody string}
*/
