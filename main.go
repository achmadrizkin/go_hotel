package main

import (
	"log"

	"go_hotel/handler"
	"go_hotel/kamar"
	"go_hotel/transaksi"
	"go_hotel/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// connect to database
	// PLEASE CREATE go-ecommerce database first.
	dsn := "root:@tcp(127.0.0.1:3306)/go_hotel?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// auto migrate (auto add table)
	db.AutoMigrate(&user.User{}, &transaksi.Transaksi{}, &kamar.Kamar{})

	// API Versioning
	v1 := r.Group("/v1")

	// KAMAR
	kamarRepository := kamar.NewRepository(db)
	kamarService := kamar.NewService(kamarRepository)
	kamarHandler := handler.NewKamarHandler(kamarService)

	v1.GET("/kamar", kamarHandler.GetKamarAll)
	v1.GET("/kamar/:nama", kamarHandler.GetKamarByNama)

	r.Run(":3000")
}
