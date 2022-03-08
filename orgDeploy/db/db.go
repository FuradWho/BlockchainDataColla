package db

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

var Db *pg.DB

func init() {
	opt, err := pg.ParseURL("postgres://furad:furad@localhost:5432/db_micro?sslmode=disable")
	if err != nil {
		panic(err)
	}

	//db
	Db = pg.Connect(opt)
}

type CompanyInfo struct {
	tableName      struct{} `pg:"schema_colla.t_companyInfo,alias:t_companyInfo"`
	Id             string   `pg:"id,pk"`
	Ip             string   `pg:"ip"`
	InvitationCode string   `pg:"invitation_code"`
}

func main() {
	com := &CompanyInfo{
		Id: uuid.NewString(),
		Ip: "1212121",
	}

	insert, err := Db.Model(com).Insert()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(insert)
}
