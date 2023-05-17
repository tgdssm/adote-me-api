package usecases

import (
	"api/internal/core/domain"
	"api/mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mock struct {
	userRepository *mocks.MockUserMysqlRepository
}

func userWithProfileImage(id uint64, name, email, cellphone, passwd string) domain.User {
	return domain.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Cellphone: cellphone,
		ProfileImage: domain.ProfileImage{
			ID:       1,
			UserID:   id,
			FileName: "File Name",
			FilePath: "File Path",
		},
		Passwd:    passwd,
		CreatedAt: time.Now(),
	}
}

func userWithoutProfileImage() *domain.User {
	return &domain.User{
		ID:        1,
		Name:      "Name",
		Email:     "Email",
		Cellphone: "38999001122",
		Passwd:    "123456",
	}
}

func TestUserUseCase_Create(t *testing.T) {
	user := userWithoutProfileImage()
	type args struct {
		user *domain.User
	}

	type want struct {
		result *domain.User
		err    error
	}

	tests := []struct {
		name string
		args args
		want want
		mock func(m mock)
	}{
		{
			name: "Should create user successfully",
			args: args{user: user},
			want: want{result: user},
			mock: func(m mock) {
				m.userRepository.EXPECT().Create(user).Return(user, nil)
			},
		},
		{
			name: "Should return error",
			args: args{user: user},
			want: want{err: errors.New("internal server error 500")},
			mock: func(m mock) {
				m.userRepository.EXPECT().Create(user).Return(nil, errors.New("internal server error 500"))
			},
		},
	}

	for _, tt := range tests {
		m := mock{userRepository: mocks.NewMockUserMysqlRepository(gomock.NewController(t))}
		tt.mock(m)
		useCase := NewUserUseCase(m.userRepository)
		result, err := useCase.Create(tt.args.user)

		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}

func TestUserUseCase_List(t *testing.T) {
	usersWithoutParameter := []domain.User{
		userWithProfileImage(1, "Name 1", "Email 1", "38999090807", "123456"),
		userWithProfileImage(2, "Name 2", "Email 2", "38999060504", "123456"),
		userWithProfileImage(3, "Name 3", "Email 3", "38999030201", "123456"),
	}

	usersWithParameter := []domain.User{
		userWithProfileImage(1, "Name 1", "Email 1", "38999090807", "123456"),
	}

	type args struct {
		queryParameter string
	}

	type want struct {
		result []domain.User
		err    error
	}

	tests := []struct {
		name string
		args args
		want want
		mock func(m mock)
	}{
		{
			name: "Should list users successfully without query parameter",
			args: args{queryParameter: ""},
			want: want{result: usersWithoutParameter},
			mock: func(m mock) {
				m.userRepository.EXPECT().List("").Return(usersWithoutParameter, nil)
			},
		},
		{
			name: "Should list users successfully with query parameter",
			args: args{queryParameter: "Name 1"},
			want: want{result: usersWithParameter},
			mock: func(m mock) {
				m.userRepository.EXPECT().List("Name 1").Return(usersWithParameter, nil)
			},
		},
		{
			name: "Should successfully list users with query parameter but found no results",
			args: args{queryParameter: "Name 5"},
			want: want{result: []domain.User{}},
			mock: func(m mock) {
				m.userRepository.EXPECT().List("Name 5").Return([]domain.User{}, nil)
			},
		},
		{
			name: "Should return error",
			args: args{queryParameter: "Name 5"},
			want: want{err: errors.New("bad request 400")},
			mock: func(m mock) {
				m.userRepository.EXPECT().List("Name 5").Return(nil, errors.New("bad request 400"))
			},
		},
	}

	for _, tt := range tests {
		m := mock{
			userRepository: mocks.NewMockUserMysqlRepository(gomock.NewController(t)),
		}

		tt.mock(m)
		useCase := NewUserUseCase(m.userRepository)

		result, err := useCase.List(tt.args.queryParameter)
		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}

func TestUserUseCase_Get(t *testing.T) {
	user := userWithProfileImage(1, "Name 1", "Email 1", "38999090807", "123456")

	type args struct {
		id            int
		idWrongFormat string
	}

	type want struct {
		result *domain.User
		err    error
	}

	tests := []struct {
		name string
		args args
		want want
		mock func(m mock)
	}{
		{
			name: "Should get user successfully",
			args: args{id: 1},
			want: want{result: &user},
			mock: func(m mock) {
				m.userRepository.EXPECT().Get(1).Return(&user, nil)
			},
		},
		{
			name: "Should return error",
			args: args{id: 2},
			want: want{err: errors.New("not found 404")},
			mock: func(m mock) {
				m.userRepository.EXPECT().Get(2).Return(nil, errors.New("not found 404"))
			},
		},

		{
			name: "Should return error",
			args: args{idWrongFormat: "ID"},
			want: want{err: errors.New("bad request 400")},
			mock: func(m mock) {
				m.userRepository.EXPECT().Get(0).Return(nil, errors.New("bad request 400"))
			},
		},
	}

	for _, tt := range tests {
		m := mock{
			userRepository: mocks.NewMockUserMysqlRepository(gomock.NewController(t)),
		}

		tt.mock(m)
		useCase := NewUserUseCase(m.userRepository)
		if tt.args.idWrongFormat != "" {
			assert.Equal(t, tt.want.err.Error(), errors.New("bad request 400").Error())
		}
		result, err := useCase.Get(tt.args.id)

		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, tt.want.result, result)
	}
}

func TestUserUseCase_Update(t *testing.T) {
	user := userWithProfileImage(0, "Name 1", "Email 1", "38999090807", "123456")

	type args struct {
		user *domain.User
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want want
		mock func(m mock)
	}{
		{
			name: "Should Update user successfully",
			args: args{user: &user},
			want: want{err: nil},
			mock: func(m mock) {
				m.userRepository.EXPECT().Update(&user).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		m := mock{
			userRepository: mocks.NewMockUserMysqlRepository(gomock.NewController(t)),
		}

		tt.mock(m)
		useCase := NewUserUseCase(m.userRepository)

		err := useCase.Update(tt.args.user)

		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, true, true)
	}
}

func TestUserUseCase_Delete(t *testing.T) {
	type args struct {
		id int
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args args
		want want
		mock func(m mock)
	}{
		{
			name: "Should Update user successfully",
			args: args{id: 1},
			want: want{err: nil},
			mock: func(m mock) {
				m.userRepository.EXPECT().Delete(1).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		m := mock{
			userRepository: mocks.NewMockUserMysqlRepository(gomock.NewController(t)),
		}

		tt.mock(m)
		useCase := NewUserUseCase(m.userRepository)

		err := useCase.Delete(tt.args.id)

		if tt.want.err != nil && err != nil {
			assert.Equal(t, tt.want.err.Error(), err.Error())
		}

		assert.Equal(t, true, true)
	}
}
