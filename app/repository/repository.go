package repository

// Object : interface for all objects
type Object interface{}

// BaseRepositoryInterface : test
type BaseRepositoryInterface interface {
	first()
}

// BaseRepository : test
type BaseRepository struct {
	model   Object
	objects []Object
}

func (r *BaseRepository) first() Object {
	// TODO : Query from db

	return r.model
}
