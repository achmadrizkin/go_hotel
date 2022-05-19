package kamar

import "time"

type KamarResponse struct {
	Id           int       `json:"id"`
	Nama         string    `json:"nama"`
	Image_Url    string    `json:"image_url"`
	Rating       string    `json:"rating"`
	Harga        string    `json:"harga"`
	Lokasi       string    `json:"lokasi"`
	Deskripsi    string    `json:"Deskripsi"`
	Jmlh_Kasur   string    `json:"jmlh_kasur"`
	Jmlh_Ruangan string    `json:"jmlh_ruangan"`
	Kode_Kamar   string    `json:"kode_kamar"`
	Status_Kamar string    `json:"status_kamar"`
	CreatedAt    time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
}
