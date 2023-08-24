package player

// Status represents the current status of a player in a game.
type Status int

const (
	// StatusUnknown is the default status of a player.
	StatusUnknown Status = iota
	// StatusFolded indicates that a player has folded.
	StatusFolded
	// StatusChecked indicates that a player has checked.
	StatusChecked
	// StatusCalled indicates that a player has called.
	StatusCalled
	// StatusRaised indicates that a player has raised.
	StatusRaised
	// StatusAllIn indicates that a player has gone all-in.
	StatusAllIn
	// StatusWon indicates that a player has won the game.
	StatusWon
	// StatusLost indicates that a player has lost the game.
	StatusLost
)

// String returns the string representation of a player status.
func (s Status) String() string {
	switch s {
	case StatusFolded:
		return "folded"
	case StatusChecked:
		return "checked"
	case StatusCalled:
		return "called"
	case StatusRaised:
		return "raised"
	case StatusAllIn:
		return "all-in"
	case StatusWon:
		return "won"
	case StatusLost:
		return "lost"
	default:
		return "unknown"
	}
}

// 確定玩家出局了沒
func (s Status) IsOut() bool {
	switch s {
	case StatusFolded, StatusLost:
		return true
	default:
		return false
	}
}
