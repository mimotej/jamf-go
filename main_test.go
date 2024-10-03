package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "io"
)


func TestMainAPIEndpoint(t *testing.T){
    expected := "Hello World!"
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    w := httptest.NewRecorder()
    HelloWorld(w, req)
    result := w.Result()
    defer result.Body.Close()
    data, err := io.ReadAll(result.Body)
    if err != nil {
        t.Errorf("Error :%v", err)
    }
    if string(data) != expected {
        t.Errorf("REST API returned incorrect string. Expected : %q, but got: %q", expected, string(data))
    }


}