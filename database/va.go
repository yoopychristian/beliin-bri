package database

import (
	"time"

	"gorm.io/gorm"
)

type VirtualAccount struct {
	IDVA         string    `gorm:"column:id_va;type:varchar(25)"`
	IDPelanggan  string    `gorm:"column:id_pelanggan;type:varchar(25)"`
	KodeBank     string    `gorm:"column:kode_bank;type:varchar(10)"`
	NamaBank     string    `gorm:"column:nama_bank;type:varchar(30)"`
	NoRekening   string    `gorm:"column:no_rekening;type:varchar(25)"`
	NamaRekening string    `gorm:"column:nama_rekening;type:varchar(25)"`
	NoVa         string    `gorm:"column:no_va;type:varchar(25)"`
	Currency     string    `gorm:"column:currency;type:varchar(25)"`
	CreatedDate  time.Time `gorm:"column:created_datetime"`
	UpdatedDate  time.Time `gorm:"column:updated_datetime"`
	VaStatus     string    `gorm:"column:va_status;type:varchar(25)"`
	IsDeleted    bool      `gorm:"column:is_deleted;type:bool"`
}

func (p *VirtualAccount) Create(db *gorm.DB, idVa, idPelanggan, namaRekening, noVa, currency, vaStatus string, createdDatetime time.Time, isdeleted bool) error {
	p.IDVA = idVa
	p.IDPelanggan = idPelanggan
	p.NoVa = noVa
	p.Currency = currency
	p.VaStatus = vaStatus
	p.CreatedDate = createdDatetime
	p.IsDeleted = isdeleted
	return db.Table("virtual_account").Create(&p).Error
}
