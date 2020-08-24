package service

import (
	"github.com/nofendian17/logAudit/entity"
	"github.com/nofendian17/logAudit/repository"
)

type UserService interface {
	Save(user entity.User) entity.User
	FindAll() []entity.User
	Update(user entity.User)
	Delete(user entity.User)
}

type userService struct {
	userRepository repository.UserRepository
}

func New(repository repository.UserRepository) UserService {
	return &userService{
		userRepository: repository,
	}
}

func (service *userService) Save(user entity.User) entity.User {
	service.userRepository.Save(user)
	return user
}

func (service *userService) Update(user entity.User) {
	service.userRepository.Update(user)
}

func (service *userService) Delete(user entity.User) {
	service.userRepository.Delete(user)
}

func (service *userService) FindAll() []entity.User {
	return service.userRepository.FindAll()
}