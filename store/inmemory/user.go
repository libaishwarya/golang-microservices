package inmemory

import (
	"errors"

	"github.com/libaishwarya/myapp/store"
	"github.com/libaishwarya/myapp/userservice"
)

type InMemoryUserStore struct {
	users map[string]*store.User
	store []store.ExternalUser
}

type InMemoryStore struct {
	catFacts []userservice.CatFact
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*store.User),
	}
}
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		catFacts: []userservice.CatFact{},
	}
}
func (s *InMemoryUserStore) CreateUser(user *store.User) error {
	if _, exists := s.users[user.Email]; exists {
		return errors.New("user already exists")
	}
	s.users[user.Email] = user
	return nil
}

func (s *InMemoryUserStore) GetUserByEmail(email string) (*store.User, error) {
	user, exists := s.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *InMemoryUserStore) StoreRes(ExternalApi *store.ExternalUser) error {
	s.store = append(s.store, *ExternalApi)
	return nil
}
func (s *InMemoryStore) StoreCatFact(fact *userservice.CatFact) error {
	s.catFacts = append(s.catFacts, *fact)
	return nil
}
