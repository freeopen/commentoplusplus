package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func dbConnect(retriesLeft int) error {
	con := os.Getenv("DB")

	fmt.Println("con ===> " + con)
	u, err := url.Parse(con)
	// fmt.Println("scheme ===> " + u.Scheme)
	if u.Scheme == "postgres" {
		u.User = url.UserPassword(u.User.Username(), "redacted")
		logger.Infof("opening connection to db: %s", u.String())

		db, err = sql.Open("postgres", con)
		if err != nil {
			logger.Errorf("cannot open connection to postgres: %v", err)
			return err
		}
		err = db.Ping()
		if err != nil {
			if retriesLeft > 0 {
				logger.Errorf("cannot talk to postgres, retrying in 10 seconds (%d attempts left): %v", retriesLeft-1, err)
				time.Sleep(10 * time.Second)
				return dbConnect(retriesLeft - 1)
			} else {
				logger.Errorf("cannot talk to postgres, last attempt failed: %v", err)
				return err
			}
		}
		maxIdleConnections, err := strconv.Atoi(os.Getenv("MAX_IDLE_PG_CONNECTIONS"))
		if err != nil {
			logger.Warningf("cannot parse COMMENTO_MAX_IDLE_PG_CONNECTIONS: %v", err)
			maxIdleConnections = 50
		}

		db.SetMaxIdleConns(maxIdleConnections)

	} else if u.Scheme == "sqlite3" {

		// fmt.Println("path ===> " + u.Path)
		db, err = sql.Open("sqlite3", u.Path)
		if err != nil {
			logger.Errorf("cannot open connection to sqlite3: %v", err)
			return err
		}
	} else {
		logger.Errorf("Unsupported database types")
	}

	if err != nil {
		logger.Errorf("invalid db connection URI: %v", err)
		return err
	}

	statement := `
		CREATE TABLE IF NOT EXISTS migrations (
			filename TEXT NOT NULL UNIQUE
		);
	`
	_, err = db.Exec(statement)
	if err != nil {
		logger.Errorf("cannot create migrations table: %v", err)
		return err
	}

	return nil
}
