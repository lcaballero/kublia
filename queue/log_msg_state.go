package queue

// Messages should normally (happy path) proceed through these states in this
// order:
//
//     Published -> Completed -> ReadyForArchive
//
// or when something goes wrong:
//
//     Published -> ProcessingError
type LogMsgState int32

const (
	Received LogMsgState = iota
	Published
	Completed
	ReadyForArchive
	ProcessingError
)
