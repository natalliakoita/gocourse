package datasource

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5438
	user     = "postgres"
	password = "qwerty"
	dbname   = "support"
)

// create new connection in database
func NewDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	return conn
}

type DS struct {
	DB *sql.DB
}

func NewDS(conn *sql.DB) DS {
	return DS{
		DB: conn,
	}
}

// add contacts in database
func (ds *DS) AddContact(name, number string) error {
	insertStatement := `INSERT INTO contacts (name, number) VALUES ($1, $2)`
	_, err := ds.DB.Exec(insertStatement, name, number)

	if err != nil {
		return err
	}
	return nil
}

// assign contact to group
func (ds *DS) AddGroupToContact(name, number string, group_id int) error {
	updateStatement := `UPDATE contacts SET group_id = $1 WHERE name = $2 AND number = $3`
	result, err := ds.DB.Exec(updateStatement, group_id, name, number)
	if err != nil {
		return err
	}
	z, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if z == 0 {
		return errors.New("Not found contact")
	}
	return nil
}

// get contact by group 
func (ds *DS) GetContactsOrderByGroup(group_id int) ([]ContactGroup, error) {
	getStatement := `SELECT c.id, c."name", c."number", COALESCE(gr."name", '') FROM  contacts c LEFT JOIN "groups" gr ON c.group_id = gr.id GROUP BY c.id , c.number, c."name", gr."name" ORDER BY gr."name" ASC`
	rows, err := ds.DB.Query(getStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	contacts := []ContactGroup{}
	for rows.Next() {
		p := ContactGroup{}
		err := rows.Scan(&p.Contact.ID, &p.Contact.Name, &p.Contact.Number, &p.Group.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		contacts = append(contacts, p)

	}
	return contacts, nil
}
