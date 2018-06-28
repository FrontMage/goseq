package goseq

// MemSequencer is the memory store impl of Sequencer
type MemSequencer struct {
	SequenceID int64
}

func (s *MemSequencer) Open(startID int64) error {
	s.SequenceID = startID
	return nil
}

func (s *MemSequencer) Next() (int64, error) {
	s.SequenceID++
	return s.SequenceID, nil
}

func (s *MemSequencer) Close() (lastSeqID int64, err error) {
	return s.SequenceID, nil
}
