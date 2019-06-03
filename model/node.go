package model

 type Node struct{
	 Id int `json: "id" `
	 Name string `json: "name"`
	 Id_law string `json: "Id_law"`
 }
 type Tree struct{
	Parent int
	Child int
}