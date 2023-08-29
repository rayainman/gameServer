package player

type ClientInterface interface {
	GetID() string
	GetUsername() string
	GetConnection() stream
}

// 繼承關係 Robot -> Player -> Audience
