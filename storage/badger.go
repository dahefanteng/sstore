package badger

import (
	"log"

	"github.com/dgraph-io/badger"
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
