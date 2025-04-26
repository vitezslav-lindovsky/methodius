package methodius

import (
	"errors"
	"sync"
)

type KeyValueStore struct {
	store map[string]string
	mutex sync.RWMutex
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[string]string),
	}
}

func (kv *KeyValueStore) Get(key string) (value string, exists bool) {
	kv.mutex.RLock()
	defer kv.mutex.RUnlock()

	value, exists = kv.store[key]

	return
}

func (kv *KeyValueStore) GetAll() map[string]string {
	kv.mutex.RLock()
	defer kv.mutex.RUnlock()

	result := make(map[string]string)

	for k, v := range kv.store {
		result[k] = v
	}

	return result
}

func (kv *KeyValueStore) Set(key, value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	kv.mutex.Lock()
	defer kv.mutex.Unlock()

	if _, exists := kv.store[key]; exists {
		return errors.New("key already exists")
	}

	kv.store[key] = value
	return nil
}

func (kv *KeyValueStore) Update(key, value string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	kv.mutex.Lock()
	defer kv.mutex.Unlock()

	kv.store[key] = value
	return nil
}

func (kv *KeyValueStore) Delete(key string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	kv.mutex.Lock()
	defer kv.mutex.Unlock()

	if _, exists := kv.store[key]; !exists {
		return errors.New("key not found")
	}

	delete(kv.store, key)

	return nil
}
