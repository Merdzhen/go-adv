package auth_test

import (
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"testing"
)

type MockUserRepository struct {}

func (repo *MockUserRepository) Create(*user.User) (*user.User, error) {
	return nil, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "a@a.ru"
	authservice := auth.NewAuthService(&MockUserRepository{})
	email, err := authservice.Register(initialEmail, "1", "Vasiliy")
	if err != nil {
		t.Fatal(err)
		return
	}
	if email != initialEmail {
		t.Fatalf("Email %s do not match %s", email, initialEmail)
		return
	}
}
