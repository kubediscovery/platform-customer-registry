package entity

type CustomerRegistry interface {}

type CustomerRegistry struct {
	ProjectName string `json:'project_name' binding:'required'`
	Repository string `json:'repository' binding:'required'`
	UserName string `json:'user_name' binding:'required'`
	UserEmail string `json:'user_email' binding:'required'`
	CreateAt string `json:'create_at,omitempety'`
	EndAt string `json:'end_at,omitempety'`
}

func (cr *CustomerRegistry) Validate() error{
	return nil
}

func (cr *CustomerRegistry) RegistryNewCustomer() (*CustomerRegistry, error){
	return nil,nil
}