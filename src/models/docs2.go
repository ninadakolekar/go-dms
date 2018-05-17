package models

// Docs2 ... docs2 core document schema
type Docs2 struct {
	DocNo       string
	Title       string
	DocType     string
	DocStatus   bool
	Initiator   string
	Creator     string
	Reviewer    []string
	Approver    []string
	Authorizer  []string
	DocDept     string
	FlowStatus  float64
	DocTemplate float64
	Template    string
	Paragraph   []string
	Images      []string
	Tables      []string
}

//Init ... add dummy variables
func (d Docs2) Init(id string) {
	d.DocNo = id
	d.Approver = []string{""}
	d.Authorizer = []string{""}
	d.Paragraph = []string{""}
	d.Creator = "ramkishan"
	d.DocDept = "department"
	d.DocStatus = false
	d.DocTemplate = 0.00
	d.DocType = ""
	d.FlowStatus = 0.00
	d.Images = []string{"1010101010101010"}
	d.Initiator = ""
	d.Reviewer = []string{""}
	d.Tables = []string{""}
	d.Template = ""
	d.Title = ""
}
