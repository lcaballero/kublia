package queue

// PublishProps represent the minimum amount of information required
// to issue an event that needs shipped to subscribers.  Potentially,
// a bulk upload could include several of these entries in a single
// stream.  Additional information, internal to the queue would be
// added for tracking.  For instance and ID and a timestamp
// describing when the event was received.
type PublishProps struct {
	PublisherID int32
	Topic       int32
	PayloadSize int32
	Payload     []byte
}

// ID is created by the application and the timestamp represents when
// the event was received.
type InternalProps struct {
	Id        int32
	Roll      int32 // Indicates log roll the entry came from
	Timestamp int64
}

// TransitionProps capture the state of processing that the event has
// undergone.  There are 2 terminal states: ReadyForArchive and
// ProcessingError.
type TransitionProps struct {
	PreviousState LogMsgState
	CurrentState  LogMsgState
}
