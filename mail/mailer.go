package mail
import "github.com/lcaballero/kublai/app/handlers"

// Mailer makes an http request to the subscriber, publishing the
// event they wanted to know about.
type Mailer struct {
	id string
	incoming chan handlers.PubEvent,

}



