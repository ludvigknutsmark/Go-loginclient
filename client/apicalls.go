package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
	"bytes"
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/* Login takes two parameters and returns an *errorString as string */
func Login(username string, password string) string {
	url := "https://securesite.press/verifyuser"
	//Creates a User struct to parse it to JSON and send as Post-data.
	u := User{Username: username, Password: password}
	jsonstring, _ := json.Marshal(u)
	//Creates a new request which sends the JSON data as buffer.
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonstring))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

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
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) == "0" {
		return errors.New("Please enter a valid username or password").Error()
	}
	/* Login successfull */
	return ""
}
/*RegisterUser is much like Login excepts it makes two requests to the server,
	one for checking for users with the same username as the sent parameter
	and one for creating a new user.*/
func RegisterUser(username string, password string, password2 string) string {

	/* Check if username exists */
	url := "https://securesite.press/getusername"
	us := User{Username: username, Password: "0"}
	jsonstring, _ := json.Marshal(us)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonstring))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Unable to get username").Error()
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if string(body) != "0" {
		return errors.New("Username already exists").Error()
	}

	url = "https://securesite.press/registeruser"
	if password != password2 {
		return errors.New("Passwords does not match").Error()
	}
	u := User{Username: username, Password: password}
	jsonstring, _ = json.Marshal(u)
	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonstring))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err = client.Do(req)
	if err != nil {
		return errors.New("Unable to create http client").Error()
	}
	defer resp.Body.Close()

	/*Register successfull*/
	return ""
}
