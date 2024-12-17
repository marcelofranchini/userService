package domain

type UserCreateRepositoryInterface interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
}
