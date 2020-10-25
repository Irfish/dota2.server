package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UseItem struct {
}

func NewUseItem() UseItem {
	p := UseItem{}
	return p
}

func (p *UseItem) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
			result["errCode"] = ERRORCODE_SERVER_ERR
		}
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

	id := GetInt64FromPostForm(c,"itemId")
	count :=GetInt64FromPostForm(c,"count")

	player:= gameManager.GetPlayer(gameID,steamID)
	if len(player.LimitItems)>0 {
		result["errCode"] = ERRORCODE_ITEM_USED
		return
	}
	if b,i:= PlayerUseItem(steamID,id, count);!b {
		result["errCode"] = i
		return
	}

	gameManager.RefreshPlayer(gameID,steamID)
	player= gameManager.GetPlayer(gameID,steamID)
	result["player"] = player
}