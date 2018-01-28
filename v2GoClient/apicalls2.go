package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* Login takes two parameters and returns an *errorString as string */
func Login(username string, password string) (int, string, string) {
	fmt.Println()
	url := "http://localhost:8080/login"

	//Creates a User struct to parse it to JSON and send as Post-data.
	u := User{Username: username, Password: password}
	jsonstring, _ := json.Marshal(u)
	//Creates a new request which sends the JSON data as buffer.
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonstring))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	//http.Client has a 10 second timeout before aborting an attempted connection
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, "", errors.New("Unable to create http client").Error()
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)
	resp.Body.Close()
	if resp.StatusCode == 200 {
		return 200, buf.String(), ""
	}
	return resp.StatusCode, "", errors.New("Please enter a valid username or password").Error()
}

func AccessProtected(token string) string {
	url := "http://localhost:8080/protected"
	//Get request for the passed URL
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	//Sets the token as a Authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	//http.Client has a 10 second timeout before aborting an attempted connection
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Unable to create http client").Error()
	}
	//Gets the return-data from the server and reads it as a string.
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return ""
	} else {
		return errors.New("Unauthorized").Error()
	}

}

/*RegisterUser is much like Login excepts it makes two requests to the server,
one for checking for users with the same username as the sent parameter
and one for creating a new user.*/
func RegisterUser(username string, password string, password2 string) string {
	if password != password2 {
		return errors.New("Passwords does not match.").Error()
	}

	/* Register user */
	url := "http://localhost:8080/register"
	us := User{Username: username, Password: password}
	jsonstring, _ := json.Marshal(us)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonstring))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Unable to get username").Error()
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return ""
	} else {
		return errors.New("Username already exists").Error()
	}
}
