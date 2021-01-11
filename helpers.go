package chessapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Converts unix timestamp string to time
func unixToTime(s string) (t time.Time, err error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	tm := time.Unix(i, 0)
	return tm, nil
}

// Unmarshals player details json into PlayerDetails struct
func unmarshalPlayerDetails(data []byte) (playerDetails PlayerDetails, err error) {
	var pd PlayerDetails
	if err = json.Unmarshal(data, &pd); err != nil {
		return pd, err
	}
	lastOnlineTime, err := unixToTime(fmt.Sprint(pd.LastOnline))
	joinedTime, err := unixToTime(fmt.Sprint(pd.Joined))
	pd.LastOnlineTime = lastOnlineTime
	pd.JoinedTime = joinedTime
	return pd, nil
}
