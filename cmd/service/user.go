package service

import (
	"context"
	"fmt"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *User) error
	GetUsers(ctx context.Context) ([]User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, user *User) error
	GetUsers(ctx context.Context) ([]User, error)
}

func NewUserService(
	userStore UserStore,
) UserService {
	return &userService{
		userStore: userStore,
	}
}

type userService struct {
	userStore UserStore
}

type User struct {
	ID         string     `json:"id"`
	FirstName  string     `json:"first_name"`
	MiddleName *string    `json:"middle_name"`
	LastName   string     `json:"last_name"`
	Type       UserType   `json:"type"`
	Status     UserStatus `json:"status"`
}

type (
	UserType string

	UserStatus string
)

var (
	UserTypeDriver   UserType = "driver"
	UserTypeCustomer UserType = "customer"

	UserTypes = []UserType{
		UserTypeDriver,
		UserTypeCustomer,
	}

	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"

	UserStatuses = []UserStatus{
		UserStatusActive,
		UserStatusInactive,
	}
)

func (v *UserType) UnmarshalText(text []byte) error {
	for _, val := range UserTypes {
		if string(text) == string(val) {
			*v = val
			return nil
		}
	}
	return fmt.Errorf("invalid UserType: %s", string(text))
}

func (v *UserStatus) UnmarshalText(text []byte) error {
	for _, val := range UserStatuses {
		if string(text) == string(val) {
			*v = val
			return nil
		}
	}
	return fmt.Errorf("invalid UserStatus: %s", string(text))
}

func (s *userService) CreateUser(ctx context.Context, user *User) error {
	// suppose to be validation / business logic here
	return s.userStore.CreateUser(ctx, user)
}

func (s *userService) GetUsers(ctx context.Context) ([]User, error) {
	return s.userStore.GetUsers(ctx)
}
