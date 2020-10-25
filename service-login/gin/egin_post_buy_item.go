package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BuyItem struct {
}

func NewBuyItem() BuyItem {
	p := BuyItem{}
	return p
}

func (p *BuyItem) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
			result["errCode"] = ERRORCODE_SERVER_ERR
		}
		c.JSON(http.StatusOK, result)
	}()

	gameID := GetStringFromPostForm(c,"gameID")
	if b,i := CheckGameID(gameID);!b {
		result["errCode"] = i
		return
	}

	steamID := GetStringFromPostForm(c,"steamId")
	if b,i := CheckSteamID(gameID,steamID);!b {
		result["errCode"] = i
		return
	}

	cost :=GetInt64FromPostForm(c,"cost")
	if b,i := CheckGold(steamID,cost);!b {
		result["errCode"] = i
		return
	}

	id := GetInt64FromPostForm(c,"itemId")
	count :=GetInt64FromPostForm(c,"count")
	if PlayerBuyItem(steamID,id,cost,count) {
		gameManager.RefreshPlayer(gameID,steamID)
		p:= gameManager.GetPlayer(gameID,steamID)
		result["player"] = p
		return
	}else {
		e = fmt.Errorf("buy item failed")
		return
	}
}