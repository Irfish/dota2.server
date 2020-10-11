package gin

import (
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
		}
		c.JSON(http.StatusOK, result)
	}()
}