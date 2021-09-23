package database

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	IDOrder           string    `gorm:"column:id_order;type:varchar(25)"`
	IDUser            string    `gorm:"column:id_user;type:varchar(25)"`
	IDStock           string    `gorm:"column:id_stock;type:varchar(25)"`
	IDPelanggan       string    `gorm:"column:id_pelanggan;type:varchar(25)"`
	JumlahBarang      int       `gorm:"column:jumlah_barang;type:int"`
	TotalHarga        string    `gorm:"column:total_harga;type:varchar(15)"`
	PilihanPengiriman string    `gorm:"column:pilihan_pengiriman;type:varchar(15)"`
	CreatedDate       time.Time `gorm:"column:created_datetime"`
	UpdatedDate       time.Time `gorm:"column:updated_datetime"`
	OrderStatus       string    `gorm:"column:order_status;type:varchar(10)"`
	IsDeleted         bool      `gorm:"column:is_deleted;type:bool"`
}

func (p *Order) Create(db *gorm.DB, idOrder, idUser, idStock, idPelanggan, totalHarga, pilihanPengiriman, orderStatus string, createdDate time.Time, jumlahBarang int, isDeleted bool) error {
	p.IDOrder = idOrder
	p.IDUser = idUser
	p.IDStock = idStock
	p.IDPelanggan = idPelanggan
	p.JumlahBarang = jumlahBarang
	p.TotalHarga = totalHarga
	p.PilihanPengiriman = pilihanPengiriman
	p.CreatedDate = createdDate
	p.OrderStatus = orderStatus
	p.IsDeleted = isDeleted
	return db.Table("order_customers").Create(&p).Error
}

func (p Order) OrderList(db *gorm.DB) ([]Order, error) {
	Orders := []Order{}
	err := db.Table("order_customers").Where("is_deleted=true").Where("order_status='Pesanan Baru'").Find(&Orders).Error
	return Orders, err
}

func (p *Order) GetByID(db *gorm.DB, id_order string) error {
	return db.Table("order_customers").Where("id_order=?", id_order).Last(&p).Error
}

func (p *Order) UpdateOrder(db *gorm.DB, idOrder, totalHarga, pilihanPengiriman string, jumlahBarang int, updateDate time.Time) error {
	sql := "update order_customers set jumlah_barang=?, total_harga=?, pilihan_pengiriman=?, updated_datetime=? where id_order=?"
	if err := db.Table("order_customers").Exec(sql, jumlahBarang, totalHarga, pilihanPengiriman, updateDate, idOrder).Error; err != nil {
		return err
	}

	return nil
}

func (p *Order) DeleteOrder(db *gorm.DB, id_order string) error {
	return db.Table("order_customers").Where("id_order=?", id_order).Delete(&p).Error
}

func (p *Order) StatusOrder(db *gorm.DB, idOrder, statusOrder string) error {
	sql := "update order_customers set order_status=? where id_order=?"
	if err := db.Table("order_customers").Exec(sql, statusOrder, idOrder).Error; err != nil {
		return err
	}

	return nil
}
