package domain

import (
	"fmt"
	"sync"
)

type MemUserRepo struct {
	data   map[int64]*User
	users  map[string]*User
	nextID int64
	mu     sync.RWMutex
}

func NewMemUserRepo() *MemUserRepo {
	return &MemUserRepo{
		data:   make(map[int64]*User),
		users:  make(map[string]*User),
		nextID: 1,
	}
}

func (m *MemUserRepo) AddUser(user *User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if user.UserID == 0 {
		user.UserID = m.nextID
		m.nextID += 1
	}

	if m.data[user.UserID] != nil {
		return fmt.Errorf("user with id %d already existed", user.UserID)
	}

	if m.users[user.Username] != nil {
		return fmt.Errorf("username %s already used", user.Username)
	}

	m.users[user.Username] = user
	m.data[user.UserID] = user
	return nil
}

func (m *MemUserRepo) UserByName(username string) (*User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.users[username], nil
}

func (m *MemUserRepo) UserByID(id int64) (*User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.data[id], nil
}

func (m *MemUserRepo) AllUsers() (map[int64]*User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.data, nil
}
