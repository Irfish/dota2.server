package orm

import (
	"errors"
	"fmt"

	orm "github.com/Irfish/component/xorm"
	"github.com/go-xorm/xorm"
)

type CardKey struct {
	Id         int64  `xorm:"pk autoincr BIGINT(10)"`
	KeyCode    string `xorm:"not null VARCHAR(10)"`
	KeyState   int    `xorm:"not null INT(11)"`
	KeyType    int    `xorm:"not null INT(11)"`
	KeyValue   int64  `xorm:"not null BIGINT(10)"`
	KeyRmb     int64  `xorm:"not null BIGINT(10)"`
	CreateTime int64  `xorm:"not null BIGINT(20)"`
	UpdateTime int64  `xorm:"not null BIGINT(20)"`
}

func (p *CardKey) Get(column string) interface{} {
	switch column {
	case "id":
		return p.Id
	case "key_code":
		return p.KeyCode
	case "key_state":
		return p.KeyState
	case "key_type":
		return p.KeyType
	case "key_value":
		return p.KeyValue
	case "key_rmb":
		return p.KeyRmb
	case "create_time":
		return p.CreateTime
	case "update_time":
		return p.UpdateTime
	}
	return nil
}

func (p *CardKey) Gets(columns ...string) []interface{} {
	ret := make([]interface{}, 0, len(columns))
	for _, column := range columns {
		ret = append(ret, p.Get(column))
	}
	return ret
}

type CardKeys []*CardKey

func NewCardKeys(cap int32) CardKeys {
	return make(CardKeys, 0, cap)
}

func (p CardKeys) ToSlice(columns ...string) [][]interface{} {
	ret := make([][]interface{}, 0, len(p))
	for _, v := range p {
		ret = append(ret, v.Gets(columns...))
	}
	return ret
}

func GetCardKey(cols []string, query string, args ...interface{}) (*CardKey, error) {
	obj := &CardKey{}
	ok, err := CardKeyXorm().
		Cols(cols...).
		Where(query, args...).
		Get(obj)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if !ok {
		return nil, fmt.Errorf("cont find CardKey by %s (%v)", query, args)
	}
	return obj, nil
}

func CardKeyXorm() *xorm.Session {
	return orm.Xorm(0).Table("card_key")
}
