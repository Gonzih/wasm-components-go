package main

type Storer interface {
	DefMutation(string, func(...interface{}) error) error
	Mutate(string, ...interface{}) error
}

type Store struct {
}

func NewStore() *Store {
	return &Store{}
}
