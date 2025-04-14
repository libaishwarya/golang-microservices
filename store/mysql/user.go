package mysql

import (
	"database/sql"

	"github.com/libaishwarya/myapp/store"
)

type MySQLUserStore struct { // It is a struct that holds a database connection (*sql.DB).
	db *sql.DB
}

func NewUserStore(db *sql.DB) *MySQLUserStore { //creates a new MySQLUserStore with that connection.
	// Group all user-related DB functions (like create, get, update) in one place.
	// Keep the *sql.DB connection inside so it can use it for database work.
	return &MySQLUserStore{db: db}
}

func (s *MySQLUserStore) CreateUser(user *store.User) error {
	_, err := s.db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, user.Password)
	return err
}

func (s *MySQLUserStore) GetUserByEmail(email string) (*store.User, error) {
	row := s.db.QueryRow("SELECT id, email, password FROM users WHERE email = ?", email)
	user := &store.User{}
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *MySQLUserStore) StoreRes(fetchApi *store.ExternalUser) error {
	_, err := s.db.Exec("INSERT INTO fetch_users (id, name, email) VALUES (?, ?, ?)", fetchApi.ID, fetchApi.Name, fetchApi.Email)
	return err
}

func (s *MySQLUserStore) StoreCatFact(catFact *store.CatFact) error {
	_, err := s.db.Exec("INSERT INTO factCat (fact, length) VALUES (?, ?)", catFact.Fact, catFact.Length)
	return err
}
func NewDB() (*sql.DB, error) {
	return sql.Open("mysql", "myuser:mypassword@tcp(localhost:3306)/mydb")
}
