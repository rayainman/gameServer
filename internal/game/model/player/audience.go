package player

type Audience struct {
	ID           string
	Username     string
	connection   stream
	InputChannel chan string
}

func NewAudience(id string, username string, connection stream) *Audience {
	return &Audience{
		ID:         id,
		Username:   username,
		connection: connection,
	}
}
