package gin

const (
	GAME_STATE_IDLE =iota
	GAME_STATE_ENTER
	GAME_STATE_END
)

const (
	PLAYER_STATE_IDLE =iota
	PLAYER_STATE_IN_GAME
)

const (
	ITEM_STATE_NULL =iota
	ITEM_STATE_HAVE

)

const (
	ITEM_FROM_BUY =1
	ITEM_FROM_REWARD=2
)

type Player struct {
	Gold int64
	Silver int64
	VipExp int64
	SteamId string
	Items   []int32
	UseTime int64
	GameState int
	Index  int
}

type Game struct {
	Players map[string]*Player
	GameID int64
	State int
	CreateTime int64
}

type GameEndData struct {
	Steam string
	GameId int64
	Score int
	Silver int
}





