package photo

import (
	"errors"
)

type Service interface {
	CreatePhoto(input CreatePhotoInput, fileLocation string) (Photo, error)
	UpdatePhoto(id GetId, input UpdatePhotoInput, fileLocation string) (Photo, error)
	GetPhotos() ([]Photo, error)
	Delete(ID int, UserID int) (Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreatePhoto(input CreatePhotoInput, fileLocation string) (Photo, error) {

	photo := Photo{}
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoURL = fileLocation
	photo.UserID = input.User.ID

	newPhoto, err := s.repository.Save(photo)
	if err != nil {
		return newPhoto, err
	}

	return newPhoto, nil
}

func (s *service) GetPhotos() ([]Photo, error) {

	photos, err := s.repository.FindAll()
	if err != nil {
		return photos, err
	}

	return photos, nil

}

func (s *service) Delete(ID int, UserID int) (Photo, error) {

	photo, err := s.repository.FindByID(ID)
	if err != nil {
		return photo, err
	}

	if photo.UserID  != UserID {
		return photo, errors.New("Not an owner of the user")
	}

	deletedPhoto, err := s.repository.Delete(photo.ID)
	if err != nil {
		return deletedPhoto, err
	}

	return deletedPhoto, nil

}

func (s *service) UpdatePhoto(id GetId, input UpdatePhotoInput, fileLocation string) (Photo, error) {

	photo, err := s.repository.FindByID(id.ID)
	if err != nil {
		return photo, err
	}

	if photo.UserID != input.User.ID {
		return photo, errors.New("Not an owner of the user")
	}

	photo.Caption = input.Caption
	photo.Title = input.Title
	photo.PhotoURL = fileLocation
	
	if err != nil {
		return photo, err
	}

	updatedUser, err := s.repository.Update(photo)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
