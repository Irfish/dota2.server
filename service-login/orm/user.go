package orm

import (
	"errors"
	"fmt"

	orm "github.com/Irfish/component/xorm"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id          int64  `xorm:"pk autoincr BIGINT(10)"`
	SteamId     string `xorm:"not null VARCHAR(10)"`
	SteamName   string `xorm:"VARCHAR(20)"`
	SteamGold   int64  `xorm:"not null BIGINT(10)"`
	SteamSilver int64  `xorm:"not null BIGINT(10)"`
	SteamVipExp int64  `xorm:"not null BIGINT(10)"`
	CreateTime  int64  `xorm:"not null BIGINT(20)"`
	UpdateTime  int64  `xorm:"not null BIGINT(20)"`
}

func (p *User) Get(column string) interface{} {
	switch column {
	case "id":
		return p.Id
	case "steam_id":
		return p.SteamId
	case "steam_name":
		return p.SteamName
	case "steam_gold":
		return p.SteamGold
	case "steam_silver":
		return p.SteamSilver
	case "steam_vip_exp":
		return p.SteamVipExp
	case "create_time":
		return p.CreateTime
	case "update_time":
		return p.UpdateTime
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
