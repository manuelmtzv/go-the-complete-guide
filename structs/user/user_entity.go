package user

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	User
	Email          string
	HashedPassword string
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" {
		return nil, fmt.Errorf("first name is required")
	}

	if lastName == "" {
		return nil, fmt.Errorf("last name is required")
	}

	if birthDate == "" {
		return nil, fmt.Errorf("birth date is required")
	}

	return &User{
		firstName,
		lastName,
		birthDate,
		time.Now(),
	}, nil
}

func NewAdmin(firstName, lastName, birthDate, email, hashedPassword string) (*Admin, error) {
	user, error := New(firstName, lastName, birthDate)

	if error != nil {
		return nil, error
	}

	if email == "" {
		return nil, fmt.Errorf("email is required")
	}

	if hashedPassword == "" {
		return nil, fmt.Errorf("hashed password is required")
	}

	return &Admin{
		*user,
		email,
		hashedPassword,
	}, nil
}

func (user *User) ToString() string {
	return fmt.Sprintf("First Name: %s\nLast Name: %s\nBirth Date: %s\nCreated At: %s\n",
		user.FirstName, user.LastName, user.BirthDate, user.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (user *User) Clear() {
	user.FirstName = ""
	user.LastName = ""
	user.BirthDate = ""
	user.CreatedAt = time.Time{}
}
