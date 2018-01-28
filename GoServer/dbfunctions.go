package main

import (
	"database/sql"
	"encoding/base64"
	"log"
)

func dbLogin(user Luser) string {
	var (
		name string
	)
	//Wrappad i en funktion?
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT username FROM userinfo WHERE username=$1 AND password=$2;", user.Username, user.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	return name
}

func dbCheckUsername(user Ruser) int {
	var username string
	//Wrappad i en funktion?
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT username FROM userinfo WHERE username=$1;", user.Username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	if username != "" {
		return 1
	}
	return 0
}

func dbRegister(user Ruser) error {
	//Wrappad i en funktion?
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		return error(err)
	}

	rows, err := db.Query("INSERT INTO userinfo(username, password, salt) VALUES($1, $2, $3);", user.Username, user.Password, user.Salt)
	if err != nil {
		return error(err)
	}
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		return error(err)
	}
	db.Close()
	return nil
}

func dbGetSalt(user Luser) []byte {
	var (
		salt string
	)
	//Wrappad i en funktion?
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT salt FROM userinfo WHERE username=$1;", user.Username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&salt)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()

	sal, _ := base64.StdEncoding.DecodeString(salt)
	return []byte(sal)
}
