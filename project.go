package gripp

// this is all the data the gripp api can return for a project. some are readonly
type Project struct {
	Createdon               Date           `json:"createdon,omitempty"`
	Updatedon               Date           `json:"updatedon,omitempty"`
	ID                      string         `json:"id,omitempty"`
	CustomFields            interface{}    `json:"customfields,omitempty"`
	TemplateSet             WithSearchName `json:"templateset"`
	SearchName              string         `json:"searchname,omitempty"`
	Name                    string         `json:"name"`
	Color                   string         `json:"color,omitempty"`
	ValidFor                int            `json:"validfor,omitempty"`
	AccountManager          int            `json:"accountmanager,omitempty"`
	FilesAvailableForClient bool           `json:"filesavailableforclient,omitempty"`
	Number                  string         `json:"number,omitempty"`
	Phase                   WithSearchName `json:"phase,omitempty"`
	Deadline                Date           `json:"deadline,omitempty"`
	Company                 WithSearchName `json:"company"`
	Contact                 WithSearchName `json:"contact,omitempty"`
	StartDate               Date           `json:"startdate,omitempty"`
	DeliveryDate            Date           `json:"deliverydate,omitempty"`
	EndDate                 Date           `json:"enddate,omitempty"`
	AddHoursSpecification   bool           `json:"addhoursspecification,omitempty"`
	Description             string         `json:"description,omitempty"`
	WorkDeliveryAddress     string         `json:"workdeliveryaddress,omitempty"`
	ClientReference         string         `json:"clientreference,omitempty"`
	IsBasis                 bool           `json:"isbasis,omitempty"`
	TotalInclVat            float64        `json:"totalinclvat,omitempty"`
	TotalExclVat            float64        `json:"totalexclvat,omitempty"`
	Archived                bool           `json:"archived,omitempty"`
	ArchivedOn              Date           `json:"archivedon,omitempty"`
	ExtendedProperties      string         `json:"extendedproperties,omitempty"`
	Tags                    []int          `json:"tags,omitempty"`
	Employees               []int          `json:"employees,omitempty"`
	EmployeesStarred        bool           `json:"employees_starred,omitempty"`
	ExtraPdf1               int            `json:"extrapdf1,omitempty"`
	ExtraPdf2               int            `json:"extrapdf2,omitempty"`
	Files                   []int          `json:"files,omitempty"`
	ProjectLines            []int          `json:"projectlines,omitempty"`
	UmbrellaProject         int            `json:"umbrella_project,omitempty"`
}

type WithSearchName struct {
	ID         uint   `json:"id"`
	SearchName string `json:"searchname"`
	Discr      string `json:"discr"`
}

func (c *Client) Projects() *RequestBuilder[Project] {
	return &RequestBuilder[Project]{
		client: c,
		base:   "project",
	}
}
