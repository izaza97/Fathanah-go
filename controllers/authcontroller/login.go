package authcontroller

import (
	"crypto/sha256"
	"encoding/hex"
	"fathanah/helper"
	"fathanah/models"
	"log"
	"net/http"

	"gorm.io/gorm"

)

//login controllers
func Login(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	usernameemail := r.Form.Get("username")
	password := r.Form.Get("password")
	userInput.Username = usernameemail
	userInput.Email = usernameemail
	userInput.Password = password

	// ambil data user berdasarkan username
	var user models.User
	if err := models.DB.Table("web-user-data").Where("username = ?", userInput.Username).Or("email= ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"Message": "FAILED"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		default:
			response := map[string]string{"Message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek apakah password valid
	pass := sha256.New()
	pass.Write([]byte(userInput.Password))
	shapass := pass.Sum(nil)
	userInput.Password = hex.EncodeToString(shapass)
	if err := models.DB.Table("web-user-data").Where("password = ?", userInput.Password).Scan(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"Message": "FAILED"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	var Ui models.Userlogin
	result := models.DB.Table("web-user-data").Where("username = ?", userInput.Username).Or("email= ?", userInput.Email).First(&Ui).Error
	if result != nil {
		log.Print(result.Error())
	}

	Ui.Message = "SUCCESS"
	w.Header().Set("Content-Type", "appication/json")
	helper.ResponseJSON(w, http.StatusOK, Ui)
}
