package dhariancontroller

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

func KDharian(w http.ResponseWriter, _ *http.Request) {
	var kdh []models.Kdh
	Lg := []models.Head{}
	Ic := []models.Icon{}
	var response models.Ddh2
	if err := models.DB.Table("pray-menu").Select("`pray-menu`.`id`, `pray-menu`.`menu`").Group("`pray-menu`.`id`").Find(&kdh).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	logo := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if logo != nil {
		log.Print(logo.Error())
	}
	response.Icon = Ic
	response.Data = kdh
	response.Logo = Lg
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}

//show doa harianbyid controller
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menu, err := strconv.ParseInt(vars["menu"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var dh []models.Dh
	dh2 := []models.Dh2{}
	Lg := []models.Head{}
	Ic := []models.Icon{}
	var Ddh models.Ddh
	title := r.FormValue("title")
	// dhInput := models.Dh{menu: int(menu)}
	menu2 := models.DB.Table("pray-data").Select("`pray-data`.`id`, `pray-menu`.`menu`").Joins("INNER JOIN `pray-menu` ON `pray-data`.`menu` = `pray-menu`.`id`").Where("`pray-data`.`menu` = ?", menu).Limit(1).Find(&dh2).Error
	if menu2 != nil {
		log.Print(menu2.Error())
	}
	if err := models.DB.Table("pray-data").Select("`pray-data`.`id`, `pray-data`.`title`, `pray-data`.`arab`, `pray-data`.`latin`, `pray-data`.`meaning`, `pray-menu`.`menu`").Joins("INNER JOIN `pray-menu` ON `pray-data`.`menu` = `pray-menu`.`id`").Where("`pray-data`.`menu` = ? AND `pray-data`.`title` LIKE ?", menu, fmt.Sprintf("%%%s%%", title)).Find(&dh).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	result := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if result != nil {
		log.Print(result.Error())
	}
	Ddh.Menu = dh2
	Ddh.Data = dh
	Ddh.Logo = Lg
	Ddh.Icon = Ic
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(Ddh)
}
