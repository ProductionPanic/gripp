package gripp

type ProjectLine struct {
	client             *Client
	Ordering           int     `json:"_ordering"`
	Internalnote       string  `json:"internalnote"`
	Amount             float64 `json:"amount"`
	Hidefortimewriting bool    `json:"hidefortimewriting"`
	Sellingprice       string  `json:"sellingprice"`
	Discount           int     `json:"discount"`
	Buyingprice        string  `json:"buyingprice"`
	Additionalsubject  string  `json:"additionalsubject"`
	Description        string  `json:"description"`
	Hidedetails        bool    `json:"hidedetails"`
	ID                 int     `json:"id"`
	Createdon          struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	} `json:"createdon"`
	Updatedon struct {
		Date         string `json:"date"`
		TimezoneType int    `json:"timezone_type"`
		Timezone     string `json:"timezone"`
	} `json:"updatedon"`
	Searchname         string      `json:"searchname"`
	Extendedproperties interface{} `json:"extendedproperties"`
	Groupcategory      interface{} `json:"groupcategory"`
	Convertto          interface{} `json:"convertto"`
	Unit               struct {
		ID         int    `json:"id"`
		Searchname string `json:"searchname"`
	} `json:"unit"`
	Invoicebasis struct {
		ID         int    `json:"id"`
		Searchname string `json:"searchname"`
	} `json:"invoicebasis"`
	Vat struct {
		ID         int    `json:"id"`
		Searchname string `json:"searchname"`
	} `json:"vat"`
	Rowtype struct {
		ID         int    `json:"id"`
		Searchname string `json:"searchname"`
	} `json:"rowtype"`
	Offerprojectbase struct {
		ID         int    `json:"id"`
		Searchname string `json:"searchname"`
		Discr      string `json:"discr"`
	} `json:"offerprojectbase"`
	Contractline interface{} `json:"contractline"`
	Product      struct {
		ID         int    `json:"id"`
		Searchname string `json:"searchname"`
		Discr      string `json:"discr"`
	} `json:"product"`
	Amountwritten string `json:"amountwritten"`
}
