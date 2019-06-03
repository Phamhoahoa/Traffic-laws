package controller

import(
	"net/http"
	"LawDemo/dao"
	"github.com/gorilla/mux"
	"encoding/json"
	"LawDemo/model"
	"strconv"
	"fmt"
	"strings"

)
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
func DeleteLawsCORS(w http.ResponseWriter, req *http.Request) {
	
	if (*req).Method == "OPTIONS" {
		setupResponse(&w, req)
		w.WriteHeader(http.StatusOK)
		return
	}
	
	DeleteLaws(w,req)
}
func GetLawFromUsers(w http.ResponseWriter, r *http.Request) {
	search:=r.URL.Query().Get("question")
	fmt.Println(r.URL.Query().Get("question"))
	if search!=""{
		mynewrequest := strings.Replace(search, " ", "+", -1)
	mynewrequest = "+" + mynewrequest
	var laws[] model.Law
	laws = dao.GetLawFromUser(mynewrequest)
	json.NewEncoder(w).Encode(laws)
	} else{
		var res model.Respone
		res.Err="Xin nhập chuỗi tìm kiếm vào!"
		json.NewEncoder(w).Encode(res)
	}
	
}
func InsertLaws(w http.ResponseWriter, r *http.Request) {
	var law model.Law
	var res model.Respone

	decoder:=json.NewDecoder(r.Body)
	fmt.Println(r.Body)
	err:=decoder.Decode(&law)
	if err!=nil{
		fmt.Println("loi1 insert_decode")
		return
	}
	defer r.Body.Close()
	fmt.Println(law.Content)
	fmt.Println(law.Id)

	err =dao.InsertLaw(&law)
	if err!=nil{
		fmt.Println("loi2 insert_dao")
		return
	}else{
		fmt.Println("insert_success")

		res.Success="Insert Success!"
		json.NewEncoder(w).Encode(res)

	}

}

func GetLaws(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	law := dao.GetLaw(id)
	fmt.Println(law)
	json.NewEncoder(w).Encode(law)
}

func GetAllLaws(w http.ResponseWriter, r *http.Request){
	var laws[] model.Law

	laws=dao.GetAllLaws()
	json.NewEncoder(w).Encode(laws)

}

func UpdateLaws(w http.ResponseWriter, r *http.Request){
	var law model.Law
	var res model.Respone

	decoder:=json.NewDecoder(r.Body)
	
	err:=decoder.Decode(&law)
	if err!=nil{
		return
	}
	defer r.Body.Close()
	
	fmt.Println(law.Content)
	fmt.Println(law.Id)
	fmt.Println(law.Object)


	err =dao.UpdateLaw(&law)
	if err!=nil{
		fmt.Println("update_err")
		return
	}else{
		fmt.Println("update_success")
		
		res.Success="Update Success!"
		json.NewEncoder(w).Encode(res)

	}
}

func DeleteLaws(w http.ResponseWriter, r *http.Request){
	var res model.Respone

	params:=mux.Vars(r)
	id,_:=strconv.Atoi(params["id"])

	fmt.Println(id)

	err:=dao.DeleteLaw(id)
	if err!=nil{
		fmt.Println("loi1 delete_dao")
		return
	}else{
		fmt.Println("Deleted!")

		res.Success="Deleted!"
		json.NewEncoder(w).Encode(res)

	}
}



// func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
//     response, _ := json.Marshal(payload)
//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(code)
//     w.Write(response)
// }