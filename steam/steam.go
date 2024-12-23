package steam

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var steamUrl = "https://api.steampowered.com/IPlayerService/GetRecentlyPlayedGames/v1"

type SteamGetRecentlyPlayedGamesResponse struct {
	Response struct {
		Total_count uint8
		Games       []struct {
			Appid            uint
			Name             string
			Playtime_2weeks  uint
			Playtime_forever uint
			Img_icon_url     string
		}
	}
}

func GetRecentlyPlayedGames() (SteamGetRecentlyPlayedGamesResponse, SteamError) {
	v := url.Values{}
	v.Add("key", os.Getenv("STEAM_API_KEY"))
	v.Add("steamid", os.Getenv("STEAM_ID"))

	response, error := http.Get(fmt.Sprint(steamUrl, "?", v.Encode()))

	if error != nil {
		return SteamGetRecentlyPlayedGamesResponse{}, SteamApiError{
			Err: error,
		}
	}

	if response.StatusCode > 299 {
		errBody, _ := io.ReadAll(response.Body)

		return SteamGetRecentlyPlayedGamesResponse{}, SteamApiError{
			Err: errors.New(fmt.Sprintf("Bad result from API: %s", errBody)),
		}
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return SteamGetRecentlyPlayedGamesResponse{}, SteamResponseError{
			Err: err,
		}
	}

	var dat SteamGetRecentlyPlayedGamesResponse

	if err := json.Unmarshal(body, &dat); err != nil {
		if err != nil {
			return SteamGetRecentlyPlayedGamesResponse{}, SteamResponseError{
				Err: err,
			}
		}
	}

	return dat, nil
}
