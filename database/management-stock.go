package database

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	IDStock      string    `gorm:"column:id_stock;type:varchar(25)"`
	IDUser       string    `gorm:"column:id_user;type:varchar(25)"`
	NamaBarang   string    `gorm:"column:nama_barang;type:varchar(20)"`
	Deskripsi    string    `gorm:"column:deskripsi_barang;type:text"`
	JumlahBarang int       `gorm:"column:jumlah_barang;type:int"`
	HargaSatuan  string    `gorm:"column:harga_barangsatuan;type:varchar(20)"`
	GambarBarang string    `gorm:"column:gambar_barang;type:text"`
	CreatedDate  time.Time `gorm:"column:created_datetime"`
	UpdatedDate  time.Time `gorm:"column:updated_datetime"`
	IsDeleted    bool      `gorm:"column:is_deleted;type:bool"`
}

func (p *Stock) Create(db *gorm.DB, idStock, idUser, namaBarang, deskripsi, hargaSatuan, gambarBarang string, jumlahBarang int, createdDate time.Time, isdeleted bool) error {
	p.IDStock = idStock
	p.IDUser = idUser
	p.NamaBarang = namaBarang
	p.Deskripsi = deskripsi
	p.JumlahBarang = jumlahBarang
	p.HargaSatuan = hargaSatuan
	p.GambarBarang = gambarBarang
	p.CreatedDate = createdDate
	p.IsDeleted = isdeleted

	return db.Table("stock_management").Create(&p).Error
}

func (p Stock) StockList(db *gorm.DB) ([]Stock, error) {
	stocks := []Stock{}
	err := db.Table("stock_management").Where("is_deleted=true").Find(&stocks).Error
	return stocks, err
}

func (p *Stock) GetByID(db *gorm.DB, id_stock string) error {
	return db.Table("stock_management").Where("id_stock=?", id_stock).Last(&p).Error
}

func (p *Stock) UpdateStock(db *gorm.DB, idStock, namaBarang, deskripsi string, jumlahBarang int, hargaSatuan, gambarBarang string, updated_datetime time.Time) error {
	sql := "update stock_management set nama_barang=?, deskripsi_barang=?, jumlah_barang=?, harga_barangsatuan=?, gambar_barang=?, updated_datetime=? where id_stock=?"
	if err := db.Table("stock_management").Exec(sql, namaBarang, deskripsi, jumlahBarang, hargaSatuan, gambarBarang, updated_datetime, idStock).Error; err != nil {
		return err
	}

	return nil
}

func (p *Stock) DeleteStock(db *gorm.DB, id_stock string) error {
	return db.Table("stock_management").Where("id_stock=?", id_stock).Delete(&p).Error
}
