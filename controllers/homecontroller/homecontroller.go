package homecontroller

import (
	"encoding/json"
	"fathanah/helper"
	"fathanah/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

)

func Home(w http.ResponseWriter, r *http.Request) {
	sld := []models.Slide{}
	Ftr := []models.Feature{}
	brt := []models.Vbrt{}
	var response models.Home
	result1 := models.DB.Table("web-slider").Select("`web-slider`.`id`, `web-slider`.`name`, `web-slider`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `web-slider`.`path` =`img-path`.`id`").Find(&sld).Error
	if result1 != nil {
		log.Print(result1.Error())
	}
	result2 := models.DB.Table("web-feature").Select("`web-feature`.`id`, `web-feature`.`name`, `web-feature`.`img`, `img-path`.`path`, `web-feature`.`url`").Joins("INNER JOIN `img-path` ON `web-feature`.`path` =`img-path`.`id`").Find(&Ftr).Error
	if result2 != nil {
		log.Print(result2.Error())
	}
	result3 := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`time`, `article-data`.`sinopsis`, `article-data`.`img`, `article-data`.`title`, `article-category`.`category`, `article-data`.`desc`, `img-path`.`path`").Joins("INNER JOIN `img-path` JOIN `article-category` ON `article-data`.`path` =`img-path`.`id` AND `article-data`.`category` =`article-category`.`id`").Limit(5).Order("time DESC").Find(&brt).Error
	if result3 != nil {
		log.Print(result3.Error())
	}
	response.Data1 = sld
	response.Data2 = Ftr
	response.Data3 = brt
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}

//logo dan icon
func Header(w http.ResponseWriter, _ *http.Request) {
	Lg := []models.Head{}
	Ic := []models.Icon{}
	var response models.Hd
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	logo := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if logo != nil {
		log.Print(logo.Error())
	}
	response.Icon = Ic
	response.Logo = Lg
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}

func UserHeader(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	Lg := []models.Head{}
	Ic := []models.Icon{}
	Nu := []models.Nuser{}
	var response models.Hd
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	logo := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if logo != nil {
		log.Print(logo.Error())
	}
	if err := models.DB.Table("web-user-data").Select("`web-user-data`.id, `web-user-data`.`name`, `web-user-pp`.`img`, `img-path`.`path`").Joins("LEFT JOIN `web-user-pp` ON `web-user-data`.`pp` = `web-user-pp`.`id` LEFT JOIN `img-path` ON `img-path`.`id` = `web-user-pp`.`path`").Where("`web-user-data`.`id` = ?", id).Scan(&Nu).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "user tidak ditemukan")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	response.Icon = Ic
	response.Logo = Lg
	response.Navuser = Nu
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}
