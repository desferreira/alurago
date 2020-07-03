package routes

import (
	"github.com/desferreira/alurago/controllers"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new-product", controllers.Create)
	http.HandleFunc("/create", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.RenderEdit)
	http.HandleFunc("/update", controllers.Edit)

}
