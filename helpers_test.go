package chessapi

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func loadTestData(t *testing.T, path string) []byte {
	if bs, err := ioutil.ReadFile(path); err == nil {
		return bs
	} else {
		t.Fatalf("Error reading test data from %v with error %v", path, err)
		return nil
	}
}

func TestUnixToTime(t *testing.T) {
	loc := time.FixedZone("UTC-8", 1*60*60)
	tm_test := time.Date(2021, time.January, 8, 11, 25, 17, 0, loc)
	if tm, err := unixToTime(fmt.Sprint(1610101517)); err == nil {
		if !tm.Equal(tm_test) {
			t.Fatalf("Bad time conversion, expected %s but got %s", tm_test.Format(time.RFC3339), tm.Format(time.RFC3339))
		}
	} else {
		t.Fatalf("Error parsing time with error %v", err)
	}
}

func TestUnmarshalPlayerDetails(t *testing.T) {
	data := loadTestData(t, "./test/data/player_details.json")
	lastOnlineTime, _ := unixToTime("1610101517")
	joinedTime, _ := unixToTime("1591477924")
	p_test := PlayerDetails{
		Avatar:         "https://images.chesscomfiles.com/uploads/v1/user/81701324.219d90b6.200x200o.81c4f77b54f6.jpeg",
		PlayerID:       81701324,
		ID:             "https://api.chess.com/pub/player/eladhel",
		Url:            "https://www.chess.com/member/eladhel",
		Name:           "Shane Grech",
		Username:       "eladhel",
		Followers:      9,
		Country:        "https://api.chess.com/pub/country/MT",
		Location:       "Malta",
		LastOnline:     1610101517,
		Joined:         1591477924,
		LastOnlineTime: lastOnlineTime,
		JoinedTime:     joinedTime,
		Status:         "premium",
		IsStreamer:     false,
	}
	if p, err := unmarshalPlayerDetails(data); err == nil {
		if p != p_test {
			t.Fatalf("Bad data, expected got %v+", p)
		}
	} else {
		t.Fatalf("Error unmarshaling test data from with error %v", err)
	}
}
