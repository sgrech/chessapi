package chessapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type chessApiService struct {
	baseUrl string
}

func InitChessApiService(baseUrl string) chessApiService {
	return chessApiService{baseUrl: baseUrl}
}

func (cpa chessApiService) GetPlayerDetails(username string) (playerDetails PlayerDetails, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/pub/player/%s", cpa.baseUrl, username))
	if err != nil {
		return PlayerDetails{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PlayerDetails{}, err
	}
	pd, err := unmarshalPlayerDetails(body)
	if err != nil {
		return PlayerDetails{}, err
	}
	return pd, nil
}
