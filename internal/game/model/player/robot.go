package player

import "math/rand"

type Robot struct {
	*Player
}

func NewRobot() *Robot {
	robot := &Robot{
		Player: &Player{
			Audience: &Audience{},
			isHuman:  false,
		},
	}
	robot.ID = robot.RandomID()
	robot.Username = robot.RandomName()

	return robot
}

// 隨機名稱
func (r *Robot) RandomName() string {
	return "Robot"
}

// 隨機產生一個id
func (r *Robot) RandomID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
