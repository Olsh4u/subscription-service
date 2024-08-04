package data

import (
	"database/sql"
	"fmt"
	"time"
)

func TestNew(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: &UserTest{},
		Plan: &PlanTest{},
	}
}

type UserTest struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Password  string
	Active    int
	IsAdmin   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan      *Plan
}

type PlanTest struct {
	ID                  int
	PlanName            string
	PlanAmount          int
	PlanAmountFormatted string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (u *UserTest) All() ([]*User, error) {
	var users []*User

	user := User{
		ID:        1,
		Email:     "gosling@rayan.com",
		FirstName: "Rayan",
		LastName:  "Gosling",
		Password:  "12345",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	users = append(users, &user)

	return users, nil
}

func (u *UserTest) ByEmail(email string) (*User, error) {
	user := User{
		ID:        1,
		Email:     "gosling@rayan.com",
		FirstName: "Rayan",
		LastName:  "Gosling",
		Password:  "12345",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	return &user, nil
}

func (u *UserTest) ById(id int) (*User, error) {
	return u.ByEmail("")
}

func (u *UserTest) Update(user User) error {
	return nil
}

func (u *UserTest) Delete() error {
	return nil
}

func (u *UserTest) DeleteByID(id int) error {
	return nil
}

func (u *UserTest) Insert(user User) (int, error) {
	return 2, nil
}

func (u *UserTest) ResetPassword(password string) error {
	return nil
}

func (u *UserTest) PasswordMatches(plainText string) (bool, error) {
	return true, nil
}

func (p *PlanTest) All() ([]*Plan, error) {
	var plans []*Plan

	plan := Plan{
		ID:         1,
		PlanName:   "Best Plan",
		PlanAmount: 1000,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}

	plans = append(plans, &plan)

	return plans, nil
}

func (p *PlanTest) ById(id int) (*Plan, error) {
	plan := Plan{
		ID:         1,
		PlanName:   "Best Plan",
		PlanAmount: 1000,
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}

	return &plan, nil
}

func (p *PlanTest) SubscribeUserToPlan(user User, plan Plan) error {
	return nil
}

func (p *PlanTest) AmountForDisplay() string {
	amount := float64(p.PlanAmount) / 100.0
	return fmt.Sprintf("$%.2f", amount)
}
