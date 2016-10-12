package channels
import "github.com/lcaballero/kublai/app/handlers"


type Channels struct {
	PublishedToHandler chan handlers.PubEvent
	OnCompletedEvents chan handlers.PubEvent
	OnPersistedEvents chan struct{}
}


func NewChannels() *Channels {
	return &Channels{
		PublishedToHandler: make(chan handlers.PubEvent, 1000),
		OnCompletedEvents: make(chan handlers.PubEvent, 1000),
		OnPersistedEvents: make(chan struct{}),
	}
}

func (c *Channels) Close() error {
	close(c.PublishedToHandler)
	close(c.OnCompletedEvents)
	close(c.OnPersistedEvents)
	return nil
}