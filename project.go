package gripp

// this is all the data the gripp api can return for a project. some are readonly
type Project struct {
	Createdon               Date           `json:"createdon,omitempty"`
	Updatedon               Date           `json:"updatedon,omitempty"`
	ID                      int            `json:"id,omitempty"`
	CustomFields            interface{}    `json:"customfields,omitempty"`
	TemplateSet             WithSearchName `json:"templateset"`
	SearchName              string         `json:"searchname,omitempty"`
	Name                    string         `json:"name"`
	Color                   string         `json:"color,omitempty"`
	ValidFor                WithSearchName `json:"validfor,omitempty"`
	AccountManager          interface{}    `json:"accountmanager,omitempty"`
	FilesAvailableForClient bool           `json:"filesavailableforclient,omitempty"`
	Number                  int            `json:"number,omitempty"`
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
	TotalInclVat            string         `json:"totalinclvat,omitempty"`
	TotalExclVat            string         `json:"totalexclvat,omitempty"`
	Archived                bool           `json:"archived,omitempty"`
	ArchivedOn              Date           `json:"archivedon,omitempty"`
	ExtendedProperties      string         `json:"extendedproperties,omitempty"`
	Tags                    []interface{}  `json:"tags,omitempty"`
	Employees               []interface{}  `json:"employees,omitempty"`
	EmployeesStarred        []interface{}  `json:"employees_starred,omitempty"`
	ExtraPdf1               int            `json:"extrapdf1,omitempty"`
	ExtraPdf2               int            `json:"extrapdf2,omitempty"`
	Files                   []interface{}  `json:"files,omitempty"`
	ProjectLines            []interface{}  `json:"projectlines,omitempty"`
	UmbrellaProject         int            `json:"umbrella_project,omitempty"`
}

type WithSearchName struct {
	ID         uint   `json:"id"`
	SearchName string `json:"searchname"`
	Discr      string `json:"discr"`
}

type ProjectRepository struct {
	builder *RequestBuilder[Project]
}

func (p *ProjectRepository) Get() ([]Project, error) {
	return p.builder.Get()
}

func (p *ProjectRepository) Filter(input ...interface{}) *ProjectRepository {
	p.builder.Filter(input...)
	return p
}

func (p *ProjectRepository) Page(firstResult, maxResults int) *ProjectRepository {
	p.builder.Page(firstResult, maxResults)
	return p
}

func (p *ProjectRepository) OrderBy(field, direction string) *ProjectRepository {
	p.builder.OrderBy(field, direction)
	return p
}

func (c *Client) Projects() *ProjectRepository {
	return &ProjectRepository{
		builder: &RequestBuilder[Project]{
			client: c,
			base:   "project",
		},
	}
}
