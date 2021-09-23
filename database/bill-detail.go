package database

import (
	"beliin-bri/shared"

	"gorm.io/gorm"
)

type BillDetail struct {
	IDPelanggan       string `gorm:"column:id_pelanggan;type:string"`
	Nama              string `gorm:"column:nama;type:string"`
	Alamat            string `gorm:"column:alamat_pengiriman;type:string"`
	NoTelepon         string `gorm:"column:no_telepon;type:string"`
	JumlahBarang      int    `gorm:"column:jumlah_barang;type:int"`
	TotalHarga        int    `gorm:"column:total_harga;type:int"`
	Email             string `gorm:"column:email;type:string"`
	PilihanPengiriman string `gorm:"column:pilihan_pengiriman;type:string"`
	NoVa              string `gorm:"column:no_va;type:string"`
}

func (p BillDetail) DetailList(db *gorm.DB, id string) (*[]shared.BillDetail, error) {
	sql := `SELECT dp.id_pelanggan, dp.nama, dp.alamat_pengiriman, dp.no_telepon, oc.jumlah_barang, oc.total_harga, oc.pilihan_pengiriman, va.no_va FROM daftar_pelanggan AS dp FULL JOIN order_customers AS oc ON oc.id_pelanggan = dp.id_pelanggan FULL JOIN virtual_account AS va ON va.id_va = dp.id_va where va.id_pelanggan = ?`

	rows, err := db.Raw(sql, id).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []shared.BillDetail{}

	for rows.Next() {
		g := shared.BillDetail{}
		rows.Scan(
			&g.IDPelanggan,
			&g.Nama,
			&g.Alamat,
			&g.NoTelepon,
			&g.JumlahBarang,
			&g.TotalHarga,
			&g.PilihanPengiriman,
			&g.NoVa)
		data = append(data, g)
	}

	return &data, nil
}

func (p BillDetail) BillList(db *gorm.DB, order string) (*[]shared.BillDetail, error) {
	sql := `SELECT dp.id_pelanggan, dp.nama, dp.alamat_pengiriman, dp.no_telepon, oc.jumlah_barang, oc.total_harga, oc.pilihan_pengiriman, va.no_va FROM daftar_pelanggan AS dp FULL JOIN order_customers AS oc ON oc.id_pelanggan = dp.id_pelanggan FULL JOIN virtual_account AS va ON va.id_va = dp.id_va where oc.order_status = ?`

	rows, err := db.Raw(sql, order).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []shared.BillDetail{}

	for rows.Next() {
		g := shared.BillDetail{}
		rows.Scan(
			&g.IDPelanggan,
			&g.Nama,
			&g.Alamat,
			&g.NoTelepon,
			&g.JumlahBarang,
			&g.TotalHarga,
			&g.PilihanPengiriman,
			&g.NoVa)
		data = append(data, g)
	}

	return &data, nil
}

func (p BillDetail) SendBill(db *gorm.DB, id string) (*[]shared.BillDetail, error) {
	sql := `SELECT dp.id_pelanggan, dp.nama, dp.alamat_pengiriman, dp.no_telepon, oc.jumlah_barang, oc.total_harga, oc.pilihan_pengiriman, dp.email, dp.id_va FROM daftar_pelanggan AS dp
	FULL JOIN order_customers AS oc ON oc.id_pelanggan = dp.id_pelanggan
	FULL JOIN virtual_account AS va ON va.id_va = dp.id_va where dp.id_pelanggan = ?`

	rows, err := db.Raw(sql, id).Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []shared.BillDetail{}

	for rows.Next() {
		g := shared.BillDetail{}
		rows.Scan(
			&g.IDPelanggan,
			&g.Nama,
			&g.Alamat,
			&g.NoTelepon,
			&g.JumlahBarang,
			&g.TotalHarga,
			&g.PilihanPengiriman,
			&g.Email,
			&g.NoVa)
		data = append(data, g)
	}

	return &data, nil
}

func (gt BillDetail) TableName() string {
	return "game_type"
}

func (gt BillDetail) List(db *gorm.DB) ([]BillDetail, error) {
	data := []BillDetail{}
	err := db.Find(&data).Error

	return data, err
}

func (gt BillDetail) OnFront(db *gorm.DB) ([]BillDetail, error) {
	data := []BillDetail{}
	err := db.Where("front=true").Order("ordered").Find(&data).Error

	return data, err
}

func (gt *BillDetail) GetByID(db *gorm.DB, id int) error {
	return db.Where("id=?", id).Last(&gt).Error
}

func (gt *BillDetail) GetByName(db *gorm.DB, name string) error {
	return db.Where("lower(type_name)=?", name).Last(&gt).Error
}
