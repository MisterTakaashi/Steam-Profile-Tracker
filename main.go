package main

import (
	"colin-tracker/steam"
	"colin-tracker/storage"
	"log"
	"time"
)

func main() {
	dat, err := steam.GetRecentlyPlayedGames()

	if err != nil {
		log.Fatalln(err)
		return
	}

	store := storage.Store(&storage.JsonStore{})
	dbError := store.Init()
	defer store.Close()
	if dbError != nil {
		log.Fatalln(dbError)
		return
	}

	for _, game := range dat.Response.Games {
		entry := storage.StorageModel{
			Date:            time.Now(),
			Game:            game.Name,
			GameId:          game.Appid,
			Playtime2Weeks:  game.Playtime_2weeks,
			PlaytimeForever: game.Playtime_forever,
		}

		store.Write(&entry)
	}
}
