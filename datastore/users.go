package datastore

import (
	"firebase-go/domain/model"
)

func (d *DB) FindUserByUUID(uuid string) (*model.User, error) {
	var user model.User
	err:= d.Where(model.User{UUID:uuid}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *DB) CreateUser(user interface{}) error {
	err := d.Create(user).Error

	if err != nil {
		return err
	}
	return nil
}