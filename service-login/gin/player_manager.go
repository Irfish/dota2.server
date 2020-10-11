package gin

import (
	"fmt"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/xorm"
	"github.com/Irfish/dota2.server/service-login/orm"
	"time"
)

type PlayerManager struct {
	Players map[string]*Player
}


func GetAndCreate(id string) (bool,orm.User ) {
	exit,u := GetUser(id)
	if exit {
		return  true,u
	}else {
		state,u:= NewUser(id)
		if state {
			return true,u
		}else {
			log.Debug("insert user failed:")
			return false,orm.User{}
		}
	}
	log.Debug("GetAndCreate user failed:")
	return false,orm.User{}
}

func GetUser(steamId string)  (bool,orm.User ) {
	u := orm.User{}
	exist, err := orm.UserXorm().Where("steam_id=?",steamId).Get(&u)
	if err !=nil {
		log.Debug("do get user err:%s",err.Error())
		return false,orm.User{}
	}
	if exist {
		return true,u
	}
	return false,orm.User{}
}

func NewUser(steamId string)  (bool,orm.User ) {
	t := time.Now().Unix()
	u := orm.User{
		SteamId: steamId,
		SteamGold:0,
		SteamSilver:0,
		SteamVipExp:0,
		CreateTime: t,
		UpdateTime: t,
	}
	affected, err := orm.UserXorm().Insert(u)
	if err !=nil {
		log.Debug("insert user err:%s",err.Error())
	}
	if affected>0 {
		return true,u
	}
	return false,orm.User{}
}

func GetItems(steamId string) []int32 {
	exit,u := GetUser(steamId)
	if exit {
		userID:= u.Id
		userBags := []orm.UserBag{}
        count,err:=orm.UserBagXorm().Where("user_id=? and item_state=?",userID,ITEM_STATE_HAVE).FindAndCount(&userBags)
		if err !=nil {
			log.Debug("insert user err:%s",err.Error())
		}
		if count>0 {
			items := make([]int32,count)
			for i,item := range userBags{
				items[i] = int32(item.ItemId)
			}
			return items
		}
	}
	return  []int32{}
}

func PlayerBuyItem(steamId string,itemId int64,cost int64,count int64) bool {
	t := time.Now().Unix()
	exit,u := GetUser(steamId)
	if !exit {
		log.Debug("user not found")
		return false
	}
	if u.SteamGold < cost {
		log.Debug("gold not enough")
		return false
	}
	var e error
	s := xorm.Xorm(0).NewSession()
	defer func() {
		if e!=nil {
			log.Debug("%s",e.Error())
			s.Rollback()
		}
		s.Close()
	}()
	if err := s.Begin() ; err != nil {
		e = fmt.Errorf("fail to session begin")
		return false
	}

	momey := u.SteamGold-cost
	user:= orm.User{
		SteamGold:momey,
		UpdateTime: t,
	}
	//更新玩家金币
	{
		s1:= s.Table("user")
		effect,err:= s1.ID(u.Id).Update(&user)
		if err !=nil {
			e = fmt.Errorf(err.Error())
			return false
		}
		if effect!=1 {
			e = fmt.Errorf("update gold failed in user table")
			return false
		}
	}
	//将道具存入背包
	{
		userBag:= orm.UserBag{
					UserId: u.Id,
					ItemId: itemId,
					ItemCount: count,
					ItemState: ITEM_STATE_HAVE,
					CreateTime: t,
					UpdateTime: t,
		}

		s2:= s.Table("user_bag")
		effect,err:= s2.Insert(&userBag)
		if err !=nil {
			e = fmt.Errorf(err.Error())
			return false
		}
		if effect!=1 {
			e = fmt.Errorf("insert buy itme failed in userbag table")
			return false
		}
		s.Commit()
	}
	//记录使用金币兑换道具日志
	{
		logUseItem := orm.LogUseItem{
			CreateTime: t,
			UpdateTime: t,
			ItemId: itemId,
			UserId: u.Id,
			ItemUseState:ITEM_FROM_BUY,
		}
		_,err:=orm.LogUseItemXorm().Insert(&logUseItem)
		if err !=nil {
			log.Debug("%s",err.Error())
		}
	}
	return true
}


func PlayerUseItem(steamId string,itemId int64,count int64)  {

}

func PlayerUseCardKey(steamId string,code string) bool {
	card := orm.CardKey{}
	exit,err:=orm.CardKeyXorm().Where(fmt.Sprintf("key_code=%s",code)).And("key_state=?",KEY_STATE_NORMAL) .Get(&card)
	if err!=nil {
		log.Debug("do sql get err:%s",err.Error())
		return false
	}
	if !exit {
		log.Debug("code(%s)  not exit",code)
		return false
	}
	exit, u := GetUser(steamId)
	if !exit {
		log.Debug("steamId(%s) not exit",steamId)
		return false
	}
	t := time.Now().Unix()

	var e error
	s := xorm.Xorm(0).NewSession()
	defer func() {
		if e!=nil {
			log.Debug("%s",e.Error())
			s.Rollback()
		}
		s.Close()
	}()
	if err := s.Begin() ; err != nil {
		e = fmt.Errorf("fail to session begin")
		return false
	}
	s1:= s.Table("card_key")

	//更新code状态
	{
		card.KeyState = KEY_STATE_USED
		card.UpdateTime= t
		effect,err:= s1.Where(fmt.Sprintf("key_code=%s",code)).Update(&card)
		if err!=nil {
			e = fmt.Errorf("%s",err.Error())
			return false
		}
		if effect !=1 {
			e = fmt.Errorf("update key_code state failed")
			return false
		}
	}
	//更新玩家金币
	{
		s2:= s.Table("user")
		u.UpdateTime = t
		u.SteamGold = u.SteamGold+card.KeyValue
		effect,err:= s2.Where("id=?",u.Id).Update(&u)
		if err!=nil {
			e = fmt.Errorf("%s",err.Error())
			return false
		}
		if effect !=1 {
			e = fmt.Errorf("update steam_gold state failed")
			return false
		}
	}
	s.Commit()

	//记录code使用日志
	{
		log:= orm.LogCardKey{}
		log.UpdateTime = t
		log.CreateTime = t
		log.UserId = u.Id
		log.CardKeyId= card.Id
		log.Cost = card.KeyRmb
		log.UseTime = 0
		effect,err:= orm.LogCardKeyXorm().Insert(&log)
		if err !=nil {
			e = fmt.Errorf("%s",err.Error())
			return false
		}
		if effect !=1 {
			e = fmt.Errorf("insert  card key log failed")
			return false
		}
	}

	return true
}



