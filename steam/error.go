package steam

type SteamError interface{}

type SteamApiError struct {
	Err error
}

type SteamResponseError struct {
	Err error
}
