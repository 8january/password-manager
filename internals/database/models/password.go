package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/8january/password-manager/internals/crypto"
)

type Password struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Service   string    `db:"service"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

func Save(db *sql.DB, name, service, password, passphrase string) {
	query := `INSERT INTO passwords (name, service, password, created_at) values (?, ?, ?, ?)`
	_, err := db.Exec(
		query,
		name,
		service,
		string(crypto.Encrypt([]byte(password), passphrase)),
		time.Now(),
	)

	if err != nil {
		log.Fatal(err)
	}

}

func Get(db *sql.DB, id int, passphrase string) string {
	query := `SELECT password FROM passwords WHERE id = ?`
	row := db.QueryRow(query, id)

	var p string
	err := row.Scan(&p)
	if err != nil {
		log.Fatal(err)
	}

	pwd := crypto.Decrypt([]byte(p), passphrase)
	return string(pwd)
}

func List(db *sql.DB) []Password {
	query := `SELECT * FROM passwords`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var passwords []Password
	for rows.Next() {
		p := &Password{}
		err := rows.Scan(&p.ID,
			&p.Name,
			&p.Service,
			&p.Password,
			&p.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}

		passwords = append(passwords, *p)
	}

	return passwords
}

func Delete(db *sql.DB, id int, passphrase string) {
	queryi := `SELECT * FROM passwords WHERE id = ?`
	row := db.QueryRow(queryi, id)
	p := &Password{}
	err := row.Scan(&p.ID,
		&p.Name,
		&p.Service,
		&p.Password,
		&p.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	_ = crypto.Decrypt([]byte(p.Password), passphrase)

	queryii := `DELETE FROM passwords WHERE id = ?`
	result, err := db.Exec(queryii, id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SENHA EXCLUIDA!")

}

func Update(db *sql.DB, id int, newPassword, oldPassphrase, newPassphrase string) {
	queryi := `SELECT password FROM passwords WHERE id = ?`
	row := db.QueryRow(queryi, id)

	var p string
	err := row.Scan(&p)
	if err != nil {
		log.Fatal(err)
	}

	_ = crypto.Decrypt([]byte(p), oldPassphrase)
	np := string(crypto.Encrypt([]byte(newPassword), newPassphrase))
	fmt.Println("NEW PWD: ", np)

	queryii := `UPDATE passwords SET password = ? WHERE id = ?`
	result, err := db.Exec(queryii, np, id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

}
