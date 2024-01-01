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
	// := means declare and intialise in one line
	// resp, err :=  means if succes store the value in resp variable else store in err variable.
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	// nil means that a pointer does not point to any memory address.
	if err != nil {
		panic(err)
	}
	// defer means runs this at last
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

	// The first return value _ (underscore) is the number of bytes copied. As it is assigned to _, it signifies that you're discarding that value and not storing it in any variable.

	// The second return value err is an error. If io.Copy executes successfully, err will be nil, indicating no error occurred during the copying process. If an error occurs during the copy operation, it will be assigned to err.


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

func myget(w http.ResponseWriter, r *http.Request) {


	// Retrieve query parameters from the URL query string
	queryParams := r.URL.Query()
	fmt.Println("Query Parameters:")
	for key, values := range queryParams {
		fmt.Printf("%s: %s\n", key, values)
	}

	// Retrieve parameters from the URL path (e.g., /users/{id})
	params := r.URL.Path[len("/users/"):]
	fmt.Println("URL Parameters:", params)

	// Respond with a message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Query Parameters and URL Parameters received "+ params)
}

func main() {
	// http://localhost:3000/hello
	http.HandleFunc("/hello", hello)

	// http://localhost:3000/typicode_get_posts
	http.HandleFunc("/typicode_get_posts", typicode_posts)

	/*
		curl --location --request POST 'http://localhost:3000/post2' \
		--header 'Content-Type: application/json' \
		--data-raw '{
		"userId": 1,
		"id": 1,
		"title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
		"body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"
		}'
	*/
	http.HandleFunc("/post", mypost)

	http.HandleFunc("/post2", mypost2)

	// http://localhost:3000/users/123?name=John&age=30
	http.HandleFunc("/users/", myget)

	http.ListenAndServe(":3000", nil)
}
