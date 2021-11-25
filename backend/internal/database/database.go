package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"web/config"
	"web/internal/model"

	_ "github.com/lib/pq" //Database driver
)

// DbConnection stores connection to database
type DbConnection struct {
	conn *sql.DB
}

// Connect to DB
func (db *DbConnection) Connect(cfg *config.Config) error {
	var err error
	db.conn, err = sql.Open("postgres", cfg.PostgresURL)
	if err != nil {
		db.conn.Close()
		return err
	}
	return nil
}

// Close connection with DB
func (db *DbConnection) Close() {
	db.conn.Close()
}

// AddUsrInfo to db
func (db *DbConnection) AddUsrInfo(user model.User) {
	db.conn.Exec("insert into users(name, email, password, id_role) values($1, $2, $3, $4)",
		user.Name, user.Email, user.Password, user.IDRole,
	)

}

// DeleteUser from db
func (db *DbConnection) DeleteUser(user model.User) error {
	_, err := db.conn.Exec("delete from comment where id_user = (select id_user from users where email = $1 limit 1);",
		user.Email,
	)
	if err != nil {
		return err
	}
	_, err = db.conn.Exec("delete from users where id_user = (select id_user from users where email = $1 limit 1);",
		user.Email,
	)
	if err != nil {
		return err
	}
	return nil
}

// MakeMod from user
func (db *DbConnection) MakeMod(user model.User) error {
	_, err := db.conn.Exec("update users set id_role = 1 where email = $1;",
		user.Email,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetUsers from Db
func (db *DbConnection) GetUsers() ([]model.User, error) {
	rows, err := db.conn.Query(`select array_to_json(array_agg(lap))
        from (select t1.name as "username", t1.email, t2.name_role AS "role" 
		from users t1 join roles t2 on t1.id_role = t2.id_role) lap;`)
	if err != nil {
		return nil, err
	}

	var arrStr string
	var users []model.User

	for rows.Next() {
		rows.Scan(&arrStr)
		json.Unmarshal([]byte(arrStr), &users)

	}

	return users, nil
}

// GetUserByEmail from Db
func (db *DbConnection) GetUserByEmail(user model.User) (model.User, error) {
	row := db.conn.QueryRow(`select name, t1.id_user, t1.email, t2.name_role from users t1 join roles t2
		on t1.id_role = t2.id_role where email = $1 and password = $2 limit 1;`, user.Email, user.Password)

	var resUser model.User
	switch err := row.Scan(&resUser.Name, &resUser.ID, &resUser.Email, &resUser.Role); err {
	case sql.ErrNoRows:
		resUser.Email = "Unidentified"
		return resUser, nil
	case nil:
		return resUser, nil
	default:
		resUser.Email = "Unidentified"
		return resUser, err
	}

}

// GetRoadsAndCourorts from DB
func (db *DbConnection) GetRoadsAndCourorts() []model.Road {
	rows, err := db.conn.Query(`select array_to_json(array_agg(lap))
		from (select t1.name_road, t2.name_courort from roads t1 join courorts t2 on 
		t1.id_courort = t2.id_courort) lap;`)
	if err != nil {
		return nil
	}

	var arrStr string
	var roads []model.Road

	for rows.Next() {
		rows.Scan(&arrStr)
		json.Unmarshal([]byte(arrStr), &roads)
	}
	return roads
}

// GetCourorts - get list of courorts
func (db *DbConnection) GetCourorts() ([]model.Courorts, error) {
	rows, err := db.conn.Query(`select array_to_json(array_agg(lap))
		from (select name_courort, city from courorts) lap;`)
	if err != nil {
		return nil, nil
	}

	var arrStr string
	var cour []model.Courorts

	for rows.Next() {
		rows.Scan(&arrStr)
		json.Unmarshal([]byte(arrStr), &cour)
	}
	return cour, nil
}

// GetCourort to service
func (db *DbConnection) GetCourort(courort string) ([]model.Courort, error) {
	var res []model.Courort
	var id int
	switch courort {
	case "rosa":
		id = 0
	case "laura":
		id = 1
	case "gorod":
		id = 2
	}
	rows, err := db.conn.Query(
		` select array_to_json(array_agg(lap))
		from (select type_road, name_road, work_status from roads where id_courort = $1) lap;`, id)
	if err != nil {
		return res, err
	}
	var arrStr string

	for rows.Next() {
		rows.Scan(&arrStr)
		json.Unmarshal([]byte(arrStr), &res)
	}
	return res, nil
}

// AddComment to service
func (db *DbConnection) AddComment(com model.Comment) error {
	row := db.conn.QueryRow(`
	select id_user from users where email = $1
	`, com.Email)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	fmt.Println(id)
	_, err = db.conn.Query(`insert into comment (id_user, id_courort, content, likes, visability) values ($1, $2, $3, $4, $5);`,
		id, com.IDCourort, com.Text, 0, 0,
	)
	return nil
}

// DeleteComment From service
func (db *DbConnection) DeleteComment(com model.Comment) error {
	_, err := db.conn.Query(`delete from comment
	where id_comment =  $1;`, com.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetComments From service
func (db *DbConnection) GetComments(com model.Comment) ([]model.Comment, error) {
	var res []model.Comment
	rows, err := db.conn.Query(
		`select array_to_json(array_agg(lap))
        from
		(select t1.content as text, t2.email as email from comment t1
			 join users t2 on t1.id_user = t2.id_user 
			 where id_courort = $1) lap;`, com.IDCourort,
	)

	if err != nil {
		return res, err
	}
	var arrStr string

	for rows.Next() {
		rows.Scan(&arrStr)
		json.Unmarshal([]byte(arrStr), &res)
	}
	return res, nil
}
