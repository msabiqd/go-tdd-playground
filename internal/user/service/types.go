package service

type CreateEmployeeRequest struct {
	Email    string                  `json:"email" validate:"required,email"`
	Password string                  `json:"password" validate:"required"`
	Families []CreateFamiliesRequest `json:"families"`
}

type CreateFamiliesRequest struct {
	EmployeeID string `json:"employee_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Relation   string `json:"relation"`
}

type CreateEmployeeResponse struct {
	ID       string       `json:"id"`
	Email    string       `json:"email"`
	Families []FamilyData `json:"families"`
}

type InsertEmployeeRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InsertEmployeeResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type InsertFamiliesDataRequest struct {
	Families []InsertFamilyDataRequest `json:"families"`
}

type InsertFamiliesDataResponse struct {
	Families []FamilyData `json:"families"`
}

type InsertFamilyDataRequest struct {
	EmployeeID string `json:"employee_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Relation   string `json:"relation"`
}

type FamilyData struct {
	ID         string `json:"id"`
	EmployeeID string `json:"employee_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Relation   string `json:"relation"`
}
