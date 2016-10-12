package queue

import (
	"time"
)

// PublishEvent represents the header for a single publishing event
// while the api might only require the Topic, PublisherId, PayloadSize
// and Payload itself.
type PublishEvent struct {
	PublishProps
	InternalProps
	TransitionProps
	PreviousOffset int32
	CurrentOffset  int32
}

// ToBinary packs the header into a byte slice.
func (a *PublishEvent) ToBinary() []byte {
	time.Now().UnixNano()
	return []byte{}
}

func (a *PublishEvent) FromBinary(offset int, bin []byte) error {
	return nil
}
