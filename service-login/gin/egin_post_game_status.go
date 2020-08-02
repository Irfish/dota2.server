package gin

import (
	"fmt"
	"github.com/Irfish/component/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	DOTA_GAMERULES_STATE_IDLE = -1
	DOTA_GAMERULES_STATE_WAIT_FOR_PLAYERS_TO_LOAD=1
	DOTA_GAMERULES_STATE_CUSTOM_GAME_SETUP=2
	DOTA_GAMERULES_STATE_HERO_SELECTION=3
	DOTA_GAMERULES_STATE_PRE_GAME=4
	DOTA_GAMERULES_STATE_GAME_IN_PROGRESS=5
)

type GameStatus struct {
	 CurState int //当前游戏转态
}

func NewGameStatus() GameStatus {
	p := GameStatus{CurState:DOTA_GAMERULES_STATE_IDLE}
	return p
}

func (p *GameStatus) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
	gameState, ok := c.GetPostForm("gameState")
	if !ok {
		e = fmt.Errorf("the key gameState not found")
		return
	}
	state, err := strconv.Atoi(gameState)
	if err != nil {
		e = fmt.Errorf("gameState can not convert to int")
		return
	}
	log.Debug("game rule state %d:",state)
	info:="game rule state : "
	switch state {
	case DOTA_GAMERULES_STATE_WAIT_FOR_PLAYERS_TO_LOAD:
		info+="wait for players to load"
		break
	case DOTA_GAMERULES_STATE_CUSTOM_GAME_SETUP:
		info+="custom game setup"
		break
	case DOTA_GAMERULES_STATE_HERO_SELECTION:
		info+="hero selection"
		break
	case DOTA_GAMERULES_STATE_PRE_GAME:
		info+="pre game"
		break
	case DOTA_GAMERULES_STATE_GAME_IN_PROGRESS:
		info+="game in progress"
		break
	default:
		info+=fmt.Sprintf("idle (state value %d)",state)
		break
	}
	result["status"] = info
}