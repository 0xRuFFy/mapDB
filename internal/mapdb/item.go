package mapdb

import (
	"fmt"
)

type MapDBItem struct {
	Key   string
	Value string
}

func NewMapDBItem(key string, value interface{}) *MapDBItem {
	return &MapDBItem{
		Key:   key,
		Value: fmt.Sprintf("%v", value),
	}
}

func (i *MapDBItem) String() string {
	return fmt.Sprintf("%s: %v", i.Key, i.Value)
}
