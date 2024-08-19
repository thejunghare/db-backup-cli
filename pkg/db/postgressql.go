package db

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"time"

	_ "github.com/lib/pq"
)

func Connect(user, password, dbname string) (*sql.DB, error) {
	const (
		host = "localhost"
		port = 5432
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connection successfull!")
	return db, nil
}

func Backup(dbname, user, password string) (string, error) {
	backupfile := fmt.Sprintf("%s_backup%s.sql", dbname, time.Now().Format("20060102_150405"))

	os.Setenv("PGPASSWORD", password)

	cmd := exec.Command("pg_dump", "-U", user, "-F", "a", "-f", backupfile, dbname)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to perform backup: %w", err)
	}

	return backupfile, nil
}
