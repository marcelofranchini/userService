package useCase

import "userService/internal/user/domain"

type CreateUserUseCase struct {
	userRepository               domain.UserCreateRepositoryInterface
	userCreateFirebaseRepository domain.FirebaseRepositoryInterface
}

func NewCreateUser(userRepository domain.UserCreateRepositoryInterface, userCreateFirebaseRepository domain.FirebaseRepositoryInterface) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository:               userRepository,
		userCreateFirebaseRepository: userCreateFirebaseRepository,
	}
}

func (uc *CreateUserUseCase) Execute(user domain.User, userFirebase domain.UserCreateRequest) error {
	err := uc.userRepository.Save(&user)
	if err != nil {
		return err
	}

	err = uc.userCreateFirebaseRepository.CreateUser(&userFirebase)
	if err != nil {
		return err
	}

	return nil
}
