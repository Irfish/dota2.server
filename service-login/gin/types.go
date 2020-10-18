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
	Items   []Item
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






