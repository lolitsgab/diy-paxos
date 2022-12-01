package storage

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Storage interface {
	Insert(k string, v int32) error
	Get(k string) (int32, error)
	Remove(k string) error
	Update(k string, v int32) error
	Upsert(k string, v int32) error
}

type InMemoryStorage struct {
	store map[string]int32
}

func NewInMemoryStorage() Storage {
	return &InMemoryStorage{map[string]int32{}}
}

func (i InMemoryStorage) Insert(k string, v int32) error {
	if _, exists := i.store[k]; exists {
		return status.New(codes.AlreadyExists, "cannot insert to an already populated key").Err()
	}
	i.store[k] = v
	return nil
}

func (i InMemoryStorage) Get(k string) (int32, error) {
	if _, exists := i.store[k]; !exists {
		return -1, status.New(codes.NotFound, "cannot get non existent key").Err()
	}
	return i.store[k], nil
}

func (i InMemoryStorage) Remove(k string) error {
	if _, exists := i.store[k]; !exists {
		return status.New(codes.NotFound, "cannot remove non existent key").Err()
	}
	delete(i.store, k)
	return nil
}

func (i InMemoryStorage) Update(k string, v int32) error {
	if _, exists := i.store[k]; !exists {
		return status.New(codes.NotFound, "cannot update a non existent key").Err()
	}
	i.store[k] = v
	return nil
}

func (i InMemoryStorage) Upsert(k string, v int32) error {
	if _, exists := i.store[k]; !exists {
		return i.Insert(k, v)
	}
	return i.Update(k, v)
}
