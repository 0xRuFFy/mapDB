package mapdb

import (
	"sync"

	"github.com/0xRuFFy/mapDB/internal/utils/errors"
)

type Database struct {
	Items map[string]*MapDBItem
	mu    sync.RWMutex
}

func NewDatabase() *Database {
	return &Database{
		Items: make(map[string]*MapDBItem),
	}
}

func (db *Database) Get(key string) (*MapDBItem, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	item, ok := db.Items[key]
	if !ok {
		return nil, errors.KeyNotFoundError()
	}

	return item, nil
}

func (db *Database) Set(key string, value string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	item := NewMapDBItem(key, value)

	db.Items[key] = item

	return nil
}

func (db *Database) Delete(key string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.Items, key)

	return nil
}

func (db *Database) Keys() []string {
	db.mu.RLock()
	defer db.mu.RUnlock()

	keys := make([]string, 0, len(db.Items))
	for key := range db.Items {
		keys = append(keys, key)
	}

	return keys
}
