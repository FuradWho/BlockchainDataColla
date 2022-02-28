package main

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB

func init() {
	opt, err := pg.ParseURL("postgres://furad:furad@localhost:5432/db_micro?sslmode=disable")
	if err != nil {
		panic(err)
	}

	//db
	db = pg.Connect(opt)
}

type User struct {
	id   int
	name string
}

//通过定义的结构体来创建数据库表
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			//Temp: true,//建表是临时的，测试用途
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := createSchema(db)
	if err != nil {
		fmt.Println(err)
	}

	user1 := &User{
		id:   10,
		name: "zhansan",
	}

	insert, err := db.Model(user1).Insert()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(insert)
}
