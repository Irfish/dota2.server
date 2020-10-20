package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RankList struct {
}

func NewRankList() RankList {
	p := RankList{}
	return p
}

func (p *RankList) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
	players:= gameRankManager.GetList()
	result["rank"] =players
}
