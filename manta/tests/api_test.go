package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/juliafem/manta/japp"
)

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// TestModelHandler Test for ModelHandler
func TestModelHandler(t *testing.T) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	f, err := os.Open("main.go")
	if err != nil {
		return
	}
	defer f.Close()

	fw, err := w.CreateFormFile("File", "test.inp")
	if err != nil {
		return
	}

	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	w.Close()

	req, _ := http.NewRequest("POST", "/model", &b)
	req.Header.Set("Content-Type", w.FormDataContentType())

	res := httptest.NewRecorder()

	handler := http.HandlerFunc(japp.ModelHandler)
	handler.ServeHTTP(res, req)

	checkResponseCode(t, http.StatusOK, res.Code)

}
