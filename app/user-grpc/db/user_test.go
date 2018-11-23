package db

import (
	"testing"

	"github.com/jeepli/ichat/config"
	"github.com/jeepli/ichat/database"
)

var gConf = &config.DbConfig{
	Address:  "127.0.0.1:5432",
	User:     "ichat",
	Database: "ichat",
	Password: "",
}

func TestCreateDB(t *testing.T) {
	db := database.NewDBHolder(gConf)
	if db == nil {
		t.Fatal("create db failed")
	}
}

func TestCRUD(t *testing.T) {
	db := database.NewDBHolder(gConf)
	if db == nil {
		t.Fatal("create db failed")
	}

	email := "lijipeng@p1.com"
	passwd := "test_passwd"
	name := "test_name"

	// insert
	var insertedUsers []User
	res, err := db.Query(&insertedUsers, "insert into users(email,password,name) values(?,?,?) returning *", email, passwd, name)

	if err != nil {
		t.Fatalf("insert err %v", err.Error())
		return
	}

	if res == nil {
		t.Fatal("insert result nil")
		return
	}

	if res.RowsAffected() != 1 {
		t.Fatal("rows affected 0")
		return
	}

	if len(insertedUsers) != 1 {
		t.Fatalf("insertedUsers len %v", len(insertedUsers))
		return
	}

	// select
	userId := insertedUsers[0].Id
	var selectUsers []User
	res, err = db.Query(&selectUsers, "select * from users where id = ?", userId)
	if err != nil {
		t.Fatalf("select err %v", err.Error())
		return
	}

	if res == nil {
		t.Fatal("select result nil")
		return
	}

	if len(selectUsers) != 1 {
		t.Fatalf("selectedUsers len %v", len(selectUsers))
		return
	}

	selectUser := selectUsers[0]
	if selectUser.Id != userId || selectUser.Email != email || selectUser.Name != name || selectUser.Password != passwd {
		t.Fatalf("selectUser info err %+v", selectUser)
		return
	}

	// delete
	res, err = db.Exec("delete from users where id = ?", userId)
	if err != nil {
		t.Fatalf("select err %v", err.Error())
		return
	}

	if res == nil {
		t.Fatal("select result nil")
		return
	}

	if res.RowsAffected() != 1 {
		t.Fatal("select result nil")
		return
	}
}
