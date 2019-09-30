package functions

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArchiveSearch(t *testing.T) {
	videoData := []byte(`[{"index": "2" , "videoId": "abc", "url": "1234"}]`)
	r, _ := http.NewRequest("POST", "", bytes.NewBuffer(videoData))
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(ArchiveSearch)
	handler.ServeHTTP(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
	}

	// check there was no error getting the body
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// check the response body
	// stringBody := string(body)
	// jw := writers.NewMessageWriter(message)
	// messageString, _ := jw.JSONString()
	// if stringBody != messageString {
	// 	t.Errorf("wrong response body: got %v want %v", body, "Hello, World!\n")
	// }

}