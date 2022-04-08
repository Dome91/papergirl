package papergirl

import "io"

type ID string
type Path string

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
}

type InMemoryRepository[E Entity] struct {
	store map[ID]E
}

func NewInMemoryRepository[E Entity]() *InMemoryRepository[E] {
	store := make(map[ID]E)
	return &InMemoryRepository[E]{
		store: store,
	}
}

func (repository *InMemoryRepository[E]) Save(entity E) error {
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
