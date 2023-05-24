package saveservice

type SaveEntry[GameState_T IGameState] struct {
	state GameState_T
}

type IGameState struct {
}

func (s *SaveEntry[GameState_T]) UpdateState(newState GameState_T) {
	s.state = newState
}
