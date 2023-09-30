package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginInput) (User, error)
	IsEmailAvailable(email string) (User, error)
	GetUserByID(ID int) (User, error)
	Delete(ID int, UserID int) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(id GetId, input UpdateUserInput) (User, error)

}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash) 

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginInput) (User, error) {

	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email) //cek balikannya dari repository findbyemail

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(email string) (User, error) {

	user, err := s.repository.FindByEmail(email) //cek balikannya dari repository findbyemail
	if err != nil {
		return user, err
	}

	if(user.Email != "") {
		return user, errors.New("User already exist")
	}

	return user, nil

}

func (s *service) GetUserByID(ID int) (User, error) {

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	return user, nil

}

func (s *service) Delete(ID int, UserID int) (User, error) {

	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID  != UserID {
		return user, errors.New("Not an owner of the user")
	}

	deletedUser, err := s.repository.Delete(user.ID)
	if err != nil {
		return deletedUser, err
	}

	return deletedUser, nil

}

func (s *service) GetAllUsers() ([]User, error) {
	
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil

}

func (s *service) UpdateUser(id GetId, input UpdateUserInput) (User, error) {

	user, err := s.repository.FindByID(id.ID)
	if err != nil {
		return user, err
	}

	
	if (user.ID != input.UserLogin.ID) {
		return user, errors.New("Not an owner of the user")
	}

	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash) 

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}