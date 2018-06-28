package goseq

// Sequencer is a sequence id generater
type Sequencer interface {
	Open(startID int64) error
	Next() (int64, error)
	Close() (lastSeqID int64, err error)
}
