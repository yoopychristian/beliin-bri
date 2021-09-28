package shared

type ResponseDetail struct {
	IDUser   string `json:"id_user"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Email    string `json:"email"`
	NoPonsel string `json:"no_ponsel"`
}

type StockReponse struct {
	IDStock      string `json:"id_stock"`
	IDUser       string `json:"id_user"`
	NamaBarang   string `json:"nama_barang"`
	HargaSatuan  string `json:"harga_barangsatuan"`
	Deskripsi    string `json:"deskripsi"`
	JumlahBarang int    `json:"jumlah_barang"`
	GambarBarang string `json:"gambar_barang"`
}

type StockListResponse struct {
	NamaBarang   string `json:"nama_barang"`
	Deskripsi    string `json:"deskripsi"`
	JumlahBarang int    `json:"jumlah_barang"`
	GambarBarang string `json:"gambar_barang"`
}

type CustomerResponse struct {
	IDPelanggan      string `json:"id_pelanggan"`
	IDUser           string `json:"id_user"`
	IDVa             string `json:"id_va"`
	Nama             string `json:"nama"`
	Email            string `json:"email"`
	AlamatPengiriman string `json:"alamat_pengiriman"`
	Kota             string `json:"kota"`
	NoTelepon        string `json:"no_telepon"`
}

type CustomerResponseEdit struct {
	IDPelanggan      string `json:"id_pelanggan"`
	Nama             string `json:"nama"`
	Email            string `json:"email"`
	AlamatPengiriman string `json:"alamat_pengiriman"`
	Kota             string `json:"kota"`
	NoTelepon        string `json:"no_telepon"`
}

type CustomerListResponse struct {
	IDPelanggan      string `json:"id_pelanggan"`
	Nama             string `json:"nama"`
	Email            string `json:"email"`
	AlamatPengiriman string `json:"alamat_pengiriman"`
	Kota             string `json:"kota"`
	NoTelepon        string `json:"no_telepon"`
}

type OrderResponse struct {
	IDOrder           string `json:"id_order"`
	IDUser            string `json:"id_user"`
	IDStock           string `json:"id_stock"`
	IDPelanggan       string `json:"id_pelanggan"`
	JumlahBarang      int    `json:"jumlah_barang"`
	TotalHarga        string `json:"total_harga"`
	PilihanPengiriman string `json:"pilihan_pengiriman"`
	OrderStatus       string `json:"order_status"`
}

type OrderResponseList struct {
	IDOrder           string `json:"id_order"`
	JumlahBarang      int    `json:"jumlah_barang"`
	TotalHarga        string `json:"total_harga"`
	PilihanPengiriman string `json:"pilihan_pengiriman"`
	OrderStatus       string `json:"order_status"`
}

type VAResponse struct {
	IDUser       string `json:"id_user"`
	NoVA         string `json:"no_va"`
	NamaRekening string `json:"nama_rekening"`
	Amount       string `json:"amount"`
	Status       string `json:"status"`
	ExpiredDate  string `json:"expiredDate"`
}

type BillDetail struct {
	IDPelanggan       string `json:"id_pelanggan"`
	Nama              string `json:"nama"`
	Alamat            string `json:"alamat_pengiriman"`
	NoTelepon         string `json:"no_telepon"`
	NamaBarang        string `json:"nama_barang"`
	JumlahBarang      int    `json:"jumlah_barang"`
	HargaBarang       int    `json:"harga_barangsatuan"`
	TotalHarga        int    `json:"total_harga"`
	Email             string `json:"email"`
	PilihanPengiriman string `json:"pilihan_pengiriman"`
	NoVa              string `json:"no_va"`
}

type NameCard struct {
	IDKartuNama string `json:"id_kartunama"`
	IDUser      string `json:"id_user"`
	NamaToko    string `json:"nama_toko"`
	BidangUsaha string `json:"bidang_usaha"`
	Alamat      string `json:"alamat"`
	NoTelepon   string `json:"no_telepon"`
}
