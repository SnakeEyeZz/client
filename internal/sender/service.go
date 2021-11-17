package sender

type IService interface {
	CreateCommander() *Sender
}

type Service struct {
	pCommander *Sender
}

func (s *Service) CreateCommander() *Sender {
	// s.pCommander = &Sender{}
	return &Sender{}
}
