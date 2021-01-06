package window

type Event int8

const EventCount = 4

const (
	SUCCESS   Event = 0
	FAILURE   Event = 1
	TIMEOUT   Event = 2
	REJECTION Event = 3
)
