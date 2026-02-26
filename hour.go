package gripp

type HourCreateRequest struct {
	Status           string  `json:"status"`
	Date             string  `json:"date"`
	Description      string  `json:"description"`
	Amount           float64 `json:"amount"`
	Employee         int     `json:"employee"`
	Offerprojectline int     `json:"offerprojectline"`
}

type Hour struct {
	Date struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	} `json:"date"`
	Amount       float64     `json:"amount"`
	Description  string      `json:"description"`
	Authorizedon interface{} `json:"authorizedon"`
	Definitiveon interface{} `json:"definitiveon"`
	ID           int         `json:"id"`
	Createdon    struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	} `json:"createdon"`
	Updatedon          interface{} `json:"updatedon"`
	Searchname         string      `json:"searchname"`
	Extendedproperties interface{} `json:"extendedproperties"`
	Task               interface{} `json:"task"`
	Status             struct {
		ID         int    `json:"id"`
		Searchname string `json:"searchname"`
	} `json:"status"`
	Employee struct {
		ID         int    `json:"id"`
		Searchname string `json:"searchname"`
		Discr      string `json:"discr"`
	} `json:"employee"`
	Offerprojectbase struct {
		ID         int    `json:"id"`
		Searchname string `json:"searchname"`
		Discr      string `json:"discr"`
	} `json:"offerprojectbase"`
	Offerprojectline struct {
		ID         int    `json:"id"`
		Searchname string `json:"searchname"`
	} `json:"offerprojectline"`
	Invoiceline  interface{} `json:"invoiceline"`
	Authorizedby interface{} `json:"authorizedby"`
	Definitiveby interface{} `json:"definitiveby"`
}

type HourRepository struct {
	builder *requestBuilder[Hour]
}

func (h *HourRepository) Get() ([]Hour, error) {
	return h.builder.Get()
}

func (h *HourRepository) GetOne() (*Hour, error) {
	return h.builder.GetOne()
}

func (h *HourRepository) ByProjectLineID(ProjectLineIDs ...int) *HourRepository {
	if len(ProjectLineIDs) == 0 {
		return h
	} else if len(ProjectLineIDs) == 1 {
		h.builder.Filter("offerprojectline", ProjectLineIDs[0])
		return h
	}
	h.builder.Filter("offerprojectline", "in", ProjectLineIDs)
	return h
}

func (h *HourRepository) ByProjectLine(projectLines ...*ProjectLine) *HourRepository {
	if len(projectLines) == 0 {
		return h
	} else if len(projectLines) == 1 {
		h.builder.Filter("offerprojectline", projectLines[0].ID)
		return h
	}
	ids := make([]int, len(projectLines))
	for i, pl := range projectLines {
		ids[i] = pl.ID
	}
	h.builder.Filter("offerprojectline", "in", ids)
	return h
}

func (h *HourRepository) Delete(hourid int) error {
	return h.builder.Delete(hourid)
}

func (h *HourRepository) Filter(input ...interface{}) *HourRepository {
	h.builder.Filter(input...)
	return h
}

func (h *HourRepository) Page(firstResult, maxResults int) *HourRepository {
	h.builder.Page(firstResult, maxResults)
	return h
}

// TODO:
//func (h *HourRepository) Create(hour HourCreateRequest) (*Hour, error) {
//	return h.builder.Create(hour)
//}
