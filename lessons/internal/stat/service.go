package stat

import (
	"go/adv-demo/pkg/event"
	"log"
)

type StatService struct {
	EventBus *event.EventBus
	StatRepository *StatRepository
}

type StatServiceDeps struct {
	EventBus *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(deps StatServiceDeps) *StatService {
	return &StatService{
		EventBus: deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Println("Bad EventLinkVisited Data ", msg.Data)
			}
			s.StatRepository.AddClick(id)
		}
	}
}
