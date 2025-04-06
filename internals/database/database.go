package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/8january/password-manager/internals/database/migration"
	"github.com/8january/password-manager/internals/database/models"
)

type Database struct {
	conn *sql.DB
}

func Init(path string) *Database {
	db, err := sql.Open("sqlite3", path+".db")
	if err != nil {
		log.Fatal(err)
	}

	return &Database{conn: db}
}

func (db *Database) Migrate() error {
	return migration.RunMigrations(db.conn)
}

func (db *Database) Save(name, service, password, passphrase string) {
	models.Save(db.conn, name, service, password, passphrase)
}

func (db *Database) Get(id int, passphrase string) string {
	return models.Get(db.conn, id, passphrase)
}

func (db *Database) List() []models.Password {
	return models.List(db.conn)
}

func (db *Database) Delete(id int, passphrase string) {
	models.Delete(db.conn, id, passphrase)
}

func (db *Database) Update(id int, newPassword, oldPassphrase, newPassphrase string) {
	models.Update(db.conn, id, newPassword, oldPassphrase, newPassphrase)
}
