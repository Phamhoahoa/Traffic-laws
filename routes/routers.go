
package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"LawDemo/controller"
	"fmt"
)
//trả về kiểu dữ liệu handler=>router
func UserRouter() http.Handler {
	router:= mux.NewRouter()
	router.HandleFunc("/user/{id}", controller.GetUsers).Methods("GET")
	return router
}
func LawRouter() http.Handler{
	
	router:=mux.NewRouter()
	fmt.Println("lalala")
	router.HandleFunc("/api/home/law/insert",controller.InsertLaws).Methods("POST")
	router.HandleFunc("/api/home/law/get/{id}",controller.GetLaws).Methods("GET")
	router.HandleFunc("/api/home/law/search",controller.GetLawFromUsers).Methods("GET")
	router.HandleFunc("/api/home/law/delete/{id}",controller.DeleteLaws).Methods("DELETE")
	router.HandleFunc("/api/home/law/getall",controller.GetAllLaws).Methods("GET")
	router.HandleFunc("/api/home/law/update",controller.UpdateLaws).Methods("POST")

	router.HandleFunc("/api/home/node/delete/{id}",controller.DeleteNodes).Methods("DELETE")	
	router.HandleFunc("/api/home/node/getsubnode/{id}",controller.GetSubNodes).Methods("GET")
	router.HandleFunc("/api/home/node/update/{id}", controller.UpdateNode).Methods("PUT")
	router.HandleFunc("/api/home/node/create", controller.CreateNode).Methods("POST")
	return router
}
