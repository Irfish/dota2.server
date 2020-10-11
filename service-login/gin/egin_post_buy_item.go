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
		}
		c.JSON(http.StatusOK, result)
	}()
	steamID := GetStringFromPostForm(c,"steamId")
	gameID := GetInt64FromPostForm(c,"gameID")

	id := GetInt64FromPostForm(c,"itemId")
	cost :=GetInt64FromPostForm(c,"cost")
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