package storage

import (
	"rockpaperscissors/challenge"
	"rockpaperscissors/user"
	"sync"
)

type MemoryStore struct {
	users       map[int]*user.User
	challenges  map[int]*challenge.Challenge
	userID      int
	challengeID int
	mu          sync.Mutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users:       make(map[int]*user.User),
		challenges:  make(map[int]*challenge.Challenge),
		userID:      1,
		challengeID: 1,
	}
}

func (m *MemoryStore) AddUser(u *user.User) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	u.ID = m.userID
	m.users[m.userID] = u
	m.userID++
	return u.ID
}

func (m *MemoryStore) FindUserByUsername(username string) (*user.User, bool) {
	for _, u := range m.users {
		if u.Username == username {
			return u, true
		}
	}
	return nil, false
}

func (m *MemoryStore) FindUserByID(id int) (*user.User, bool) {
	u, ok := m.users[id]
	return u, ok
}

func (m *MemoryStore) ListUsers() []*user.User {
	var list []*user.User
	for _, u := range m.users {
		list = append(list, u)
	}
	return list
}

func (m *MemoryStore) ListAIUsers() []*user.User {
	var list []*user.User
	for _, u := range m.users {
		if u.IsAI {
			list = append(list, u)
		}
	}
	return list
}

func (m *MemoryStore) UpdateUser(u *user.User) error {
	m.users[u.ID] = u
	return nil
}

func (m *MemoryStore) AddChallenge(c *challenge.Challenge) int {
	m.mu.Lock()
	defer m.mu.Unlock()
	c.ID = m.challengeID
	m.challenges[m.challengeID] = c
	m.challengeID++
	return c.ID
}

func (m *MemoryStore) FindChallengeByID(id int) (*challenge.Challenge, bool) {
	c, ok := m.challenges[id]
	return c, ok
}

func (m *MemoryStore) ListChallenges() []*challenge.Challenge {
	var list []*challenge.Challenge
	for _, c := range m.challenges {
		list = append(list, c)
	}
	return list
}

func (m *MemoryStore) UpdateChallenge(c *challenge.Challenge) error {
	m.challenges[c.ID] = c
	return nil
}
