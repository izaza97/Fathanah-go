package authcontroller

import (
	"crypto/sha256"
	"encoding/hex"
	"fathanah/helper"
	"fathanah/models"
	"net/http"

)

//register controller
func Register(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	name := r.Form.Get("name")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	passwordconfirm := r.Form.Get("passwordconfirm")
	profil := r.Form.Get("profil")
	userInput.Name = name
	userInput.Username = username
	userInput.Email = email
	userInput.Password = password
	userInput.Pp = profil
	// amankan pass menggunakan sha256
	pass := sha256.New()
	pass.Write([]byte(userInput.Password))
	shapass := pass.Sum(nil)
	userInput.Password = hex.EncodeToString(shapass)

	// ambil data user berdasarkan username
	var user models.User
	if name == "" || username == "" || email == "" || password == "" || passwordconfirm == "" {
		response := map[string]string{"Message": "FAILED"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	} else if err := models.DB.Table("web-user-data").Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		default:
			// insert ke database
			if password == passwordconfirm {
				models.DB.Table("web-user-data").Create(&userInput)
				response := map[string]string{"Message": "SUCCESS"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			} else {
				response := map[string]string{"Message": "FAILED"}
				helper.ResponseJSON(w, http.StatusInternalServerError, response)
				return
			}
		}
	} else {
		response := map[string]string{"Message": "FAILED"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}
}
