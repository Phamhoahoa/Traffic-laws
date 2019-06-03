package dao

import(
	"LawDemo/model"
	"fmt"
	"context"
)

func GetUser(id int) model.User {
	Db := OpenDbConnection()
	query := fmt.Sprintf("SELECT * from user where id = %d", id)
	fmt.Println(query)
	var user model.User
	ctx := context.Background()
	_ = Db.QueryRowContext(ctx, query).Scan(&user.Id,&user.Username, &user.Password)
	defer Db.Close()
	return user
}