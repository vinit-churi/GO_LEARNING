//go:build solution
// +build solution

package main

import "fmt"

type Reader interface {
	Load(id string) (string, bool)
}

type Writer interface {
	Save(id, value string)
}

type ReaderStore interface {
	Reader
	Writer
}

type MemoryStore struct {
	values map[string]string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{values: map[string]string{}}
}

func (m *MemoryStore) Load(id string) (string, bool) {
	value, ok := m.values[id]
	return value, ok
}

func (m *MemoryStore) Save(id, value string) {
	m.values[id] = value
}

func SaveReader(store ReaderStore, id string) error {
	store.Save(id, "reader:"+id)
	return nil
}

func UpsertLabel(store ReaderStore, id string, value string) string {
	if existing, ok := store.Load(id); ok {
		return existing
	}
	store.Save(id, value)
	return value
}

func CopyIfMissing(store ReaderStore, src, dst string) error {
	if _, ok := store.Load(dst); ok {
		return nil
	}
	value, ok := store.Load(src)
	if !ok {
		return fmt.Errorf("missing source: %s", src)
	}
	store.Save(dst, value)
	return nil
}

func main() {
	store := NewMemoryStore()
	_ = SaveReader(store, "alpha")
	fmt.Println(UpsertLabel(store, "alpha", "override"))
}
