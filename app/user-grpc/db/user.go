package db

import (
	"github.com/jeepli/ichat/database"
	"github.com/pkg/errors"
)

var (
	insertUserSql        = `insert into users(email,name,password) values(?,?,?) returning *`
	selectUserByIdSql    = `select * from users where id in (?)`
	selectUserByEmailSql = `select * from users where email = ?`
)

type User struct {
	Id       string `pg:"id"`
	Email    string `pg:"email"`
	Password string `pg;"password"`
	Name     string `pg:"name"`
}

type UserDB struct {
	db *database.DBHolder
}

func NewUserDb(dh *database.DBHolder) *UserDB {
	return &UserDB{
		db: dh,
	}
}

// insert user
func (self *UserDB) InsertUser(u User) (*User, error) {
	var insertedUsers []User
	_, err := self.db.Query(&insertedUsers,
		insertUserSql,
		u.Email,
		u.Password,
		u.Name,
	)

	if err != nil {
		return nil, err
	}

	if len(insertedUsers) != 1 {
		return nil, errors.Errorf("insert user err.")
	}

	return &insertedUsers[0], nil
}

// select by ids
func (self *UserDB) SelectUsersByIds(ids []string) ([]User, error) {
	var selectedUsers []User
	_, err := self.db.Query(&selectedUsers,
		selectUserByIdSql,
		database.ArrayToSequence(ids),
	)

	if err != nil {
		return nil, err
	}

	return selectedUsers, nil
}

// select by email
func (self *UserDB) SelectUserByEmail(email string) ([]User, error) {
	var selectedUsers []User
	_, err := self.db.Query(&selectedUsers,
		selectUserByEmailSql,
		email,
	)

	if err != nil {
		return nil, err
	}

	return selectedUsers, nil
}
