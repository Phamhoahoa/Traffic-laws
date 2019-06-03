package controller

import(
	"net/http"
	"LawDemo/dao"
	"github.com/gorilla/mux"
	"encoding/json"
	"LawDemo/model"
	"strconv"
	"fmt"

)

type UpdateNodeBody struct{
	Name string
	Id_law string
}

type CreateNodeBody struct{
	Name string
	Id_law string
	Parent int
}

func CreateNode(response http.ResponseWriter, request *http.Request){
	var node model.Node
	var tree model.Tree
	var body CreateNodeBody
	_ = json.NewDecoder(request.Body).Decode(&body)
	node.Name = body.Name
	node.Id_law= body.Id_law
	tree.Parent = body.Parent
	defer request.Body.Close()
	err := dao.CreateNode(&node, &tree)
	
	if err != nil{
		return
	}
}
func UpdateNode(response http.ResponseWriter, request *http.Request){
    vars := mux.Vars(request)
    id,_ := strconv.Atoi(vars["id"])
	var body UpdateNodeBody
	_ = json.NewDecoder(request.Body).Decode(&body)
	var node model.Node
	node.Id = id
	node.Name = body.Name
	node.Id_law=body.Id_law
	
	_ = dao.UpdateNode(id, node)
 }

func DeleteNodes(w http.ResponseWriter, r *http.Request){
	var res model.Respone

	params:=mux.Vars(r)
	id,_:=strconv.Atoi(params["id"])

	err:=dao.DeleteNode(id)
	if err!=nil{
		fmt.Println("loi1 delete_node_dao")
		return
	}else{
		fmt.Println("Node deleted!")

		res.Success="Node deleted!"
		json.NewEncoder(w).Encode(res)
	}
}

func GetSubNodes(w http.ResponseWriter, r *http.Request){
	var subNode[] model.Node

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	subNode=dao.GetSubNode(id)
	i:=len(subNode)
	if i!=0 {
	 json.NewEncoder(w).Encode(subNode)
	}else{
		var IdLaws []int 
		IdLaws = dao.GetIdLaw(id)
		var laws[] model.Law
		laws=dao.GetSliceLaw(IdLaws)
		json.NewEncoder(w).Encode(laws)
	}
}