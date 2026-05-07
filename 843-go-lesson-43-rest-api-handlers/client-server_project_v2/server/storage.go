// Хранилище + загрузка/сохранение в файл.
package server

import (
	"encoding/json"
	"os"
	"sync"
)

type UserStorage struct {
	mu     sync.Mutex
	users  []User
	nextID int
	file   string
}

func NewUserStorage(file string) *UserStorage {
	s := &UserStorage{
		users:  []User{},
		nextID: 1,
		file:   file,
	}
	s.load()
	return s
}

func (s *UserStorage) load() {
	data, err := os.ReadFile(s.file)
	if err != nil {
		return
	}

	json.Unmarshal(data, &s.users)

	maxID := 0
	for _, u := range s.users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	s.nextID = maxID + 1
}

func (s *UserStorage) save() {
	data, _ := json.MarshalIndent(s.users, "", "  ")
	os.WriteFile(s.file, data, 0644)
}

func (s *UserStorage) GetAll() []User {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]User{}, s.users...)
}

func (s *UserStorage) Create(name string) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := User{
		ID:   s.nextID,
		Name: name,
	}
	s.nextID++
	s.users = append(s.users, user)
	s.save()
	return user
}

func (s *UserStorage) Update(id int, name string) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.users {
		if s.users[i].ID == id {
			s.users[i].Name = name
			s.save()
			return s.users[i], true
		}
	}
	return User{}, false
}

func (s *UserStorage) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.users {
		if s.users[i].ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			s.save()
			return true
		}
	}
	return false
}
