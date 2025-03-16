package repository

type UserRepositoryInterface interface {
	InsertEmployee(InsertEmployeeRequest) (InsertEmployeeResponse, error)
	InsertFamilies(InsertFamiliesDataRequest) (InsertFamiliesDataResponse, error)
}
