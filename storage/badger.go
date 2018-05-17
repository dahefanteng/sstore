package badger

import (
	pb "github.com/coreos/etcd/raft/raftpb"
	"github.com/dgraph-io/badger"
	"log"
)

type Badger struct {
	DB       *badger.DB
	Dir      string
	ValueDir string
}

func NewBadger(dir, valueDir string) *Badger {
	opts := badger.DefaultOptions
	opts.Dir = dir
	opts.ValueDir = valueDir
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	b := &Badger{
		Dir:      dir,
		ValueDir: valueDir,
	}
	return b
}

func (b *Badger) Set(key, value string) error {

	err = b.DB.Update(func(txn *badger.Txn) error {
		e := txn.Set([]byte(key), []byte(value))
		return e
	})
	return err
}

func (b *Badger) Get(key string) (string, error) {
	var value []byte
	err := b.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		value, err = item.Value()
		if err != nil {
			return err
		}
	})
	return string(value), err
}

func (b *Badger) Clear(key []byte) error {
	if err := b.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		return err
	}); err != nil {
		return err
	}
	return nil
}

// Implement raft.Storage

// InitialState returns the saved HardState and ConfState information.
func (b *Badger) InitialState() (pb.HardState, pb.ConfState, error) {

}

// Entries returns a slice of log entries in the range [lo,hi).
// MaxSize limits the total size of the log entries returned, but
// Entries returns at least one entry if any.
func (b *Badger) Entries(lo, hi, maxSize uint64) ([]pb.Entry, error) {

}

// Term returns the term of entry i, which must be in the range
// [FirstIndex()-1, LastIndex()]. The term of the entry before
// FirstIndex is retained for matching purposes even though the
// rest of that entry may not be available.
func (b *Badger) Term(i uint64) (uint64, error) {

}

// LastIndex returns the index of the last entry in the log.
func (b *Badger) LastIndex() (uint64, error) {

}

// FirstIndex returns the index of the first log entry that is
// possibly available via Entries (older entries have been incorporated
// into the latest Snapshot; if storage only contains the dummy entry the
// first log entry is not available).
func (b *Badger) FirstIndex() (uint64, error) {

}

// Snapshot returns the most recent snapshot.
// If snapshot is temporarily unavailable, it should return ErrSnapshotTemporarilyUnavailable,
// so raft state machine could know that Storage needs some time to prepare
// snapshot and call Snapshot later.
func (b *Badger) Snapshot() (pb.Snapshot, error) {

}
