package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPStatusHandler(t *testing.T) {
	t.Run("it should return httpCode 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/status", nil)
		if err != nil {
			t.Error(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(StatusHandler)
		handler.ServeHTTP(resp, req)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("it should return status = OK", func (t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/status", nil)
		if err != nil {
			t.Error(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(StatusHandler)
		handler.ServeHTTP(resp, req)
		var expectedJson map[string] interface{}
		err = json.Unmarshal(resp.Body.Bytes(), &expectedJson)
		if err != nil {
			t.Error(err)
		}

		if expectedJson["application-name"] != "resource monitoring" {
			t.Error("Invalid JSON format")
		}

		if expectedJson["status"] != "OK" {
			t.Error("Status != OK")
		}
	})
}

func TestHTTPResourceHandler(t *testing.T) {
	t.Run("it should return httpCode 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/resouces", nil)
		if err != nil {
			t.Error(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(StatusHandler)
		handler.ServeHTTP(resp, req)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("it should return correct format", func (t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/resouces", nil)
		if err != nil {
			t.Error(err)
		}

		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(ResourcesHandler)
		handler.ServeHTTP(resp, req)
		var expectedJson map[string] interface{}
		err = json.Unmarshal(resp.Body.Bytes(), &expectedJson)
		if err != nil {
			t.Error(err)
		}

		if expectedJson["Memory"] == nil {
			t.Error("Memory field must not be nil")
		}

		if expectedJson["Cpu"] == nil {
			t.Error("Cpu field must not be nil")
		}
	})
}
