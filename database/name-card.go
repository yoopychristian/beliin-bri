package database

import (
	"gorm.io/gorm"
)

type NameCard struct {
	IDKartuNama string `gorm:"column:id_kartunama;type:varchar(25)"`
	IDUser      string `gorm:"column:id_user;type:varchar(25)"`
	NamaToko    string `gorm:"column:nama_toko;type:varchar(25)"`
	BidangUsaha string `gorm:"column:bidang_usaha;type:varchar(25)"`
	Alamat      string `gorm:"column:alamat_toko;type:text"`
	NoTelepon   string `gorm:"column:no_telepon;type:varchar;size:20"`
	IsDeleted   bool   `gorm:"column:is_deleted;type:bool"`
}

func (p *NameCard) Create(db *gorm.DB, idNameCard, idUser, namaToko, bidangUsaha, alamat, noTelepon string, isdeleted bool) error {
	p.IDKartuNama = idNameCard
	p.IDUser = idUser
	p.NamaToko = namaToko
	p.BidangUsaha = bidangUsaha
	p.Alamat = alamat
	p.NoTelepon = noTelepon
	p.IsDeleted = isdeleted

	return db.Table("kartu_nama").Create(&p).Error
}

func (p NameCard) NameCardList(db *gorm.DB) ([]NameCard, error) {
	NameCards := []NameCard{}
	err := db.Table("kartu_nama").Where("is_deleted=true").Find(&NameCards).Error
	return NameCards, err
}

func (p *NameCard) GetByID(db *gorm.DB, id_NameCard string) ([]NameCard, error) {
	NameCards := []NameCard{}
	err := db.Table("kartu_nama").Where("id_kartunama=?", id_NameCard).Find(&NameCards).Error
	return NameCards, err
}

func (p *NameCard) UpdateNameCard(db *gorm.DB, idNameCard, idUser, namaToko, bidangUsaha, alamat, noTelepon string) error {
	sql := "update kartu_nama set id_user=?, nama_toko=?, bidang_usaha=?, alamat_toko=?, no_telepon=? where id_kartunama=?"
	if err := db.Table("kartu_nama").Exec(sql, idUser, namaToko, bidangUsaha, alamat, noTelepon, idNameCard).Error; err != nil {
		return err
	}

	return nil
}

func (p *NameCard) DeleteNameCard(db *gorm.DB, id_NameCard string) error {
	return db.Table("kartu_nama").Where("id_kartunama=?", id_NameCard).Delete(&p).Error
}
