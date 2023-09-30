package photo

import "gorm.io/gorm"

type Repository interface {
	Save(photo Photo) (Photo, error)
	Update(photo Photo) (Photo, error)
	FindByID(ID int) (Photo, error)
	FindAll() ([]Photo, error)
	Delete(ID int) (Photo, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Photo, error) {
	var photos []Photo
	
	err := r.db.Preload("User").Find(&photos).Error
	if err != nil {
		return photos, err
	}

	return photos, nil
}

func (r *repository) Update(photo Photo) (Photo, error) {

	err := r.db.Save(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) Save(photo Photo) (Photo, error) {

	err := r.db.Create(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil
}

func (r *repository) FindByID(ID int) (Photo, error) {
	var photo Photo

	err := r.db.Where("ID = ?", ID).Find(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil

}

func (r *repository) Delete(ID int) (Photo, error) {
	var photo Photo

	err := r.db.Where("ID = ?", ID).Delete(&photo).Error
	if err != nil {
		return photo, err
	}

	return photo, nil

}	