package gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PayPalOrderCreate struct {

}

func NewPayPalOrderCreate() PayPalOrderCreate {
	p := PayPalOrderCreate{}
	return p
}

func (p *PayPalOrderCreate) handle(c *gin.Context) {
	var e error
	result := gin.H{}
	defer func() {
		if e != nil {
			result["err"] = e.Error()
		}
		c.JSON(http.StatusOK, result)
	}()
}
