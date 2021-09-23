package database

import (
	"errors"

	"gorm.io/gorm"
)

type Registration struct {
	IDUser    string `gorm:"column:id_user;primaryKey"`
	Nama      string `gorm:"column:nama;type:varchar;size:25"`
	Username  string `gorm:"column:username;type:varchar;size:25"`
	Password  string `gorm:"column:password;type:varchar;size:255"`
	Email     string `gorm:"column:email;type:varchar;size:20"`
	NoTelepon string `gorm:"column:no_telepon;type:varchar;size:20"`
	NoKTP     string `gorm:"column:no_ktp;type:varchar;size:20"`
	IsDeleted bool   `gorm:"column:is_deleted;type:bool"`
}

func (p *Registration) Create(db *gorm.DB, idUser, nama, username, password, email, noTelepon, noKTP string, isdeleted bool) error {
	p.IDUser = idUser
	p.Nama = nama
	p.Username = username
	p.Password = password
	p.Email = email
	p.NoTelepon = noTelepon
	p.NoKTP = noKTP
	p.IsDeleted = isdeleted

	return db.Table("user_registration").Create(&p).Error
}

func (p Registration) CheckAvailability(db *gorm.DB, username, email, noPonsel string) bool {
	err := db.Table("user_registration").Where("username=? or email=? or no_ponsel=?", username, email, noPonsel).Last(&p).Error
	return err == nil
}

func (p Registration) AccountExist(db *gorm.DB, username, email, noPonsel string) (bool, error) {
	result := db.Table("user_registration").Where("username=? or email=? or no_ponsel=?", username, email, noPonsel).Take(&username, &email, &noPonsel)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}

func (p Registration) EmailExist(db *gorm.DB, email string) (bool, error) {
	result := db.Table("user_registration").Where("email=?", email).Take(&email)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}

func (p Registration) UsernameExist(db *gorm.DB, username string) (bool, error) {
	result := db.Table("user_registration").Where("username=?", username).Take(&username)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}
