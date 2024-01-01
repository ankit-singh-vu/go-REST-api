package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

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

func mypost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Echo the received data back as the response
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}

	// // Print the received data to the console
	// fmt.Println("Received data:", string(body))

	// // Optionally, you can unmarshal the JSON and marshal it again to pretty print in console
	// var data interface{}
	// if err := json.Unmarshal(body, &data); err != nil {
	// 	fmt.Println("Error unmarshaling JSON:", err)
	// 	return
	// }
	// prettyJSON, err := json.MarshalIndent(data, "", "  ")
	// if err != nil {
	// 	fmt.Println("Error marshaling JSON:", err)
	// 	return
	// }
	// fmt.Println("Received JSON:")
	// fmt.Println(string(prettyJSON))
}

func mypost2(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into a Post struct
	var postData Post
	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Print specific keys from the received JSON data to the console
	fmt.Println("Received data - Title:", postData.Title)
	fmt.Println("Received data - Body:", postData.Body)

	// Construct the response JSON object with one key
	responseData := map[string]string{
		"title": postData.Title,
		"body": postData.Body}

	// Marshal the responseData into JSON
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error creating response JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Send the response JSON back to the client
	_, err = w.Write(responseJSON)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/typicode_get_posts", typicode_posts)
	http.HandleFunc("/post", mypost)
	http.HandleFunc("/post2", mypost2)
	http.ListenAndServe(":3000", nil)
}
