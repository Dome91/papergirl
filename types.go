package papergirl

import (
	"errors"
	"io"
	_log "log"
	"reflect"
	"strconv"
)

var (
	log         Logger
	ErrNotFound = errors.New("not found")
)

type (
	ID   string
	Path string
)

type Logger interface {
	Info(msg string)
}

type Entity interface {
	ID() ID
}

type Repository[E Entity] interface {
	Save(entity E) error
	FindByID(id ID) (E, error)
	FindAll() ([]E, error)
	DeleteByID(id ID) error
	DeleteAll() error
}

type Storage interface {
	Store(path Path, reader io.Reader) error
	Retrieve(path Path, consumer func(io.Reader) error) error
}

type SimpleLogger struct {
}

func NewSimpleLogger() Logger {
	return &SimpleLogger{}
}

func (*SimpleLogger) Info(msg string) {
	_log.Println("INFO: " + msg)
}

type InMemoryRepository[E Entity] struct {
	id    int
	store map[ID]E
}

func NewInMemoryRepository[E Entity]() *InMemoryRepository[E] {
	store := make(map[ID]E)
	return &InMemoryRepository[E]{
		store: store,
	}
}

func (repository *InMemoryRepository[E]) Save(entity E) error {
	if entity.ID() == "" {
		id := strconv.Itoa(repository.id)
		repository.id = repository.id + 1
		ptr := reflect.ValueOf(&entity)
		v := reflect.Indirect(ptr)
		v.FieldByName("Id").SetString(id)
	}

	repository.store[entity.ID()] = entity
	return nil
}

func (repository *InMemoryRepository[E]) FindByID(id ID) (E, error) {
	entity := repository.store[id]
	return entity, nil
}

func (repository *InMemoryRepository[E]) FindAll() ([]E, error) {
	var entities []E
	for _, entity := range repository.store {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (repository *InMemoryRepository[E]) DeleteByID(id ID) error {
	delete(repository.store, id)
	return nil
}

func (repository *InMemoryRepository[E]) DeleteAll() error {
	for id := range repository.store {
		delete(repository.store, id)
	}

	return nil
}

type InMemoryStorage struct {
	storage map[Path][]byte
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		storage: make(map[Path][]byte),
	}
}

func (storage *InMemoryStorage) Store(path Path, reader io.Reader) error {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	storage.storage[path] = bytes
	return nil
}
