package appender
import (
	"github.com/lcaballero/hitman"
	"github.com/lcaballero/kublai/app/handlers"
	"time"
)

type Appender struct {
	incomingPubEvents chan handlers.PubEvent
	onCompletedEvent chan handlers.PubEvent
	onPersistedEvent chan struct{}
}

func NewAppender(
	publishedToHandler chan handlers.PubEvent,
	onCompletedEvent chan handlers.PubEvent,
	onPersistedEvents chan struct{},
	) *Appender {

	return &Appender{
		incomingPubEvents: publishedToHandler,
		onCompletedEvent: onCompletedEvent,
		onPersistedEvent: onPersistedEvents,
	}
}

func (a *Appender) Name() string {
	return "Appender"
}

func (a *Appender) Start() hitman.KillChannel {
	kill := hitman.NewKillChannel()
	go func() {
		flushTic := time.NewTicker(1*time.Second)

		for {
			select {
			case cleaner := <-kill:
				cleaner.WaitGroup.Done()
				return

			case <-flushTic.C:
				// flush events to disk based on time

			case <-a.onCompletedEvent:
				// completed event occurs once event is on disk

			case <-a.incomingPubEvents:
				// flush events to disk based on size
			}
		}
	}()
	return kill
}

func (a *Appender) informOfEvents() {
	// Inform at least once that events have been persisted and so these
	// events can be pushed to subscribers.
	select {
	case a.onPersistedEvent <- struct{}{}:
	default:
		return
	}
}

