package allscreens

type Screen int
const (
	Search Screen = iota
	Result
	Detail
	Watchlist
)

type Model struct {
	CurrentScreen Screen
}

func InitializeScreen() Model {
	return Model{
		CurrentScreen: Search,
	}
}

func (m *Model) SetScreen(screen Screen) {
	m.CurrentScreen = screen
}

