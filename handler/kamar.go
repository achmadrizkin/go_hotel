package handler

import (
	"net/http"

	"go_hotel/kamar"

	"github.com/gin-gonic/gin"
)

type kamarHandler struct {
	kamarService kamar.Service
}

func NewKamarHandler(kamarService kamar.Service) *kamarHandler {
	return &kamarHandler{kamarService}
}

func (h *kamarHandler) GetKamarAll(c *gin.Context) {
	allproductss, err := h.kamarService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	var allproductssResponse []kamar.KamarResponse

	for _, b := range allproductss {
		allproductsResponse := converToAllProductResponse(b)
		allproductssResponse = append(allproductssResponse, allproductsResponse)
	}

	//BEWARE DONT TOUCH THIS CODE
	if allproductssResponse != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": allproductssResponse,
		})
	}
}

func converToAllProductResponse(b kamar.Kamar) kamar.KamarResponse {
	return kamar.KamarResponse { 
		Id:           b.Id,
		Nama:         b.Nama,
		Image_Url:    b.Image_Url,
		Rating:       b.Rating,
		Harga:        b.Harga,
		Lokasi:       b.Lokasi,
		Deskripsi:    b.Deskripsi,
		Jmlh_Kasur:   b.Jmlh_Kasur,
		Jmlh_Ruangan: b.Jmlh_Ruangan,
		Kode_Kamar:   b.Kode_Kamar,
		Status_Kamar: b.Status_Kamar,
	}
}
