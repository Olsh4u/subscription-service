package data

type UserInterface interface {
	All() ([]*User, error)
	ByEmail(email string) (*User, error)
	ById(id int) (*User, error)
	Update(user User) error
	Delete() error
	DeleteByID(id int) error
	Insert(user User) (int, error)
	ResetPassword(password string) error
	PasswordMatches(plainText string) (bool, error)
}

type PlanInterface interface {
	All() ([]*Plan, error)
	ById(id int) (*Plan, error)
	SubscribeUserToPlan(user User, plan Plan) error
	AmountForDisplay() string
}
