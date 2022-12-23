package store

type Store interface {
	Polls() PollStore
	Players() PlayerStore
	VoteRecords() VoteRecordStore
}
