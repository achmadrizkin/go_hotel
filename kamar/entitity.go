package kamar

import "time"

type Kamar struct {
	Id           int
	Nama         string
	Image_Url    string
	Rating       string
	Harga        string
	Lokasi       string
	Deskripsi    string
	Jmlh_Kasur   string
	Jmlh_Ruangan string
	Kode_Kamar   string
	Status_Kamar string
	CreatedAt    time.Time
	UpdateAt     time.Time
}
