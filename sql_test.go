package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id,nama) VALUES(2,'rona theresa idang ngo')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestExecQuery(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, nama FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}
	defer rows.Close()
}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, nama, email, balance, rating, birth_date,married,created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id :", id, "Name :", name, "Email :", email, "Balance :", balance, "Rating :", rating, "Birth Date :", birthDate, "Married :", married, "Created At :", createdAt)
	}
	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin'; #"
	password := "salah"
	query := "SELECT username FROM users where username = '" + username + "' AND password =  '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)
	fmt.Println(query)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}

	defer rows.Close()
}

func TestHandleSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin'; #"
	password := "salah"
	query := "SELECT username FROM users where username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	fmt.Println(query)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}

	defer rows.Close()
}

func TestAutoIncrementLastInsertId(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	name := "meja belajar"
	query := "INSERT INTO category(name) values (?)"
	result, err := db.ExecContext(ctx, query, name)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(lastInsertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO category(name) values (?)"
	stmt, err := db.PrepareContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	for i := 0; i < 100; i++ {
		name := "kategori ke " + strconv.Itoa(i)
		result, err := stmt.ExecContext(ctx, name)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Last insert Id = ", id)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}
	query := "INSERT INTO category(name) values (?)"

	for i := 0; i < 100; i++ {
		name := "kategori ke " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, query, name)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Last insert Id = ", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

}
