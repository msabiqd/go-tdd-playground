package service

import (
	"errors"
	"go-tdd-playground/internal/user/repository"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserService_CreateEmployee(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepositoryInterface
	}
	type args struct {
		req CreateEmployeeRequest
	}

	type test struct {
		name     string
		fields   fields
		args     args
		mock     func(test, *gomock.Controller) test
		wantResp CreateEmployeeResponse
		wantErr  error
	}

	employeeEmail := faker.Email()
	employeeId := faker.UUIDHyphenated()
	families := []FamilyData{
		{
			ID:         faker.UUIDHyphenated(),
			EmployeeID: employeeId,
			FirstName:  faker.FirstName(),
			LastName:   faker.LastName(),
			Relation:   faker.Word(),
		},
	}
	tests := []test{
		{
			name:   "Should success insert employee with families",
			fields: fields{},
			args: args{
				req: CreateEmployeeRequest{
					Email:    employeeEmail,
					Password: faker.Password(),
					Families: []CreateFamiliesRequest{
						{
							FirstName: families[0].FirstName,
							LastName:  families[0].LastName,
							Relation:  families[0].Relation,
						},
					},
				},
			},
			mock: func(tt test, ctrl *gomock.Controller) test {
				userRepository := repository.NewMockUserRepositoryInterface(ctrl)

				userRepository.EXPECT().InsertEmployee(repository.InsertEmployeeRequest{
					Email:    employeeEmail,
					Password: tt.args.req.Password,
				}).Times(1).Return(repository.InsertEmployeeResponse{
					ID:    employeeId,
					Email: employeeEmail,
				}, nil)

				userRepository.EXPECT().InsertFamilies(repository.InsertFamiliesDataRequest{
					Families: []repository.InsertFamilyDataRequest{
						{
							EmployeeID: employeeId,
							FirstName:  families[0].FirstName,
							LastName:   families[0].LastName,
							Relation:   families[0].Relation,
						},
					},
				}).Times(1).Return(repository.InsertFamiliesDataResponse{
					Families: []repository.FamilyData{
						{
							ID:         families[0].ID,
							EmployeeID: employeeId,
							FirstName:  families[0].FirstName,
							LastName:   families[0].LastName,
							Relation:   families[0].Relation,
						},
					},
				}, nil)

				tt.fields.UserRepository = userRepository

				return tt
			},
			wantResp: CreateEmployeeResponse{
				ID:    employeeId,
				Email: employeeEmail,
			},
			wantErr: nil,
		},
		{
			name:   "Should success insert employee with empty families",
			fields: fields{},
			args: args{
				req: CreateEmployeeRequest{
					Email:    employeeEmail,
					Password: faker.Password(),
					Families: []CreateFamiliesRequest{},
				},
			},
			mock: func(tt test, ctrl *gomock.Controller) test {
				userRepository := repository.NewMockUserRepositoryInterface(ctrl)

				userRepository.EXPECT().InsertEmployee(repository.InsertEmployeeRequest{
					Email:    employeeEmail,
					Password: tt.args.req.Password,
				}).Times(1).Return(repository.InsertEmployeeResponse{
					ID:    employeeId,
					Email: employeeEmail,
				}, nil)

				tt.fields.UserRepository = userRepository

				return tt
			},
			wantResp: CreateEmployeeResponse{
				ID:    employeeId,
				Email: employeeEmail,
			},
			wantErr: nil,
		},
		{
			name:   "Should return error when repositor return rror",
			fields: fields{},
			args: args{
				req: CreateEmployeeRequest{
					Email:    employeeEmail,
					Password: faker.Password(),
				},
			},
			mock: func(tt test, ctrl *gomock.Controller) test {
				userRepository := repository.NewMockUserRepositoryInterface(ctrl)
				userRepository.EXPECT().InsertEmployee(repository.InsertEmployeeRequest{
					Email:    employeeEmail,
					Password: tt.args.req.Password,
				}).Times(1).Return(
					repository.InsertEmployeeResponse{},
					errors.New("dummy error"))
				tt.fields.UserRepository = userRepository
				return tt
			},
			wantResp: CreateEmployeeResponse{},
			wantErr:  errors.New("dummy error"),
		},
		{
			name:   "Should return error when email is empty",
			fields: fields{},
			args: args{
				req: CreateEmployeeRequest{
					Email:    "",
					Password: faker.Password(),
				},
			},
			mock: func(tt test, ctrl *gomock.Controller) test {
				userRepository := repository.NewMockUserRepositoryInterface(ctrl)
				tt.fields.UserRepository = userRepository
				return tt
			},
			wantResp: CreateEmployeeResponse{},
			wantErr:  ErrInvalidParam,
		},
		{
			name:   "Should return error when password is empty",
			fields: fields{},
			args: args{
				req: CreateEmployeeRequest{
					Email:    faker.Email(),
					Password: "",
				},
			},
			mock: func(tt test, ctrl *gomock.Controller) test {
				userRepository := repository.NewMockUserRepositoryInterface(ctrl)
				tt.fields.UserRepository = userRepository
				return tt
			},
			wantResp: CreateEmployeeResponse{},
			wantErr:  ErrInvalidParam,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()
			tt = tt.mock(tt, mockController)
			s := NewUserService(tt.fields.UserRepository)
			gotResp, err := s.CreateEmployee(tt.args.req)
			require.Equal(t, tt.wantResp, gotResp)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.Equal(t, err, tt.wantErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
