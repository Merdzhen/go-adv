package main

import (
	"bytes"
	"encoding/json"
	"go/adv-demo/internal/auth"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()
	
	data, _ := json.Marshal(&auth.LoginRequest{
		Email: "mail@melu4.rur",
		Password: "123",
	})

	res, err := http.Post(ts.URL + "/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("Expected %d status code, got %d", 200, res.StatusCode)
	}
}
