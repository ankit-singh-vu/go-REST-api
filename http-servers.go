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



func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/typicode_get_posts", typicode_posts)
	http.HandleFunc("/post", mypost)
	http.ListenAndServe(":3000", nil)
}
