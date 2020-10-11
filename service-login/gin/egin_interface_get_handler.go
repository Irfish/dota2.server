package gin

import "github.com/gin-gonic/gin"

func (s *Gin) GinGetHandler() map[string]func(*gin.Context) {
	ret := make(map[string]func(*gin.Context))
	{
		handler := NewServiceTime()
		ret["/get/serviceTime"] = handler.handle
	}
	{
		handler := NewGenCardKey()
		ret["/get/gen_card_key"] = handler.handle
	}
	return ret
}
