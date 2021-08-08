package datasource

type Contact struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Number string `db:"number"`
	Group  string `db:"group_id"`
}

type Group struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
}

type ContactGroup struct {
	Contact Contact
	Group Group
}
