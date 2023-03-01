package articlecontroller

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

func Kberita(w http.ResponseWriter, _ *http.Request) {
	kb := []models.Cart{}
	Lg := []models.Head{}
	var response models.Kdb
	result := models.DB.Table("article-category").Scan(&kb).Error
	if result != nil {
		log.Print(result.Error())
	}
	result1 := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if result1 != nil {
		log.Print(result1.Error())
	}
	Ic := []models.Icon{}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	response.Data = kb
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}

//show berita controller
func Showm(w http.ResponseWriter, r *http.Request) {
	var brt []models.Art
	Ac := []models.Cart{}
	var article models.Dart
	title := r.FormValue("title")
	category := models.DB.Table("article-category").Scan(&Ac).Error
	if category != nil {
		log.Print(category.Error())
	}
	if err := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`time`, `article-data`.`img`, `article-data`.`title`, `article-category`.`category`, `article-data`.`desc`, `img-path`.`path`").Joins("INNER JOIN `article-category` JOIN `img-path` ON `article-data`.`category` = `article-category`.`id` AND `article-data`.`path` =`img-path`.`id`").Where("`article-data`.`title` LIKE ?", fmt.Sprintf("%%%s%%", title)).Find(&brt).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	article.Category = Ac
	article.Data = brt
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(article)
}

func Showc(w http.ResponseWriter, r *http.Request) {
	var brt []models.Art
	Pac := []models.Cart{}
	Ac := []models.Cart{}
	var article models.Cartl
	vars := mux.Vars(r)
	title := r.FormValue("title")
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	categorypick := models.DB.Table("article-category").Where("id = ?", id).Scan(&Pac).Error
	if categorypick != nil {
		log.Print(categorypick.Error())
	}
	category := models.DB.Table("article-category").Scan(&Ac).Error
	if category != nil {
		log.Print(category.Error())
	}
	if err := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`time`, `article-data`.`img`, `article-data`.`title`, `article-category`.`category`, `article-data`.`desc`, `img-path`.`path`").Joins("INNER JOIN `article-category` JOIN `img-path` ON `article-data`.`category` = `article-category`.`id` AND `article-data`.`path` =`img-path`.`id`").Where("`article-data`.`title` LIKE ?", fmt.Sprintf("%%%s%%", title)).Where("`article-category`.`id` = ?", id).Find(&brt).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	article.PickC = Pac
	article.Category = Ac
	article.Data = brt
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(article)
}

//show berita controller
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var art []models.Art
	var artn []models.Artn
	var article models.Vart
	if err := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`time`, `article-data`.`img`, `article-data`.`title`, `article-category`.`category`, `article-data`.`desc`, `img-path`.`path`").Joins("INNER JOIN `article-category` JOIN `img-path` ON `article-data`.`category` = `article-category`.`id` AND `article-data`.`path` =`img-path`.`id`").Where("`article-data`.`id`= ?", id).Find(&art).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	result2 := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`title`, `article-data`.`time`").Limit(5).Order("time DESC").Find(&artn).Error
	if result2 != nil {
		log.Print(result2.Error())
	}
	article.Data = art
	article.Data2 = artn
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(article)
}
