package storage

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Metadata struct {
	Version int32
}

type Value struct {
	Value int32
	Meta  Metadata
}

// Storage exposes an API to interface with a KV Storage container.
type Storage interface {
	Insert(k string, v int32) error
	Get(k string) (int32, error)
	Remove(k string) error
	Update(k string, v int32) error
	Upsert(k string, v int32) error
}

// InMemoryStorage implements Storage with an in-memory table.
type InMemoryStorage struct {
	store map[string]Value
}

// NewInMemoryStorage creates a new InMemoryStorage instance.
func NewInMemoryStorage() Storage {
	return &InMemoryStorage{map[string]Value{}}
}

// Insert inserts a KV pair into the Storage container.
func (i *InMemoryStorage) Insert(k string, v int32) error {
	if _, exists := i.store[k]; exists {
		return status.New(codes.AlreadyExists, "cannot insert to an already populated key").Err()
	}
	i.store[k] = Value{Value: v}
	return nil
}

// Get fetches value by key from the Storage container.
func (i *InMemoryStorage) Get(k string) (int32, error) {
	if _, exists := i.store[k]; !exists {
		return -1, status.New(codes.NotFound, "cannot get non existent key").Err()
	}
	return i.store[k].Value, nil
}

// Remove deletes a value by key from a storage container.
func (i *InMemoryStorage) Remove(k string) error {
	if _, exists := i.store[k]; !exists {
		return status.New(codes.NotFound, "cannot remove non existent key").Err()
	}
	delete(i.store, k)
	return nil
}

// Update changes the value associated with a key entry in the Storage container iff it exists.
func (i *InMemoryStorage) Update(k string, v int32) error {
	if _, exists := i.store[k]; !exists {
		return status.New(codes.NotFound, "cannot update a non existent key").Err()
	}
	i.store[k] = Value{Value: v}
	return nil
}

// Upsert changes the value associated with a key entry in the Storage container, or creates it if it does not exist.
func (i *InMemoryStorage) Upsert(k string, v int32) error {
	if _, exists := i.store[k]; !exists {
		return i.Insert(k, v)
	}
	return i.Update(k, v)
}
