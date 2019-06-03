package controller

import(
	"net/http"
	"github.com/gorilla/mux"
	"LawDemo/dao"
	"strconv"
	"encoding/json"
)

func GetUsers(response http.ResponseWriter, request *http.Request) {
	
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	user := dao.GetUser(id)
	json.NewEncoder(response).Encode(user)
}
