package gripp

// this is all the data the gripp api can return for a project. some are readonly
type Project struct {
	Createdon               Date                     `json:"createdon,omitempty"`
	Updatedon               Date                     `json:"updatedon,omitempty"`
	ID                      int                      `json:"id,omitempty"`
	CustomFields            interface{}              `json:"customfields,omitempty"`
	TemplateSet             WithSearchName[int]      `json:"templateset"`
	SearchName              string                   `json:"searchname,omitempty"`
	Name                    string                   `json:"name"`
	Color                   string                   `json:"color,omitempty"`
	ValidFor                WithSearchName[int]      `json:"validfor,omitempty"`
	AccountManager          interface{}              `json:"accountmanager,omitempty"`
	FilesAvailableForClient bool                     `json:"filesavailableforclient,omitempty"`
	Number                  int                      `json:"number,omitempty"`
	Phase                   WithSearchName[int]      `json:"phase,omitempty"`
	Deadline                Date                     `json:"deadline,omitempty"`
	Company                 WithSearchName[int]      `json:"company"`
	Contact                 WithSearchName[int]      `json:"contact,omitempty"`
	StartDate               Date                     `json:"startdate,omitempty"`
	DeliveryDate            Date                     `json:"deliverydate,omitempty"`
	EndDate                 Date                     `json:"enddate,omitempty"`
	AddHoursSpecification   bool                     `json:"addhoursspecification,omitempty"`
	Description             string                   `json:"description,omitempty"`
	WorkDeliveryAddress     string                   `json:"workdeliveryaddress,omitempty"`
	ClientReference         string                   `json:"clientreference,omitempty"`
	IsBasis                 bool                     `json:"isbasis,omitempty"`
	TotalInclVat            string                   `json:"totalinclvat,omitempty"`
	TotalExclVat            string                   `json:"totalexclvat,omitempty"`
	Archived                bool                     `json:"archived,omitempty"`
	ArchivedOn              Date                     `json:"archivedon,omitempty"`
	Tags                    []WithSearchName[string] `json:"tags,omitempty"`
	Employees               []WithSearchName[string] `json:"employees,omitempty"`
	ProjectLines            []ProjectLine            `json:"projectlines,omitempty"`
}

type ProjectRepository struct {
	builder *requestBuilder[Project]
}

func (p *ProjectRepository) Get() ([]Project, error) {
	return p.builder.Get()
}

func (p *ProjectRepository) GetOne() (*Project, error) {
	return p.builder.GetOne()
}

func (p *ProjectRepository) Delete(projectid int) error {
	return p.builder.Delete(projectid)
}

func (p *ProjectRepository) Filter(input ...interface{}) *ProjectRepository {
	p.builder.Filter(input...)
	return p
}

func (p *ProjectRepository) Archived(archived bool) *ProjectRepository {
	p.builder.Filter("archived", archived)
	return p
}

func (p *ProjectRepository) ByEmployee(employeeIDs ...int) *ProjectRepository {
	if len(employeeIDs) == 0 {
		return p
	}
	p.builder.Filter("employees", "in", employeeIDs)
	return p
}

func (p *ProjectRepository) Search(column string, value string) *ProjectRepository {
	p.builder.Filter(column, "like", "%"+value+"%")
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
