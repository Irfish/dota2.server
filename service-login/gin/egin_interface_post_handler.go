package gin

import "github.com/gin-gonic/gin"

func (s *Gin) GinPostHandler() map[string]func(*gin.Context) {
	ret := make(map[string]func(*gin.Context), 0)
	{
		handler := NewLoginByAccount()
		ret["/post/login"] = handler.handle
	}
	{
		handler := NewUserRegister()
		ret["/post/user_register"] = handler.handle
	}
	{
		handler := NewGameStatus()
		ret["/post/game_status"] = handler.handle
	}
	{
		handler := NewPayPalOauth()
		ret["/post/pay_oauth"] = handler.handle
	}
	{
	 	handler := NewPayPalOrderCreate()
		ret["/post/pay_order_create"] = handler.handle
	}
	{
		handler := NewEnterGame()
		ret["/post/enter_game"] = handler.handle
	}
	{
		handler := NewGameEnd()
		ret["/post/game_end"] = handler.handle
	}
	{
		handler := NewBuyItem()
		ret["/post/buy_item"] = handler.handle
	}
	{
		handler := NewUseCardKey()
		ret["/post/use_card_key"] = handler.handle
	}
	{
		handler := NewUseItem()
		ret["/post/use_item"] = handler.handle
	}
	{
		handler := NewRankList()
		ret["/post/rank_list"] = handler.handle
	}
	return ret
}
