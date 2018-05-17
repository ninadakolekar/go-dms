package models

// InactiveDoc ... Inactive Document
type InactiveDoc struct {
	DocNo        string
	Title        string
	DocType      string
	DocStatus    bool
	Initiator    string
	Creator      string
	Reviewer     []string
	Approver     []string
	Authorizer   []string
	DocDept      string
	FlowStatus   uint32
	DocTemplate  uint32
	DocumentBody string
}
