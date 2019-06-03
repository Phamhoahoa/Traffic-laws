package dao
import(
	"LawDemo/model"
	"fmt"
	"strings"
	"strconv"
	"context"

)
func UpdateNode(id int, node model.Node) error{
	db := OpenDbConnection()
	statement := fmt.Sprintf("UPDATE nodes set Name = '%s', Id_law='%s' where Id = %d", node.Name, node.Id_law, node.Id)
	_, err := db.Exec(statement)
  	if err != nil {
  	 	fmt.Println(err)
   	}
   	db.Close()
	return err
}
func CreateNode(node *model.Node, tree *model.Tree) error{
	db := OpenDbConnection()
	statement := fmt.Sprintf("INSERT INTO nodes(Name) VALUES('%s')", node.Name)
	_ ,err:= db.Exec(statement)
	query := fmt.Sprintf("SELECT LAST_INSERT_ID();")
	ctx := context.Background()
 	_ = db.QueryRowContext(ctx,query).Scan(&node.Id)
 	
	statement1 := fmt.Sprintf("INSERT INTO tree(Parent, Child, Depth) VALUES (%d, %d, 0)", node.Id, node.Id)
	statement2 := fmt.Sprintf("INSERT INTO tree(Parent, Child, Depth) SELECT Parent, %d, Depth+1 FROM tree WHERE Child = %d",node.Id, tree.Parent)
	
  	_, err = db.Exec(statement1)
 	_, err = db.Exec(statement2)
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
	return err
}
func DeleteNode(id int) error {
	Db:=OpenDbConnection()

	statement1:=fmt.Sprintf("DELETE p FROM tree p JOIN tree a USING (Child) WHERE a.Parent = %d;",id)
	_,err:=Db.Exec(statement1)

	Db.Close()
	return err
}
func GetSubNode(id int) []model.Node{
	Db:=OpenDbConnection()

	var subNode[] model.Node
	
	statement:=fmt.Sprintf("SELECT c.* FROM nodes c JOIN tree t  ON (c.Id = t.Child) WHERE t.Parent = %d  AND t.Depth = 1;", id)
	rows,_:=Db.Query(statement)
	for rows.Next(){
		var node model.Node
		_=rows.Scan(&node.Id, &node.Name, &node.Id_law)
		subNode =append(subNode,node)
	}
	Db.Close()
	return subNode
}
func GetIdLaw(id int) []int{
	Db:=OpenDbConnection()

	var IdLawString string
	statement:=fmt.Sprintf("select Id_law from nodes where Id=%d",id)
	_ = Db.QueryRow(statement).Scan(&IdLawString)
	var IdLawSliceString []string
	
	IdLawSliceString =strings.Split(IdLawString,",")
	var IdLawSliceInt []int
	for i:=0; i<len(IdLawSliceString); i++{
	
		value,_:=strconv.Atoi(IdLawSliceString[i])
		IdLawSliceInt=append(IdLawSliceInt,value)
		
	}
	Db.Close()
	return IdLawSliceInt
}