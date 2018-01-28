package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/urfave/negroni"
)

//connectionstring
//this should lay as a password encrypted file on the disk
var ConnStr = "X"

type Response struct {
	Data string `json:"data"`
}

func main() {
	Privkey = generateKeyPair()
	Pubkey = &Privkey.PublicKey
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", login)
	router.HandleFunc("/register", register)

	router.Handle("/protected", negroni.New(
		negroni.HandlerFunc(validateToken),
		negroni.Wrap(http.HandlerFunc(returnProtect)),
	))

	http.ListenAndServe(":8080", router)
}

func returnProtect(w http.ResponseWriter, r *http.Request) {
	resp := Response{"Protected resource accessed."}
	jsonResponse(resp, w)
}

func login(w http.ResponseWriter, r *http.Request) {
	//Ta input och kontrollera
	var user Luser
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	salt := dbGetSalt(user)
	user.Password = hashPass(user.Password, salt)
	name := dbLogin(user)

	//If correct password and username (default is to not accept)
	if name != "" {
		resp := createSendToken(w)
		jsonResponse(resp, w)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	//Ta input och kontrollera
	var user Ruser
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	exist := dbCheckUsername(user)

	//username does not already exist. Create the user
	if exist != 1 {
		//Create the salt and hash the password
		salt := make([]byte, 16)
		rand.Read(salt)
		user.Password = hashPass(user.Password, salt)
		user.Salt = base64.StdEncoding.EncodeToString(salt)

		err = dbRegister(user)
	}

	if exist != 1 {
		if err == nil {
			resp := Response{"User created successfully."}
			jsonResponse(resp, w)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func validateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return Pubkey, nil
		})
	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized) //443
			fmt.Fprintf(w, "Invalid token")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
	}
}

func createSendToken(w http.ResponseWriter) Token {
	//If username AND password OK. Create and send token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"foo": "bar", //could be anything like username etc.
		"nbf": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(Privkey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error signing token")
		panic(err)
	}

	resp := Token{tokenString}
	return resp
}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func hashPass(password string, salt []byte) string {
	nPass := append([]byte(password), salt...)
	sum := sha256.Sum256(nPass)
	newSum := sum[:]
	return base64.StdEncoding.EncodeToString(newSum)
}
