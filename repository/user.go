package repository

import "github.com/ehardi19/bitment-backend/domain"

func (r *Repository) GetAllUser() ([]domain.User, error) {
	var user []domain.User

	if err := r.DB.Model(&user).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) CreateUser(user domain.User) (domain.User, error) {
	if err := r.DB.Model(&user).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) GetUserByID(id int64) (domain.User, error) {
	var user domain.User

	if err := r.DB.Model(&user).Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User

	if err := r.DB.Model(&user).Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *Repository) UpdateUser(user domain.User) error {
	if err := r.DB.Model(&user).Where("id = ?", user.ID).Update(&user).Error; err != nil {
		return err
	}

	return nil
}
