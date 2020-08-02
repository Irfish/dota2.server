package orm

import (
	"errors"
	"fmt"

	orm "github.com/Irfish/component/xorm"
	"github.com/go-xorm/xorm"
)

type User struct {
	CreateTime int64  `xorm:"not null BIGINT(20)"`
	HeadUrl    string `xorm:"not null VARCHAR(1000)"`
	Id         int64  `xorm:"pk autoincr BIGINT(20)"`
	Level      int    `xorm:"not null INT(10)"`
	Phone      int    `xorm:"not null INT(11)"`
	Pwd        string `xorm:"not null VARCHAR(1024)"`
	SteamId    string `xorm:"not null VARCHAR(100)"`
	UserName   string `xorm:"not null VARCHAR(50)"`
}

func (p *User) Get(column string) interface{} {
	switch column {
	case "create_time":
		return p.CreateTime
	case "head_url":
		return p.HeadUrl
	case "id":
		return p.Id
	case "level":
		return p.Level
	case "phone":
		return p.Phone
	case "pwd":
		return p.Pwd
	case "steam_id":
		return p.SteamId
	case "user_name":
		return p.UserName
	}
	return nil
}

func (p *User) Gets(columns ...string) []interface{} {
	ret := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		ret = append(ret, p.Get(column))
	}
	return ret
}

type Users []*User

func NewUsers(cap int32) Users {
	return make(Users, 0, cap)
}

func (p Users) ToSlice(columns ...string) [][]interface{} {
	ret := make([][]interface{}, 0, len(p))
	for _, v := range p {
		ret = append(ret, v.Gets(columns...))
	}
	return ret
}

func GetUser(cols []string, query string, args ...interface{}) (*User, error) {
	obj := &User{}
	ok, err := UserXorm().
		Cols(cols...).
		Where(query, args...).
		Get(obj)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if !ok {
		return nil, fmt.Errorf("cont find User by %s (%v)", query, args)
	}
	return obj, nil
}

func UserXorm() *xorm.Session {
	return orm.Xorm(0).Table("user")
}
