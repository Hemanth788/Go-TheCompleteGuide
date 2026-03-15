package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

// strict embedding
type Admin struct {
	email string
	password string
	User
}

// receiver, turns func into method of the received type
// pass by value, not reference, copy
func (
	// user *User
	user User) OutputUserDetails(
	// userAddr User - need not pass if converted to method
) {
	// to use this, the receiver declaration should be switched to a pointer
	// fmt.Println((*userAddr).firstName, (*userAddr).lastName, (*userAddr).birthDate)
	// pointers to structs are exempted from de-referencing
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

// mutation method, ref/pointer/addr in the receiver declaration, no copy
func (userAddr *User) ClearUserName() {
	userAddr.firstName = ""
	userAddr.lastName = ""
}

// constructor
func New(firstNameInput string, lastNameInput string, birthDateInput string) (*User, error) {
	if firstNameInput == "" || lastNameInput == "" || birthDateInput == "" {
		return nil, errors.New("All inputs are required")
	}
	return &User{
		firstName: firstNameInput,
		lastName: lastNameInput,
		birthDate: birthDateInput,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email string, password string) *Admin {
	return &Admin{
		email: email,
		password: password,
		// this should be the case, can't have User properties at this level even with anonymous embedding
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "N/A",
			createdAt: time.Now(),
		},
	}
}