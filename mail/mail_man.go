package mail
import "github.com/lcaballero/hitman"


type MailMan struct {
	onPersistedWake chan struct{}
}

func NewMailMan(
	onPersistedWake chan struct{},
	) *MailMan {

	m := &MailMan{
		onPersistedWake: onPersistedWake,
	}
	return m
}

func (m *MailMan) Name() string {
	return "MailMan"
}

func (m *MailMan) Start() hitman.KillChannel {
	kill := hitman.KillChannel()
	go func() {
		for {
			select {
			case cleaner := <-kill:
				cleaner.WaitGroup.Done()
				return

			case <-m.onPersistedWake:
				// there's work to be done
			}
		}
	}()
	return kill
}

