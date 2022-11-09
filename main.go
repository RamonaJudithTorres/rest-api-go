package main

import (
	"net/http"

	"github.com/RamonaJudithTorres/rest-api-go/models"
	"github.com/RamonaJudithTorres/rest-api-go/routes"
	"github.com/culturadevops/GORM/libs"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	//ID       uint    `gorm:"primary_key;auto_increment"`
	Account  string `gorm:"type:varchar(20);not null;index:username"`
	Password string `gorm:"type:char(32);not null;"`
}

func main() {

	dbConfig := libs.Configure("./", "mysql")
	libs.DB = dbConfig.InitMysqlDB()
	libs.DB.AutoMigrate(&User{})
	libs.DB.AutoMigrate(models.Song{})
	libs.DB.AutoMigrate(models.Response{})

	router := mux.NewRouter()
	// auth routes:
	router.HandleFunc("/login", routes.Login)
	router.HandleFunc("/refresh", routes.Refresh)

	router.HandleFunc("/song", routes.PostSong).Methods("POST")
	http.ListenAndServe(":3000", router)
}
