package authcontroller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fathanah/helper"
	"fathanah/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

//show userbyid controller
func Show(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user []models.UserS
	var useri models.Ud
	if err := models.DB.Table("web-user-data").Select("`web-user-data`.id, `web-user-data`.`name`, `web-user-data`.`username`, `web-user-data`.`email`, `web-user-data`.`password`,`web-user-pp`.`img`").Joins("INNER JOIN `web-user-pp` ON `web-user-data`.`pp` = `web-user-pp`.`id`").Where("`web-user-data`.`id` = ?", id).Scan(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "user tidak ditemukan")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	useri.Data = user
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(useri)

}

func Showun(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user []models.UU
	var useri models.Un
	if err := models.DB.Table("web-user-data").First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "user tidak ditemukan")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	useri.Data = user
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(useri)
}

//update user controller
func Updateprofile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var userInput models.UserU
	errr := r.ParseForm()
	if errr != nil {
		panic(err)
	}
	name := r.Form.Get("name")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	pp := r.Form.Get("pp")
	userInput.Username = username
	userInput.Name = name
	userInput.Email = email
	userInput.Pp = pp

	defer r.Body.Close()

	//
	var user models.UU
	if err := models.DB.Table("web-user-data").Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		default:
			// insert ke database
			models.DB.Table("web-user-data").Where("id = ?", id).Updates(&userInput)
			response := map[string]string{"Message": "SUCCESS"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	} else {
		response := map[string]string{"Message": "FAILED"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}
}

func Updatepw(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var userInput models.UP
	var oldpass models.UP
	errr := r.ParseForm()
	if errr != nil {
		panic(err)
	}
	oldpassword := r.Form.Get("oldpassword")
	password := r.Form.Get("password")
	passwordconfirm := r.Form.Get("passwordconfirm")
	oldpass.Password = oldpassword
	userInput.Password = password

	defer r.Body.Close()

	//
	oldpas := sha256.New()
	oldpas.Write([]byte(oldpass.Password))
	shaoldpas := oldpas.Sum(nil)
	oldpass.Password = hex.EncodeToString(shaoldpas)
	pass := sha256.New()
	pass.Write([]byte(userInput.Password))
	shapass := pass.Sum(nil)
	userInput.Password = hex.EncodeToString(shapass)

	var user models.UP
	if err := models.DB.Table("web-user-data").Where("password = ? AND id = ?", oldpass.Password, id).First(&user).Error; err != nil {
		switch err {
		default:
			response := map[string]string{"Message": "FAILED"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}
	} else {
		// insert ke database
		if password == passwordconfirm {
			models.DB.Table("web-user-data").Where("id = ?", id).Updates(&userInput)
			response := map[string]string{"Message": "SUCCESS"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		} else {
			response := map[string]string{"Message": "password not same"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}
}

func Forgotpass(w http.ResponseWriter, r *http.Request) {

	// mengambil inputan json
	var userInput models.User
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	email := r.Form.Get("email")
	userInput.Email = email

	// ambil data user berdasarkan username
	var user models.User
	if err := models.DB.Table("web-user-data").Where("email= ?", userInput.Email).First(&user).Error; err != nil {
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

	var Ui models.Userlogin
	result := models.DB.Table("web-user-data").Where("email = ?", userInput.Email).First(&Ui).Error
	if result != nil {
		log.Print(result.Error())
		Ui.Message = "FAILED"
		w.Header().Set("Content-Type", "appication/json")
		helper.ResponseJSON(w, http.StatusOK, Ui)
	} else {
		Ui.Message = "SUCCESS"
		w.Header().Set("Content-Type", "appication/json")
		helper.ResponseJSON(w, http.StatusOK, Ui)
	}
}

func Newpass(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var userInput models.UP
	errr := r.ParseForm()
	if errr != nil {
		panic(err)
	}
	password := r.Form.Get("newpassword")
	passwordconfirm := r.Form.Get("passwordconfirm")
	userInput.Password = password
	pass := sha256.New()
	pass.Write([]byte(userInput.Password))
	shapass := pass.Sum(nil)
	userInput.Password = hex.EncodeToString(shapass)

	//
	if password == passwordconfirm {
		models.DB.Table("web-user-data").Where("id = ?", id).Updates(&userInput)
		response := map[string]string{"Message": "SUCCESS"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	} else {
		response := map[string]string{"Message": "password not same"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
}

func Imgp(w http.ResponseWriter, _ *http.Request) {
	img := []models.Imgp{}
	var response models.Ip
	result := models.DB.Table("web-user-pp").Select("`web-user-pp`.id, `web-user-pp`.`img`, `web-user-pp`.`name`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `web-user-pp`.`path` =`img-path`.`id`").Scan(&img).Error
	if result != nil {
		log.Print(result.Error())
	}
	response.Data = img
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}
