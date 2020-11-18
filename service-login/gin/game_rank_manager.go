package gin

import (
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
	"github.com/Irfish/component/str"
	"math/rand"
	"sort"
)

var gameRankManagers = NewGameRankManagers()

const (
	RANK_MAX_NUM = 20
	RANK_LIST_REDIS_KEY = "service.login.gameRankManager"
)

type GameRankManagers struct {
	GameLevelToRankList map[int]*GameRankManager
}

func NewGameRankManagers()  *GameRankManagers {
	m := &GameRankManagers{}
	m.GameLevelToRankList = make(map[int]*GameRankManager,0)
	m.GameLevelToRankList[1] = NewGameRankManager()
	m.GameLevelToRankList[2] = NewGameRankManager()
	m.GameLevelToRankList[3] = NewGameRankManager()
	m.GameLevelToRankList[4] = NewGameRankManager()
	return m
}

func (m *GameRankManagers)Update(gameLevel int,ranks []GameRank )  {
	l := m.GameLevelToRankList[gameLevel]
	l.Update(ranks)
	_, e1 := redis.RedisHset(RANK_LIST_REDIS_KEY, "info", gameRankManagers)
	if e1 != nil {
		log.Debug(e1.Error())
	}
}

func (m *GameRankManagers)GetList() map[int]*GameRankManager {
	m.Update(1,[]GameRank{})
	m.Update(2,[]GameRank{})
	m.Update(3,[]GameRank{})
	m.Update(4,[]GameRank{})
	return m.GameLevelToRankList
}

type GameRankManager struct {
	RankList  []GameRank
}

func NewGameRankManager()  *GameRankManager {
	m := &GameRankManager{}
	m.RankList = make([]GameRank,0)
	for i:=0;i<RANK_MAX_NUM;i++ {
		r:=GameRank{
			SteamID:str.RandomNumber(8),
			Score:int64(rand.Intn(1000)),
			PlayTime: 99999999 + int64(rand.Intn(1000)),
		}
		m.RankList = append(m.RankList,r)
	}
	sort.Sort(m)
	return m
}

func (m *GameRankManager)GetList() []GameRank{
	m.Update([]GameRank{})
	return m.RankList
}

func (m *GameRankManager)Update(ranks []GameRank)  {
	m.RankList = append(m.RankList,ranks...)
	list := make(map[string]GameRank,0)
	for _,r:=range m.RankList {
		s,b:= list[r.SteamID]
		if !b {
			list[r.SteamID] = r
		}else {
			if s.PlayTime>r.PlayTime || (s.PlayTime==r.PlayTime && s.Score<r.Score){
				list[r.SteamID] =r
			}
		}
	}
	m.RankList = make([]GameRank,0)
	for _,r:=range list {
		m.RankList = append(m.RankList,r)
	}
	sort.Sort(m)
	len:=m.Len()
	if len> RANK_MAX_NUM {
		m.RankList = m.RankList[:len-RANK_MAX_NUM-1]
	}
}

func (m *GameRankManager) GetRankedSteamIds() []string {
	ids := make([]string,0)
	for _,v:=range m.RankList{
		ids=append(ids,v.SteamID)
	}
	return ids
}


func (m *GameRankManager) Len() int {
	return len(m.RankList)
}

func (m *GameRankManager) Less(i, j int) bool {
	if m.RankList[i].PlayTime == m.RankList[j].PlayTime {
		return m.RankList[i].Score > m.RankList[j].Score
	}
	return m.RankList[i].PlayTime < m.RankList[j].PlayTime
}

func (m *GameRankManager) Swap(i, j int) {
	m.RankList[i], m.RankList[j] = m.RankList[j], m.RankList[i]
}

func LoadGameRankManager()  {
	r, err := redis.RedisHget(RANK_LIST_REDIS_KEY, "info")
	if err != nil {
		log.Debug(err.Error())
		return
	}
	if r == nil {
		log.Debug("redis.RedisHget key :%s  is nil",RANK_LIST_REDIS_KEY)
		gameRankManagers = NewGameRankManagers()
	}else {
		info := r.(*GameRankManagers)
		gameRankManagers = info
	}
}

