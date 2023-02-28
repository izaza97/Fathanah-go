package mozaikcontroller

import (
	"encoding/json"
	"fathanah/helper"
	"fathanah/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

)

func Mozaik(w http.ResponseWriter, r *http.Request) {
	var mzk []models.Mozaikdata //model data mozaik
	hd := []models.Head{}       //model data logo
	Ic := []models.Icon{}       //model data icon
	var Dmzk models.Dmozaik     //model data akhir

	//query GET data mozaik
	if err := models.DB.Table("mozaik-data").Select("`mozaik-data`.`id`, `mozaik-data`.`time`, `mozaik-data`.`img`, `mozaik-data`.`title`,  `mozaik-data`.`desc`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `mozaik-data`.`path` =`img-path`.`id`").Find(&mzk).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	//query GET data logo header
	header := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&hd).Error
	if header != nil {
		log.Print(header.Error())
	}
	//query GET data icon
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	Dmzk.Icon = Ic
	Dmzk.Logo = hd
	Dmzk.Data = mzk
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(Dmzk)
}

func Mozaikview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var mzk []models.Mozaikdata
	var mzkl []models.Mozaikl
	hd := []models.Head{}
	Ic := []models.Icon{}
	var Dmzk models.Vmozaik

	if err := models.DB.Table("mozaik-data").Select("`mozaik-data`.`id`, `mozaik-data`.`time`, `mozaik-data`.`img`, `mozaik-data`.`title`,  `mozaik-data`.`desc`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `mozaik-data`.`path` =`img-path`.`id`").Where("`mozaik-data`.`id` =  ?", id).Find(&mzk).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	result2 := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`title`, `article-data`.`time`").Limit(5).Order("time DESC").Find(&mzkl).Error
	if result2 != nil {
		log.Print(result2.Error())
	}

	header := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&hd).Error
	if header != nil {
		log.Print(header.Error())
	}

	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	Dmzk.Icon = Ic
	Dmzk.Logo = hd
	Dmzk.Data = mzk
	Dmzk.Sidedata = mzkl
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(Dmzk)
}
