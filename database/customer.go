package database

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	IDPelanggan      string    `gorm:"column:id_pelanggan;type:varchar(25)"`
	IDUser           string    `gorm:"column:id_user;type:varchar(25)"`
	IDVa             string    `gorm:"column:id_va;type:varchar(25)"`
	Nama             string    `gorm:"column:nama;type:varchar(20)"`
	Email            string    `gorm:"column:email;type:varchar(20)"`
	AlamatPengiriman string    `gorm:"column:alamat_pengiriman;type:text"`
	Kota             string    `gorm:"column:kota;type:varchar(20)"`
	NoTelepon        string    `gorm:"no_telepon;type:varchar(15)"`
	CreatedDate      time.Time `gorm:"column:created_datetime"`
	UpdatedDate      time.Time `gorm:"column:updated_datetime"`
	IsDeleted        bool      `gorm:"column:is_deleted;type:bool"`
}

func (p *Customer) Create(db *gorm.DB, idPelanggan, idUser, idVa, nama, email, alamat, kota, noTelepon string, createdDate time.Time, isdeleted bool) error {
	p.IDPelanggan = idPelanggan
	p.IDUser = idUser
	p.IDVa = idVa
	p.Nama = nama
	p.Email = email
	p.AlamatPengiriman = alamat
	p.Kota = kota
	p.NoTelepon = noTelepon
	p.CreatedDate = createdDate
	p.IsDeleted = isdeleted

	return db.Table("daftar_pelanggan").Create(&p).Error
}

func (p Customer) CustomerList(db *gorm.DB) ([]Customer, error) {
	customers := []Customer{}
	err := db.Table("daftar_pelanggan").Where("is_deleted=true").Find(&customers).Error
	return customers, err
}

func (p *Customer) GetByID(db *gorm.DB, id_pelanggan string) error {
	return db.Table("daftar_pelanggan").Where("id_pelanggan=?", id_pelanggan).Last(&p).Error
}

func (p *Customer) UpdateCustomer(db *gorm.DB, idPelanggan, nama, email, alamat, kota, noTelepon string, updated_datetime time.Time) error {
	sql := "update daftar_pelanggan set nama=?, email=?, alamat_pengiriman=?, kota=?, no_telepon=?, updated_datetime=? where id_pelanggan=?"
	if err := db.Table("daftar_pelanggan").Exec(sql, nama, email, alamat, kota, noTelepon, updated_datetime, idPelanggan).Error; err != nil {
		return err
	}

	return nil
}

func (p *Customer) DeleteCustomer(db *gorm.DB, id_stock string) error {
	return db.Table("daftar_pelanggan").Where("id_pelanggan=?", id_stock).Delete(&p).Error
}
