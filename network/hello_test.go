package network

//func handleError(t *testing.T, err error) {
//	t.Helper()
//	if err != nil {
//		t.Fatal("failed", err)
//	}
//}
//
//func TestConn(t *testing.T) {
//	ln, err := net.Listen("tcp", "127.0.0.1:0")
//	handleError(t, err)
//	defer ln.Close()
//
//	http.HandleFunc("/hello", helloHandler)
//	go http.Serve(ln, nil)
//
//	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
//	handleError(t, err)
//
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	handleError(t, err)
//
//	if string(body) != "hello world" {
//		t.Fatal("expected hello world, but got", string(body))
//	}
//}

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestConn(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}

func BenchmarkHello(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello")
	}
}
