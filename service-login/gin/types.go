package gin

const (
	GAME_STATE_IDLE =iota
	GAME_STATE_ENTER
	GAME_STATE_END
)

const (
	GAME_RESULT_WIN =1
	GAME_RESULT_LOOSE =2
)

const (
	PLAYER_STATE_IDLE =iota
	PLAYER_STATE_IN_GAME
)

const (
	ITEM_STATE_USED =0
	ITEM_STATE_HAVE =2
	ITEM_FROM_BUY =3
	ITEM_FROM_REWARD=4
	ITEM_STATE_USE_30 =30
	ITEM_STATE_USE_10 =10

)

const (
	ERRORCODE_SERVER_ERR = 1 //服务器异常
	ERRORCODE_GAME_NOT_EXIT = 2//游戏ID不存在
	ERRORCODE_PLAYER_NOT_EXIT = 3//玩家不存在
	ERRORCODE_SILVER_NOT_ENOUGH = 4//银币不足
	ERRORCODE_GOLD_NOT_ENOUGH = 5//金币不足
	ERRORCODE_CARD_KEY_USED = 6//卡密无效
	ERRORCODE_ITEM_NOT_EXIT = 7//道具不存在
	ERRORCODE_ITEM_USED = 8//体验卡已经被使用
)

type Player struct {
	Gold int64
	Silver int64
	VipExp int64
	SteamId string
	Items   []Item
	LimitItems   []LimitItem
	UseTime int64
	GameState int
	Index  int
}

type Game struct {
	Players map[string]*Player
	GameID string
	State int
	CreateTime int64
}

type GameEndData struct {
	Steam string
	GameId string
	Score int
	Silver int
}

type GameRank struct {
	SteamID string
	Score  int64
	PlayTime int64
}

type GameResult struct {
	SteamId string
	GameID string
	Score int64
	Silver int64
}

type Item struct {
	Id int64
	Count int64
}

type LimitItem struct {
	Id int64
	LimitTime int //时间
	StartTime int64 //开始使用时间
	LeftTime  int64 //剩余时间
}






