package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello world")
}

func typicode_posts(w http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	response := resp
	// fmt.Println("Response status:", resp.Status)
	// fmt.Fprint(w, resp.Body)
	// scanner := bufio.NewScanner(resp.Body)
	// for i := 0; scanner.Scan() && i < 5; i++ {
	// 	// fmt.Println()
	// 	fmt.Fprint(w, scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }
	

	// Set the Content-Type header to display JSON in the browser
	w.Header().Set("Content-Type", "application/json")

	// Copy the response body directly to the browser
	_, err = io.Copy(w, response.Body)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}


}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/typicode_get_posts", typicode_posts)
	http.ListenAndServe(":3000", nil)
}
