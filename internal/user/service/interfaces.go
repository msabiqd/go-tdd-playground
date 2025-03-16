package service

type UserServiceInterface interface {
	CreateEmployee(CreateEmployeeRequest) (CreateEmployeeResponse, error)
}
