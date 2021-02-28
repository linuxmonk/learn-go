package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {

	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(req.Context(), time.Duration(150)*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)

	tr := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Response: ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
