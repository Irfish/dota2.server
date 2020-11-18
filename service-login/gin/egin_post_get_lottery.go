package gin

import (
	"github.com/Irfish/component/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetLottery struct {
}

func NewGetLottery() GetLottery {
	p := GetLottery{}
	return p
}

func (p *GetLottery) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
			result["errCode"] = ERRORCODE_SERVER_ERR
		}
		log.Debug("result: %v",result)
		c.JSON(http.StatusOK, result)
	}()
	gameID := GetStringFromPostForm(c,"gameId")
	if b,i := CheckGameID(gameID);!b {
		result["errCode"] = i
		return
	}

	steamID := GetStringFromPostForm(c,"steamId")
	if b,i := CheckSteamID(gameID,steamID);!b {
		result["errCode"] = i
		return
	}

	lotteryType := GetInt64FromPostForm(c,"lotteryType")
	lotteryId := 0
	lotteryCount :=0
	switch lotteryType {
	case LOTTERY_TYPE_SLIVER:
		if b,i := CheckSilver(steamID,50);!b {
			result["errCode"] = i
			return
		}
		if b, id, c := lotteryManager.GetSilverLottery(steamID);b {
			lotteryId = id
			lotteryCount=c
		}else {
			result["errCode"] = ERRORCODE_SERVER_ERR
			return
		}
		break
	case LOTTERY_TYPE_GOLD:
		if b,i := CheckGold(steamID,50);!b {
			result["errCode"] = i
			return
		}
		if b, id,c := lotteryManager.GetGoldLottery(steamID);b {
			lotteryId = id
			lotteryCount=c
		}else {
			result["errCode"] = ERRORCODE_SERVER_ERR
			return
		}
		break
	}
	result["lotteryId"] =lotteryId
	result["lotteryCount"] =lotteryCount
	gameManager.RefreshPlayer(gameID,steamID)
	player:= gameManager.GetPlayer(gameID,steamID)
	result["player"] = player
}