package player

// Audience 接口定义了 Audience 结构体应该实现的方法

// Audience 结构体表示观众
type Audience struct {
	ID           string
	Username     string
	Connection   stream
	InputChannel chan string
}

// NewAudience 创建一个新的观众实例
func NewAudience(id string, username string, connection stream) *Audience {
	return &Audience{
		ID:         id,
		Username:   username,
		Connection: connection,
	}
}

// GetID 返回观众的 ID
func (a *Audience) GetID() string {
	return a.ID
}

// GetUsername 返回观众的用户名
func (a *Audience) GetUsername() string {
	return a.Username
}

// GetConnection 返回观众的连接信息
func (a *Audience) GetConnection() stream {
	return a.Connection
}
