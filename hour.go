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
	Amount       int         `json:"amount"`
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
