package chessapi

import (
	"testing"
)

func TestChessApiService(t *testing.T) {
	cpa := InitChessApiService("https://api.chess.com")
	if cpa.baseUrl != "https://api.chess.com" {
		t.Fatalf("baseUrl property not set correctly, got %s", cpa.baseUrl)
	}
}
