package db

import "go-http-boilerplate/src/domain/users/model"

// UserStore just public for type
type UserStore struct {
	users map[string]model.User
}

// newUserStore private (capsulation)
func newUserStore() *UserStore {
	return &UserStore{
		users: map[string]model.User{},
	}
}

func (us *UserStore) Save(u model.User) {
	us.users[u.Username] = u
}

func (us *UserStore) Exist(u model.User) bool {
	_, has := us.users[u.Username]
	return has
}

func (us *UserStore) Find(u model.User) (*model.User, bool) {
	dbu, has := us.users[u.Username]
	return &dbu, has
}
