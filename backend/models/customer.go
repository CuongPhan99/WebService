package model

type Customers struct {
	Id          uint   `json:"id" pg:"type:serial"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	CompanyName string `json:"company_name"`
}
