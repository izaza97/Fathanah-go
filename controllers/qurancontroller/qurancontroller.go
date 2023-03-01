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
	var response models.Dqs
	surah := r.FormValue("surah")
	result := models.DB.Table("quran-surah").Where("`quran-surah`.`name` LIKE ?", fmt.Sprintf("%%%s%%", surah)).Find(&Qrn).Error
	if result != nil {
		log.Print(result.Error())
	}
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
	Srh := []models.Surah{}
	Pck := []models.Picked{}
	bck := []models.Back{}
	nxt := []models.Next{}
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

	if surah == 1 {
		Back := models.DB.Table("quran-surah").Where("id = 114").Scan(&bck).Error
		if Back != nil {
			log.Print(Back.Error())
		}
	} else {
		Back := models.DB.Table("quran-surah").Where("id >= ?-1 LIMIT 1", surah).Scan(&bck).Error
		if Back != nil {
			log.Print(Back.Error())
		}
	}

	if surah == 114 {
		Next := models.DB.Table("quran-surah").Where("id = 1").Scan(&nxt).Error
		if Next != nil {
			log.Print(Next.Error())
		}
	} else {
		Next := models.DB.Table("quran-surah").Where("id >= ?+1 LIMIT 1", surah).Scan(&nxt).Error
		if Next != nil {
			log.Print(Next.Error())
		}
	}

	response.Pickedsurah = Pck
	response.Surah = Srh
	response.Data = dh
	response.Back = bck
	response.Next = nxt
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}
