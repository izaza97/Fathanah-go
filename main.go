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

	r.HandleFunc("/asmaulhusna", ahd.Ashusna).Methods("GET")

	r.HandleFunc("/doaharian", dhariancontroller.KDharian).Methods("GET")
	r.HandleFunc("/doaharians/{menu}", dhariancontroller.Show).Methods("GET")
	r.HandleFunc("/doaharians/{menu}", dhariancontroller.Show).Methods("POST")

	r.HandleFunc("/articlee", articlecontroller.Kberita).Methods("GET")
	r.HandleFunc("/article", articlecontroller.Showm).Methods("GET")
	r.HandleFunc("/article", articlecontroller.Showm).Methods("POST")
	r.HandleFunc("/article/category/{id}", articlecontroller.Showc).Methods("GET")
	r.HandleFunc("/article/category/{id}", articlecontroller.Showc).Methods("POST")
	r.HandleFunc("/article/{id}", articlecontroller.Show).Methods("GET")

	r.HandleFunc("/alquran", qurancontroller.Qrnsurah).Methods("GET")
	r.HandleFunc("/alquran", qurancontroller.Qrnsurah).Methods("POST")
	r.HandleFunc("/alquran/{surah}", qurancontroller.Shows).Methods("GET")

	r.HandleFunc("/mozaikislam", mozaikcontroller.Mozaik).Methods("GET")
	r.HandleFunc("/mozaikislam/{id}", mozaikcontroller.Mozaikview).Methods("GET")

	r.HandleFunc("/sign/{id}/Message", mssgcontroller.Message).Methods("POST")

	r.HandleFunc("/about", asetcontroller.About).Methods("GET")

	r.HandleFunc("/sign/logo-icon/{id}", homecontroller.UserHeader).Methods("GET")

	r.HandleFunc("/sign/{id}", authcontroller.Show).Methods("GET")
	r.HandleFunc("/sign/{id}/updateuser", authcontroller.Updateprofile).Methods("POST")
	r.HandleFunc("/sign/{id}/userdata", authcontroller.Showun).Methods("GET")
	r.HandleFunc("/sign/{id}/updatepw", authcontroller.Updatepw).Methods("POST")
	r.HandleFunc("/forgotpass", authcontroller.Forgotpass).Methods("POST")
	r.HandleFunc("/sign/{id}/forgotpass", authcontroller.Newpass).Methods("POST")

	r.HandleFunc("/sign/{user}/diary", diarycontroller.GetAllDiary).Methods("GET")
	r.HandleFunc("/sign/{user}/diary/create", diarycontroller.CreateDiary).Methods("POST")
	r.HandleFunc("/sign/{user}/diary/{no}", diarycontroller.GetDiary).Methods("GET")
	r.HandleFunc("/sign/{user}/diary/{no}/update", diarycontroller.UpdateDiary).Methods("POST")
	r.HandleFunc("/sign/{user}/diary/{no}/delete", diarycontroller.DeleteDiary).Methods("POST")

	fmt.Println("Connected to port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
