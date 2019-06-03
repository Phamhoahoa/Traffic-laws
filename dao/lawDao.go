package dao
import (
	"LawDemo/model"
	"fmt"
	"context"
	
)
func GetLawFromUser(s string) []model.Law {
	Db := OpenDbConnection()

	var laws[] model.Law
    
	query := fmt.Sprintf("select * from law where match(content,object) against ('%s' in boolean mode)",s)
	fmt.Println(query)
	rows,_:= Db.Query(query)
	for rows.Next(){
		var law model.Law
		_=rows.Scan(&law.Id, &law.Content, &law.Object, &law.Money, &law.Extra, &law.Hold, &law.Source)
		laws=append(laws,law)
	}
	defer Db.Close()
	return laws
 }

func InsertLaw(law *model.Law) error {
	 Db := OpenDbConnection()
	 
 	statement := fmt.Sprintf("INSERT INTO law(content,object,money,extra,hold,source) VALUES('%s', '%s', '%s', '%s', '%s', '%s')", law.Content, law.Object, law.Money,law.Extra,law.Hold,law.Source)
	_, err := Db.Exec(statement)
	return err
}

func GetLaw(id int) model.Law{
	Db:=OpenDbConnection()

	var law model.Law
	ctx := context.Background()
	
	statement:=fmt.Sprintf("select * from law where id = %d", id)
	
	_ = Db.QueryRowContext(ctx, statement).Scan(&law.Id, &law.Content,&law.Object,&law.Money,&law.Extra,&law.Hold,&law.Source)
	defer Db.Close()
	return law
}

func GetSliceLaw(id []int) []model.Law{
	Db:=OpenDbConnection()

	var laws[] model.Law
	
	for i:=0;i<len(id);i++{
		var law model.Law
		statement:=fmt.Sprintf("select *from law where id=%d",id[i])
		_=Db.QueryRow(statement).Scan(&law.Id, &law.Content,&law.Object,&law.Money,&law.Extra,&law.Hold,&law.Source)
		laws=append(laws,law)
	}
	defer Db.Close()
	return laws
}

func GetAllLaws()[]model.Law{
	Db:=OpenDbConnection()

	var laws[] model.Law

	statement:=fmt.Sprintf("select * from law")
	rows,_:=Db.Query(statement)
	for rows.Next(){
		var law model.Law
		_=rows.Scan(&law.Id, &law.Content, &law.Object, &law.Money, &law.Extra, &law.Hold, &law.Source)
		laws=append(laws,law)
	}
	defer Db.Close()
	return laws
}

func UpdateLaw( law *model.Law) error{
	Db:=OpenDbConnection()

	statement:=fmt.Sprintf("update law set content='%s', object='%s', money='%s', extra='%s', hold='%s', source='%s' where id=%d",law.Content,law.Object,law.Money,law.Extra,law.Hold,law.Source,law.Id)
	_,err:=Db.Exec(statement)
	return err
}
func DeleteLaw(id int) error {
	Db:=OpenDbConnection()

	statement:=fmt.Sprintf("delete from law where id=%d",id)
	_,err:=Db.Exec(statement)
	return err
}

