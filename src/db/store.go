package db

// Singleton Design
func init() {
	db = newStore()
}

var db *store

type store struct {
	us *UserStore
}

// newStore private (capsulation)
func newStore() *store {
	return &store{
		us: newUserStore(),
	}
}

// GetUserStore get private UserStore instance
func GetUserStore() *UserStore {
	return db.us
}
