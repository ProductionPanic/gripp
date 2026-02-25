package gripp

type Employee struct {
	Username      string                   `json:"username"`
	Phone         string                   `json:"phone"`
	Active        bool                     `json:"active"`
	Email         string                   `json:"email"`
	Mobile        string                   `json:"mobile"`
	Street        string                   `json:"street"`
	Streetnumber  string                   `json:"streetnumber"`
	Adresline2    string                   `json:"adresline2"`
	Zipcode       string                   `json:"zipcode"`
	City          string                   `json:"city"`
	Country       string                   `json:"country"`
	Function      string                   `json:"function"`
	Initials      string                   `json:"initials"`
	Firstname     string                   `json:"firstname"`
	Infix         string                   `json:"infix"`
	Lastname      string                   `json:"lastname"`
	Screenname    string                   `json:"screenname"`
	ID            int                      `json:"id"`
	Createdon     Date                     `json:"createdon"`
	Updatedon     Date                     `json:"updatedon"`
	Searchname    string                   `json:"searchname"`
	Dateofbirth   Date                     `json:"dateofbirth"`
	Employeesince Date                     `json:"employeesince"`
	Discr         string                   `json:"discr"`
	Userphoto     WithSearchName[int]      `json:"userphoto"`
	Role          WithSearchName[int]      `json:"role"`
	Salutation    WithSearchName[int]      `json:"salutation"`
	Department    WithSearchName[int]      `json:"department"`
	Tags          []WithSearchName[string] `json:"tags"`
}

func (c *Client) Employee() *EmployeeRepository {
	return &EmployeeRepository{
		builder: &requestBuilder[Employee]{
			client: c,
			base:   "employee",
		},
	}
}

type EmployeeRepository struct {
	builder *requestBuilder[Employee]
}

func (e *EmployeeRepository) Get() ([]Employee, error) {
	return e.builder.Get()
}

func (e *EmployeeRepository) GetOne() (*Employee, error) {
	return e.builder.GetOne()
}

func (e *EmployeeRepository) Delete(employeeid int) error {
	return e.builder.Delete(employeeid)
}

func (e *EmployeeRepository) Filter(input ...interface{}) *EmployeeRepository {
	e.builder.Filter(input...)
	return e
}

func (e *EmployeeRepository) Page(firstResult, maxResults int) *EmployeeRepository {
	e.builder.Page(firstResult, maxResults)
	return e
}

func (e *EmployeeRepository) OrderBy(field string, direction string) *EmployeeRepository {
	e.builder.OrderBy(field, direction)
	return e
}
