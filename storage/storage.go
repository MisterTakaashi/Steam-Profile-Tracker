package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const fileName = "./data.db"

type Store interface {
	Init() error
	Close()
	Write(entry *StorageModel) error
}

type JsonStore struct {
	db *sql.DB
}

func (store *JsonStore) Init() error {
	db, err := sql.Open("sqlite3", fileName)

	if err != nil {
		return err
	}
	store.db = db

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS entry (date datetime, game string, game_id int, playtime_2weeks int, playtime_forever int)")
	if err != nil {
		return err
	}

	return nil
}

func (store *JsonStore) Close() {
	store.db.Close()
}

func (store *JsonStore) Write(entry *StorageModel) error {
	log.Println(fmt.Sprintf("Storing stats for game %s", entry.Game))

	query, err := store.db.Prepare("INSERT INTO entry(date, game, game_id, playtime_2weeks, playtime_forever) VALUES(?, ?, ?, ?, ?)")
	defer query.Close()
	if err != nil {
		return err
	}

	_, err = query.Exec(entry.Date, entry.Game, entry.GameId, entry.Playtime2Weeks, entry.PlaytimeForever)
	if err != nil {
		return err
	}

	return nil
}
