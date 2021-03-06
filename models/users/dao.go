package users

import (
	"github.com/crob1140/codewiz-server/datastore"
)

type Dao struct {
	DB *datastore.DB
}

func NewDao(db *datastore.DB) *Dao {
	db.AddTableWithName(User{}, "Users")
	return &Dao{DB : db}
}

func (dao *Dao) GetByID(id uint64) (*User, error) {
	user, err := dao.DB.Get(User{}, "SELECT * FROM Users WHERE ID = ?", id)
	if err != nil || user == nil {
		return nil, err
	}
	return user.(*User), err
}

func (dao *Dao) GetByUsername(username string) (*User, error) {
	user, err := dao.DB.Get(User{}, "SELECT * FROM Users WHERE Username = ?", username)
	if err != nil || user == nil {
		return nil, err
	}
	return user.(*User), err
}

func (dao *Dao) Update(user *User) error {
	return dao.DB.Update(user)
}

func (dao *Dao) Delete(user *User) error {
	return dao.DB.Delete(user)
}

func (dao *Dao) Insert(user *User) error {
	err := dao.DB.Insert(user)
	return err
}	