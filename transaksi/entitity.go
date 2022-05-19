package transaksi

type Transaksi struct {
	Id               int
	Kode_Transaksi   string
	Email_Pemesan    string
	Kode_Kamar       string
	Status_Transaksi string
	Pembayaran_via   string
	Biaya            string
	Check_In         string
	Check_Out        string
}
