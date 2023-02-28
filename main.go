package main

import (
	"fathanah/controllers/ahd"
	"fathanah/controllers/articlecontroller"
	"fathanah/controllers/asetcontroller"
	"fathanah/controllers/authcontroller"
	"fathanah/controllers/dhariancontroller"
	"fathanah/controllers/diarycontroller"
	"fathanah/controllers/homecontroller"
	"fathanah/controllers/mozaikcontroller"
	"fathanah/controllers/mssgcontroller"
	"fathanah/controllers/qurancontroller"
	"fathanah/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")

	r.HandleFunc("/Img", authcontroller.Imgp).Methods("GET")
	r.HandleFunc("/home", homecontroller.Home).Methods("GET")
	r.HandleFunc("/logo-icon", homecontroller.Header).Methods("GET")

	r.HandleFunc("/ashusna", ahd.Ashusna).Methods("GET")

	r.HandleFunc("/dharian", dhariancontroller.KDharian).Methods("GET")
	r.HandleFunc("/dharians/{menu}", dhariancontroller.Show).Methods("GET")
	r.HandleFunc("/dharians/{menu}", dhariancontroller.Show).Methods("POST")

	r.HandleFunc("/articlee", articlecontroller.Kberita).Methods("GET")
	r.HandleFunc("/article", articlecontroller.Showm).Methods("GET")
	r.HandleFunc("/article", articlecontroller.Showm).Methods("POST")
	r.HandleFunc("/article/category/{id}", articlecontroller.Showc).Methods("GET")
	r.HandleFunc("/article/category/{id}", articlecontroller.Showc).Methods("POST")
	r.HandleFunc("/article/{id}", articlecontroller.Show).Methods("GET")

	r.HandleFunc("/quran", qurancontroller.Qrnsurah).Methods("GET")
	r.HandleFunc("/quran", qurancontroller.Qrnsurah).Methods("POST")
	r.HandleFunc("/quran/{surah}", qurancontroller.Shows).Methods("GET")

	r.HandleFunc("/mozaik", mozaikcontroller.Mozaik).Methods("GET")
	r.HandleFunc("/mozaik/{id}", mozaikcontroller.Mozaikview).Methods("GET")

	r.HandleFunc("/sign/{id}", authcontroller.Show).Methods("GET")
	r.HandleFunc("/sign/{id}/update", authcontroller.Updateprofile).Methods("POST")
	r.HandleFunc("/sign/{id}/updateuser", authcontroller.Updateusername).Methods("POST")
	r.HandleFunc("/sign/{id}/sun", authcontroller.Showun).Methods("GET")
	r.HandleFunc("/sign/{id}/updateun", authcontroller.Updateusername).Methods("POST")
	r.HandleFunc("/sign/{id}/updatepw", authcontroller.Updatepw).Methods("POST")

	r.HandleFunc("/sign/{user}/diary", diarycontroller.GetAllDiary).Methods("GET")
	r.HandleFunc("/sign/{user}/diary/create", diarycontroller.CreateDiary).Methods("POST")
	r.HandleFunc("/sign/{user}/diary/{no}", diarycontroller.GetDiary).Methods("GET")
	r.HandleFunc("/sign/{user}/diary/{no}/update", diarycontroller.UpdateDiary).Methods("POST")
	r.HandleFunc("/sign/{user}/diary/{no}/delete", diarycontroller.DeleteDiary).Methods("POST")

	r.HandleFunc("/sign/{id}/message", mssgcontroller.Message).Methods("POST")

	r.HandleFunc("/about", asetcontroller.About).Methods("GET")

	fmt.Println("Connected to port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
