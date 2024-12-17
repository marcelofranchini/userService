package domain

type FirebaseRepositoryInterface interface {
	CreateUser(user *UserCreateRequest) error
}
