package gin

import (
	"fmt"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/str"
	"github.com/Irfish/component/xorm"
	"github.com/Irfish/dota2.server/service-login/orm"
	"os"
	"time"
)

const (
	KEY_STATE_NORMAL =1
	KEY_STATE_USED =2
)

const (
	KEY_TYPE_1 =1
)

func GenCardKeys(money, value , count int) bool {
	if count==0 {
		return true
	}
	t := time.Now().Unix()
	keys := make(map[string]bool)
	cards:= make([]orm.CardKey,0,count)
	i:=0;
	for i<count {
		code := str.RandomString(8)
		if b,_:=keys[code];!b {
			keys[code] = true
			i++
			card:=orm.CardKey{
				KeyCode: code,
				KeyState: KEY_STATE_NORMAL,
				KeyRmb: int64(money),
				KeyValue: int64(value),
				KeyType: KEY_TYPE_1,
				CreateTime: t,
				UpdateTime: t,
			}

			cards = append(cards,card)
		}
	}

	var e error
	s := xorm.Xorm(0).NewSession()
	defer func() {
		if e!=nil {
			log.Debug("GenCardKeys err:",e.Error())
			s.Rollback()
		}
		s.Close()
	}()
	if err := s.Begin() ; err != nil {
		e = fmt.Errorf("fail to session begin:%s",err.Error())
		return false
	}

	s1:= s.Table("card_key")
	effect,err:= s1.Insert(cards)
	if err!=nil {
		e = fmt.Errorf("insert cards failed:%s",err.Error())
		return false
	}
	if effect == int64(count) {
		s.Commit()
		return true
	}else {
		e = fmt.Errorf("insert cards failed %d",effect)
		return false
	}
}

func InsertCardKey( card orm.CardKey)  {
	_,err:= orm.CardKeyXorm().Insert(card)
	if err!=nil {
		log.Debug(err.Error())
	}
}

func ExportTxtFile() bool  {
	var cards []orm.CardKey
	err:= orm.CardKeyXorm().Where("key_state=?",KEY_STATE_NORMAL).Find(&cards)
	if err!=nil {
		log.Debug("get card keys err:%s",err.Error())
		return false
	}
	keyMap :=make(map[int64]map[string]bool,0)
	for _,v:=range cards{
		keys,b:= keyMap[v.KeyRmb]
		if !b {
			keys = make(map[string]bool,0)
		}
		keys[v.KeyCode] = true
		keyMap[v.KeyRmb] = keys
	}
	for rmb,keys:=range keyMap{
		fileName := fmt.Sprintf("card/keys_%d_%d.txt",rmb,time.Now().Unix())
		go SaveFile(keys,fileName)
	}
	return true
}

func SaveFile(keys map[string]bool,fileName string) bool {
	//打开文件，新建文件
	f, err := os.Create(fileName) //传递文件路径
	if err != nil {           //有错误
		log.Debug("Create file err:",err.Error())
		return false
	}
	//使用完毕，需要关闭文件
	defer f.Close()
	for k,_:=range keys{
		_, err := f.WriteString(k+"\n")
		if err != nil {
			log.Debug("Create err:",err.Error())
			return false
		}
	}
	log.Debug("saved file : %s",fileName)
	return true
}
