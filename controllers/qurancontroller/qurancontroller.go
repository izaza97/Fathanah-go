package qurancontroller

import (
	"encoding/json"
	"fathanah/helper"
	"fathanah/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

)

func Qrnsurah(w http.ResponseWriter, r *http.Request) {
	Qrn := []models.Qrns{}
	Lg := []models.Head{}
	Ic := []models.Icon{}
	var response models.Dqs
	surah := r.FormValue("surah")
	result := models.DB.Table("quran-surah").Where("`quran-surah`.`name` LIKE ?", fmt.Sprintf("%%%s%%", surah)).Find(&Qrn).Error
	if result != nil {
		log.Print(result.Error())
	}
	result1 := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if result1 != nil {
		log.Print(result1.Error())
	}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	response.Icon = Ic
	response.Logo = Lg
	response.Data = Qrn
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}

func Shows(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	surah, err := strconv.ParseInt(vars["surah"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var dh []models.Qrn
	Lg := []models.Head{}
	Ic := []models.Icon{}
	Srh := []models.Surah{}
	Pck := []models.Picked{}
	var response models.Dqrn
	// dhInput := models.Dh{menu: int(menu)}
	Pickedsurah := models.DB.Table("quran-surah").Where("id = ?", surah).Scan(&Pck).Error
	if Pickedsurah != nil {
		log.Print(Pickedsurah.Error())
	}
	if err := models.DB.Table("quran-data").Select("`quran-data`.`id`, `quran-surah`.`name`, `quran-data`.`arab`, `quran-data`.`latin`, `quran-data`.`meaning`").Joins("INNER JOIN `quran-surah` ON `quran-data`.`surah` = `quran-surah`.`id`").Where("`quran-data`.`surah` = ?", surah).Find(&dh).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	surat := models.DB.Table("quran-surah").Scan(&Srh).Error
	if surat != nil {
		log.Print(surat.Error())
	}
	logo := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if logo != nil {
		log.Print(logo.Error())
	}

	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}

	// back := models.DB.Table("quran-surah").Where("`id` >= ?-1 LIMIT 1", bck)

	// next := models.DB.Table("quran-surah").Where("`id` >= ?+1 LIMIT 1", nxt)

	response.Icon = Ic
	response.Logo = Lg
	response.Pickedsurah = Pck
	response.Surah = Srh
	response.Data = dh
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}
