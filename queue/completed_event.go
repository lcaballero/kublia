package queue


// CompletedEvent is fixed width so that once the ID is determined
// the position in the file can be calculated and so PreviousOffset,
// and CurrentOffset can also be calculated based on the ID.
type CompletedEvent struct {
	PublishProps
	InternalProps
	TransitionProps
}
