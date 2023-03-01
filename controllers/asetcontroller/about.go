package asetcontroller

import (
	"encoding/json"
	"fathanah/models"
	"log"
	"net/http"

)

func About(w http.ResponseWriter, _ *http.Request) {
	ab := []models.Ab{}
	ab2 := []models.Ab{}
	var response models.Dab
	Data := models.DB.Table("about-data").Select("`about-data`.`id`, `about-data`.`desc`, `about-data`.`name`, `about-data`.`img`, `img-asset`.`path`").Joins("INNER JOIN `img-asset` ON `about-data`.`img` =`img-asset`.`id`").Where("`about-data`.`id` = 1").Scan(&ab).Error
	if Data != nil {
		log.Print(Data.Error())
	}
	Data2 := models.DB.Table("about-data").Select("`about-data`.`id`,`about-data`.`desc`, `about-data`.`name`, `about-data`.`img`, `img-asset`.`path`").Joins("INNER JOIN `img-asset` ON `about-data`.`img` =`img-asset`.`id`").Where("`about-data`.`id` = 2").Scan(&ab2).Error
	if Data2 != nil {
		log.Print(Data.Error())
	}
	response.Data1 = ab
	response.Data2 = ab2
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}
