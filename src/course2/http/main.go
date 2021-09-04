package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://httpbin.org/?a=1&b=2")

	if err != nil {
		fmt.Println("Error::", err)
		os.Exit(1)
	}
	fmt.Println(resp)

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Got this many bytes::", len(bs))
	return len(bs), nil
}
