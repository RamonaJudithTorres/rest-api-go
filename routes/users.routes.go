package routes

import (
	"encoding/json"
	"net/http"

	"fmt"
	"io/ioutil"
	"os"

	"github.com/RamonaJudithTorres/rest-api-go/models"
	"github.com/culturadevops/GORM/libs"
	"github.com/dgrijalva/jwt-go"
)

func PostSong(w http.ResponseWriter, r *http.Request) {

	song := r.FormValue("song")
	album := r.FormValue("album")

	url := "https://itunes.apple.com/search?limit=5&term=" + song + album


	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err.Error())
	}

	var responseObject models.Response
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Song); i++ {
		name := responseObject.Song[i]
		createdSong := libs.DB.Create(&name)
		err := createdSong.Error
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		json.NewEncoder(w).Encode(&name)
	}

	// AUTENTICACION JWT
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

}
