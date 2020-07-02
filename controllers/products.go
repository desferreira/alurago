package controllers

import (
	"fmt"
	"github.com/desferreira/alurago/db"
	"github.com/desferreira/alurago/models"
	"html/template"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request){
	database := db.CreateConnection()
	defer db.CloseConnection(database)

	products := models.GetAll()

	temp.ExecuteTemplate(w, "index", products)
}

func Create(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "create", nil)
}

func Insert(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		formatedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Wrong conversion format")
		}

		formatedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			fmt.Println("Wrong conversion format")
		}

		models.Create(name, description, formatedPrice, formatedQuantity)
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		fmt.Println("Redirecting to home")
	}
}

func Delete(w http.ResponseWriter, r *http.Request){
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)
	models.Delete(id)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}