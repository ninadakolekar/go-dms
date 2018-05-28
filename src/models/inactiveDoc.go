package models

// InactiveDoc ... Inactive Document
type InactiveDoc struct {
	DocNo           string
	Title           string
	DocType         string
	DocProcess      string
	DocEffDate      string
	DocExpDate      string
	DocStatus       bool
	Initiator       string
	Creator         string
	Reviewer        []string
	Approver        []string
	Authorizer      []string
	DocDept         string
	FlowStatus      float64
	FlowList        []string
	CurrentFlowUser float64
	DocTemplate     float64
	InitTS          string
	CreateTS        string
	ReviewTS        string
	AuthTS          string
	ApproveTS       string
	DocumentBody    []string
	QA              string
}
