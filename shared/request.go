package shared

import "time"

type ParamAccountCheck struct {
	Username string `json:"username" form:"username" url:"username"`
	NoPonsel string `json:"no_ponsel"  form:"no_ponsel" url:"no_ponsel"`
	Email    string `json:"email" form:"email" url:"email"`
}
type ParamAccountEdit struct {
	PlayerID int    `json:"player_id" form:"player_id" url:"player_id"`
	FullName string `json:"full_name" form:"full_name" url:"full_name"`
	Phone    string `json:"phone"  form:"phone" url:"phone"`
	Password string `json:"password" form:"password" url:"password"`
	Email    string `json:"email" form:"email" url:"email"`
}

type ParamAccountRegister struct {
	IDUser    string `json:"id_user" form:"id_user" url:"id_user"`
	Nama      string `json:"nama" form:"nama" url:"nama"`
	Username  string `json:"username" form:"username" url:"username"`
	Password  string `json:"password" form:"password" url:"password"`
	Email     string `json:"email" form:"email" url:"email"`
	NoPonsel  string `json:"no_telepon"  form:"no_telepon" url:"no_telepon"`
	NoKTP     string `json:"no_ktp"  form:"no_ktp" url:"no_ktp"`
	IsDeleted bool   `json:"is_deleted"  form:"is_deleted" url:"is_deleted"`
}

type ParamLogin struct {
	Username string `json:"username" form:"username" url:"username"`
	Password string `json:"password" form:"password" url:"password"`
}

type ParamStock struct {
	IDStock      string    `json:"id_stock" form:"id_stock" url:"id_stock"`
	IDUser       string    `json:"id_user" form:"id_user" url:"id_user"`
	NamaBarang   string    `json:"nama_barang" form:"nama_barang" url:"nama_barang"`
	Deskripsi    string    `json:"deskripsi" form:"deskripsi" url:"deskripsi"`
	JumlahBarang int       `json:"jumlah_barang" form:"jumlah_barang" url:"jumlah_barang"`
	HargaSatuan  string    `json:"harga_barangsatuan" form:"harga_barangsatuan" url:"harga_barangsatuan"`
	GambarBarang string    `json:"gambar_barang" form:"gambar_barang" url:"gambar_barang"`
	CreatedDate  time.Time `json:"created_datetime" form:"created_datetime" url:"created_datetime"`
	UpdatedDate  time.Time `json:"updated_datetime" form:"updated_datetime" url:"updated_datetime"`
	IsDeleted    bool      `json:"is_deleted" form:"is_deleted" url:"is_deleted"`
}

type ParamStockEdit struct {
	IDStock      string `json:"id_stock" form:"id_stock" url:"id_stock"`
	NamaBarang   string `json:"nama_barang" form:"nama_barang" url:"nama_barang"`
	Deskripsi    string `json:"deskripsi" form:"deskripsi" url:"deskripsi"`
	JumlahBarang int    `json:"jumlah_barang" form:"jumlah_barang" url:"jumlah_barang"`
	HargaSatuan  string `json:"harga_barangsatuan" form:"harga_barangsatuan" url:"harga_barangsatuan"`
	GambarBarang string `json:"gambar_barang" form:"gambar_barang" url:"gambar_barang"`
}

type ParamDeleteStock struct {
	IDStock string `json:"id_stock" form:"id_stock" url:"id_stock"`
}

type ParamCustomer struct {
	IDUser           string    `json:"id_user" form:"id_user" url:"id_user"`
	IDVa             string    `json:"id_va" form:"id_va" url:"id_va"`
	Nama             string    `json:"nama" form:"nama" url:"nama"`
	Email            string    `json:"email" form:"email" url:"email"`
	AlamatPengiriman string    `json:"alamat_pengiriman" form:"alamat_pengiriman" url:"alamat_pengiriman"`
	Kota             string    `json:"kota" form:"kota" url:"kota"`
	NoTelepon        string    `json:"no_telepon" form:"no_telepon" url:"no_telepon"`
	CreatedDate      time.Time `json:"created_datetime" form:"created_datetime" url:"created_datetime"`
	UpdatedDate      time.Time `json:"updated_datetime" form:"updated_datetime" url:"updated_datetime"`
	IsDeleted        bool      `json:"is_deleted" form:"is_deleted" url:"is_deleted"`
}

type ParamCustomerEdit struct {
	IDPelanggan      string `json:"id_pelanggan" form:"id_pelanggan" url:"id_pelanggan"`
	Nama             string `json:"nama" form:"nama" url:"nama"`
	Email            string `json:"email" form:"email" url:"email"`
	AlamatPengiriman string `json:"alamat_pengiriman" form:"alamat_pengiriman" url:"alamat_pengiriman"`
	Kota             string `json:"kota" form:"kota" url:"kota"`
	NoTelepon        string `json:"no_telepon" form:"no_telepon" url:"no_telepon"`
}

type ParamDeleteCustomer struct {
	IDPelanggan string `json:"id_pelanggan" form:"id_pelanggan" url:"id_pelanggan"`
}

type ParamOrder struct {
	IDOrder           string    `json:"id_order" form:"id_order" url:"id_order"`
	IDUser            string    `json:"id_user" form:"id_user" url:"id_user"`
	IDStock           string    `json:"id_stock" form:"id_stock" url:"id_stock"`
	IDPelanggan       string    `json:"id_pelanggan" form:"id_pelanggan" url:"id_pelanggan"`
	JumlahBarang      int       `json:"jumlah_barang" form:"jumlah_barang" url:"jumlah_barang"`
	TotalHarga        string    `json:"total_harga" form:"total_harga" url:"total_harga"`
	PilihanPengiriman string    `json:"pilihan_pengiriman" form:"pilihan_pengiriman" url:"pilihan_pengiriman"`
	CreatedDate       time.Time `json:"created_datetime" form:"created_datetime" url:"created_datetime"`
	UpdatedDate       time.Time `json:"updated_datetime" form:"updated_datetime" url:"updated_datetime"`
	OrderStatus       string    `json:"order_status" form:"order_status" url:"order_status"`
	IsDeleted         bool      `json:"is_deleted" form:"is_deleted" url:"is_deleted"`
}

type ParamOrderEdit struct {
	IDOrder           string `json:"id_order" form:"id_order" url:"id_order"`
	JumlahBarang      int    `json:"jumlah_barang" form:"jumlah_barang" url:"jumlah_barang"`
	TotalHarga        string `json:"total_harga" form:"total_harga" url:"total_harga"`
	PilihanPengiriman string `json:"pilihan_pengiriman" form:"pilihan_pengiriman" url:"pilihan_pengiriman"`
}

type ParamOrderID struct {
	IDOrder string `json:"id_order" form:"id_order" url:"id_order"`
}

type ParamVirtualAccount struct {
	IDVA         string    `json:"id_va" form:"id_va" url:"id_va"`
	IDPelanggan  string    `json:"id_pelanggan" form:"id_pelanggan" url:"id_pelanggan"`
	Amount       string    `json:"amount" form:"amount" url:"amount"`
	NamaRekening string    `json:"nama_rekening" form:"nama_rekening" url:"nama_rekening"`
	NoVa         string    `json:"no_va" form:"no_va" url:"no_va"`
	CreatedDate  time.Time `json:"created_datetime" form:"created_datetime" url:"created_datetime"`
	VaStatus     string    `json:"va_status" form:"va_status" url:"va_status"`
	IsDeleted    bool      `json:"is_deleted" form:"is_deleted" url:"is_deleted"`
}

type ParamID struct {
	ID string `json:"id" form:"id" url:"id"`
}

type ParamNameCard struct {
	IDKartuNama string `json:"id_kartunama" form:"id_kartunama" url:"id_kartunama"`
	IDUser      string `json:"id_user" form:"id_user" url:"id_user"`
	NamaToko    string `json:"nama_toko" form:"nama_toko" url:"nama_toko"`
	BidangUsaha string `json:"bidang_usaha" form:"bidang_usaha" url:"bidang_usaha"`
	Alamat      string `json:"alamat" form:"alamat" url:"alamat"`
	NoTelepon   string `json:"no_telepon" form:"no_telepon" url:"no_telepon"`
	IsDeleted   bool   `json:"is_deleted" form:"is_deleted" url:"is_deleted"`
}

type ParamNameCardEdit struct {
	IDKartuNama string `json:"id_kartunama" form:"id_kartunama" url:"id_kartunama"`
	IDUser      string `json:"id_user" form:"id_user" url:"id_user"`
	NamaToko    string `json:"nama_toko" form:"nama_toko" url:"nama_toko"`
	BidangUsaha string `json:"bidang_usaha" form:"bidang_usaha" url:"bidang_usaha"`
	Alamat      string `json:"alamat" form:"alamat" url:"alamat"`
	NoTelepon   string `json:"no_telepon" form:"no_telepon" url:"no_telepon"`
}
