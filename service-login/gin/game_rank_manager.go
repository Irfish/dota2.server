package gin

import (
	"github.com/Irfish/component/log"
	"github.com/Irfish/component/redis"
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
	LoadConfigBlackUserList()
	log.Debug("GameRankManagers:Update %d %v",gameLevel,ranks)
	l := m.GameLevelToRankList[gameLevel]
	l.Update(ranks)
	_, e1 := redis.RedisHset(RANK_LIST_REDIS_KEY, "info", gameRankManagers)
	if e1 != nil {
		log.Debug(e1.Error())
	}
}

func (m *GameRankManagers)GetList() map[int]*GameRankManager {
	return m.GameLevelToRankList
}

type GameRankManager struct {
	RankList  []GameRank
}

func NewGameRankManager()  *GameRankManager {
	m := &GameRankManager{}
	m.RankList = make([]GameRank,0)
	return m
}

func (m *GameRankManager)GetList() []GameRank{
	m.Update([]GameRank{})
	return m.RankList
}

func (m *GameRankManager)Update(ranks []GameRank)  {
	log.Debug("GameRankManagers:Update before %v",m.RankList)
	m.RankList = append(m.RankList,ranks...)
	list := make(map[string]GameRank,0)
	for _,r:=range m.RankList {
		isBlack :=false
		for _,black:=range blackUserList{
			if r.SteamID ==black.SteamId {
				isBlack =true
				break
			}
		}
		if !isBlack {
			s,b:= list[r.SteamID]
			if !b {
				list[r.SteamID] = r
			}else {
				if s.PlayTime>r.PlayTime || (s.PlayTime==r.PlayTime && s.Score<r.Score){
					list[r.SteamID] =r
				}
			}
		}
	}
	m.RankList = make([]GameRank,0)
	for _,r:=range list {
		m.RankList = append(m.RankList,r)
	}
	sort.Sort(m)
	len:= m.Len()
	if len > RANK_MAX_NUM {
		m.RankList = m.RankList[:len-RANK_MAX_NUM-1]
	}
	log.Debug("GameRankManagers:Update after %v",m.RankList)
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

