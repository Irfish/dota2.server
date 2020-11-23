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

	{
	//	handler := NewEinGetClearAllDB()
	//	ret["/get/clear_all_db"] = handler.handle
	}

	{
		handler := NewEginGetReloadBlackList()
		ret["/get/reload_black_list"] = handler.handle
	}

	return ret
}
