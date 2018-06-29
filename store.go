package main

type Storer interface {
	Set(string, interface{})
	Get(string) interface{}
	GetString(string) string
	GetInt(string) int
	Update(string, func(interface{}) interface{})
}

type Store struct {
	data map[string]interface{}
}

func (s *Store) Set(key string, value interface{}) {
	s.data[key] = value
	globalObserver.Notify(key)
}

func (s *Store) Get(key string) interface{} {
	globalObserver.Register(key)
	return s.data[key]
}

func (s *Store) GetString(key string) string {
	return s.Get(key).(string)
}

func (s *Store) GetInt(key string) int {
	return s.Get(key).(int)
}

func (s *Store) Update(key string, fn func(interface{}) interface{}) {
	s.data[key] = fn(s.data[key])
	globalObserver.Notify(key)
}

func NewStore() Storer {
	return &Store{data: make(map[string]interface{}, 0)}
}
