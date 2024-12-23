package steam

type SteamApiError struct {
	Err error
}

func (e *SteamApiError) Error() string {
	return e.Error()
}

type SteamResponseError struct {
	Err error
}

func (e *SteamResponseError) Error() string {
	return e.Error()
}
