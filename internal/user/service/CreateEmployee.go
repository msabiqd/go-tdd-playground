package service

import (
	"go-tdd-playground/internal/user/repository"

	validator "github.com/go-playground/validator/v10"
)

func (s UserService) CreateEmployee(req CreateEmployeeRequest) (resp CreateEmployeeResponse, err error) {
	// if req.Email == "" {
	// 	err = ErrInvalidParam
	// 	return
	// }

	// if req.Password == "" {
	// 	err = ErrInvalidParam
	// 	return
	// }
	validate := validator.New()
	validationErr := validate.Struct(req)
	if validationErr != nil {
		err = ErrInvalidParam
		return
	}

	repoOutput, err := s.UserRepository.InsertEmployee(repository.InsertEmployeeRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return
	}

	if len(req.Families) > 0 {
		_, _ = s.UserRepository.InsertFamilies(repository.InsertFamiliesDataRequest{
			Families: []repository.InsertFamilyDataRequest{
				{
					EmployeeID: repoOutput.ID,
					FirstName:  req.Families[0].FirstName,
					LastName:   req.Families[0].LastName,
					Relation:   req.Families[0].Relation,
				},
			}})
	}

	return CreateEmployeeResponse{
		Email: repoOutput.Email,
		ID:    repoOutput.ID,
		// Families: []FamilyData{
		// 	// {
		// 	// 	ID:         repoFamiliesOutput.Families[0].ID,
		// 	// 	EmployeeID: repoFamiliesOutput.Families[0].EmployeeID,
		// 	// 	FirstName:  repoFamiliesOutput.Families[0].FirstName,
		// 	// 	LastName:   repoFamiliesOutput.Families[0].LastName,
		// 	// 	Relation:   repoFamiliesOutput.Families[0].Relation,
		// 	// },
		// },
	}, nil
}
