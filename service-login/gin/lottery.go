package gin

import (
	"fmt"
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
	"math/rand"
	"time"
)

var (
	lotteryManager =NewLottery()
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

const (
	LOTTERY_TYPE_SLIVER =1 //银币抽奖
	LOTTERY_TYPE_GOLD =2 //金币抽奖
)

const (
	PLAYER_LOTTERY_COUNT_REDIS_KEY = "service.login.player_lottery_count_redis_key"
)

type PlayerLottery struct {
	CountGold int
	CountSilver int
	Id int64
}

func GetPlayerLotteryKey(steamID string) string {
	p:= gameManager.GetUser(steamID)
	if p!=nil {
		return fmt.Sprintf("%s_%d", PLAYER_LOTTERY_COUNT_REDIS_KEY,p.ID)
	}
	return ""
}

func GetPlayerLotteryCount(steamID string) *PlayerLottery {
	key := GetPlayerLotteryKey(steamID)
	if key=="" {
		return &PlayerLottery{CountSilver: 0,CountGold: 0}
	}
	r, err := redis.RedisHget(key, "count")
	if err != nil {
		log.Debug(err.Error())
		return &PlayerLottery{CountSilver: 0,CountGold: 0}
	}
	if r == nil {
		log.Debug("redis.RedisHget key :%s  is nil", key)
		return &PlayerLottery{CountSilver: 0,CountGold: 0}
	}else {
		p:= r.(*PlayerLottery)
		return p
	}
	return &PlayerLottery{CountSilver: 0,CountGold: 0}
}

func UpdatePlayerLotteryCount(lotteryType int, steamID string,reset bool)  {
	p := GetPlayerLotteryCount(steamID)
	if p==nil {
		return
	}

	if lotteryType==LOTTERY_TYPE_GOLD {
		if reset {
			p.CountGold = 1
		}else {
			p.CountGold = p.CountGold+1
		}
	}else {
		if reset {
			p.CountSilver = 1
		}else {
			p.CountSilver = p.CountSilver+1
		}
	}
	key := GetPlayerLotteryKey(steamID)
	if key!="" {
		_, e1 := redis.RedisHset(key, "count", p)
		if e1 != nil {
			log.Debug(e1.Error())
		}
	}
}

func SliceOutOfOrder(in []int) []int {
	l := len(in)
	for i := l - 1; i > 0; i-- {
		j := r.Intn(i)
		in[j], in[i] = in[i], in[j]
	}
	return in
}

type LotteryConfig struct {
	Index  int
	Type  int
	Rate  int
}

type LotteryManager struct {
	SilverSeed []int
	GoldSeed []int
}

func NewLottery()  *LotteryManager {
	l := &LotteryManager{}
	{
		var silverConfig [14]LotteryConfig
		silverConfig[0] = LotteryConfig{0,0,0}
		silverConfig[1] = LotteryConfig{1,1,50}
		silverConfig[2] = LotteryConfig{2,2,50}
		silverConfig[3] = LotteryConfig{3,3,25}
		silverConfig[4] = LotteryConfig{4,4,25}
		silverConfig[5] = LotteryConfig{5,5,50}
		silverConfig[6] = LotteryConfig{6,6,50}
		silverConfig[7] = LotteryConfig{7,7,150}
		silverConfig[8] = LotteryConfig{8,8,125}
		silverConfig[9] = LotteryConfig{9,9,100}
		silverConfig[10] = LotteryConfig{10,10,75}
		silverConfig[11] = LotteryConfig{11,11,50}
		silverConfig[12] = LotteryConfig{12,12,150}
		silverConfig[13] = LotteryConfig{13,13,100}

		l.SilverSeed =make([]int,0)
		for i:=1;i<14;i++ {
			for j:=0;j<silverConfig[i].Rate;j++ {
				l.SilverSeed = append(l.SilverSeed,silverConfig[i].Type)
			}
		}
		l.SilverSeed = SliceOutOfOrder(l.SilverSeed)
	}

	{
		var goldConfig [14]LotteryConfig
		goldConfig[0] = LotteryConfig{0,0,0}
		goldConfig[1] = LotteryConfig{1,1,20}
		goldConfig[2] = LotteryConfig{2,2,20}
		goldConfig[3] = LotteryConfig{3,3,85}
		goldConfig[4] = LotteryConfig{4,4,85}
		goldConfig[5] = LotteryConfig{5,5,85}
		goldConfig[6] = LotteryConfig{6,6,85}
		goldConfig[7] = LotteryConfig{7,7,85}
		goldConfig[8] = LotteryConfig{8,8,85}
		goldConfig[9] = LotteryConfig{9,9,140}
		goldConfig[10] = LotteryConfig{10,10,130}
		goldConfig[11] = LotteryConfig{11,11,80}
		goldConfig[12] = LotteryConfig{12,12,60}
		goldConfig[13] = LotteryConfig{13,13,40}

		l.GoldSeed =make([]int,0)
		for i:=1;i<14;i++ {
			for j:=0;j<goldConfig[i].Rate;j++ {
				l.GoldSeed = append(l.GoldSeed,goldConfig[i].Type)
			}
		}
		l.GoldSeed = SliceOutOfOrder(l.GoldSeed)
	}
	return l
}

func (m *LotteryManager)GetGoldLottery(steamID string) (bool, int,int) {
	if !UpdateGold(steamID,50) {
		return false,0,0
	}

	m.SilverSeed = SliceOutOfOrder(m.SilverSeed)
	l :=m.SilverSeed[0]
	p := GetPlayerLotteryCount(steamID)
	count := 0
	if p!=nil {
		count =p.CountGold
	}
	count= count+1
	if count==20 {
		l =0
	}
	itemId := int64(0)
	switch l {
	case 0://
		itemId = 28
		break
	case 1://高级信使 22,23
		itemId = 22
		break
	case 2://高级信使 22,23
		itemId = 23
		break
	case 3://高级英雄 13,14,15,16,17,18
		itemId = 15
		break
	case 4://高级英雄 13,14,15,16,17,18
		itemId = 18
		break
	case 5://中级英雄 7，8，9，10，11，12
		itemId = 9
		break
	case 6://中级英雄 7，8，9，10，11，12
		itemId = 10
		break
	case 7://vip  200
		if !UpdateExp(steamID,200) {
			return false,0,0
		}
		break
	case 8://vip  300
		if !UpdateExp(steamID,300) {
			return false,0,0
		}
		break
	case 9://vip  400
		if !UpdateExp(steamID,400) {
			return false,0,0
		}
		break
	case 10://vip  500
		if !UpdateExp(steamID,500) {
			return false,0,0
		}
		break
	case 11://vip  600
		if !UpdateExp(steamID,600) {
			return false,0,0
		}
		break
	case 12://银币  300
		if !UpdateSilver(steamID,-300) {
			return false,0,0
		}
		break
	case 13://银币  500
		if !UpdateSilver(steamID,-500) {
			return false,0,0
		}
		break
	}
	if itemId!=0 {
		if !UpdateItem(steamID,itemId,1) {
			return false,0,0
		}
	}
	UpdatePlayerLotteryCount(LOTTERY_TYPE_GOLD,steamID,count==20)
	return true,l,count
}

func (m *LotteryManager)GetSilverLottery(steamID string) (bool,int,int) {
	if !UpdateSilver(steamID,50) {
		return false,0,0
	}
	m.GoldSeed = SliceOutOfOrder(m.GoldSeed)
	l :=m.GoldSeed[0]
	p := GetPlayerLotteryCount(steamID)
	count := 0
	if p!=nil {
		count =p.CountSilver
	}
	count= count+1
	itemId := int64(0)
	switch l {
	case 0://
		itemId = 27
		break
	case 1://中级信使 20,21
		itemId = 20
		break
	case 2://中级信使 20,21
		itemId = 21
		break
	case 3://低级英雄 1，2，3，4，5，6
		itemId = 1
		break
	case 4://低级英雄 1，2，3，4，5，6
		itemId = 2
		break
	case 5://低级英雄 1，2，3，4，5，6
		itemId = 3
		break
	case 6://低级英雄 1，2，3，4，5，6
		itemId = 4
		break
	case 7://低级英雄 1，2，3，4，5，6
		itemId = 5
		break
	case 8://低级英雄 1，2，3，4，5，6
		itemId = 6
		break
	case 9://vip  50
		if !UpdateExp(steamID,50) {
			return false,0,0
		}
		break
	case 10://vip  75
		if !UpdateExp(steamID,75) {
			return false,0,0
		}
		break
	case 11://vip  100
		if !UpdateExp(steamID,100) {
			return false,0,0
		}
		break
	case 12://vip  125
		if !UpdateExp(steamID,125) {
			return false,0,0
		}
		break
	case 13://vip  150
		if !UpdateExp(steamID,150) {
			return false,0,0
		}
		break
	}
	if itemId!=0 {
		if !UpdateItem(steamID,itemId,1) {
			return false,0,0
		}
	}
	UpdatePlayerLotteryCount(LOTTERY_TYPE_SLIVER,steamID,count==20)
	return true,l,count
}
