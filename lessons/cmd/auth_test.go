package main

import (
	"bytes"
	"encoding/json"
	"go/adv-demo/internal/auth"
	"io"

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
	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Expected %d status code, got %d", 200, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body) 
	if err != nil {
		t.Fatal(err)
	}

	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
		if err != nil {
		t.Fatal(err)
	}

	if resData.Token == "" {
		t.Fatal("Token empty")
	}
}


func TestLoginFail(t *testing.T) {
	ts := httptest.NewServer(App())
	defer ts.Close()
	
	data, _ := json.Marshal(&auth.LoginRequest{
		Email: "mail@melu4.rur",
		Password: "wrong",
	})

	res, err := http.Post(ts.URL + "/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 401 {
		t.Fatalf("Expected %d status code, got %d", 401, res.StatusCode)
	}
}
