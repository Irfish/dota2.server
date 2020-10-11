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
		}
		c.JSON(http.StatusOK, result)
	}()
	steamID := GetStringFromPostForm(c,"steamId")
	id := GetInt64FromPostForm(c,"itemId")
	count :=GetInt64FromPostForm(c,"count")
	PlayerUseItem(steamID,id, count)
}