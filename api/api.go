package api

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	// "github.com/gorilla/mux"

	webapp "github.com/kroppt/cs252-lab6-webapp/webapp"
)

// UserModel is a struct, which holds user information.
type UserModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type id struct {
	ID string
}

var storedID = id{""}

// AuthUser will check authentication of the given username (if valid)
// with the given password.
func AuthUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data *UserModel
	err := decoder.Decode(data)
	checkError(err)

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password),
		bcrypt.MaxCost)
	checkError(err)

	// check that conversion back is successful
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.Password))
	checkError(err)

	rows, err := webapp.DataBase.DB.Query(
		"SELECT User.PasswordHash, UserSession.SessionKey FROM User LEFT JOIN" +
			" UserSession ON (User.ID=UserSession.UserID) WHERE Username=" +
			data.Username)
	defer rows.Close()
	checkError(err)

	var rowData struct {
		PasswordHash string
		SessionKey   string
	}
	err = rows.Scan(&rowData)
	checkError(err)

	var output struct {
		Success bool
	}
	if rowData.PasswordHash == string(hash) {
		expireCookie := time.Now().Add(time.Hour * 1)
		cookie := http.Cookie{
			Name:     "Auth",
			Value:    string(rowData.SessionKey),
			Expires:  expireCookie,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		output.Success = true
		out, err := json.Marshal(output)
		checkError(err)
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	} else {
		output.Success = false
		out, err := json.Marshal(output)
		checkError(err)
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	}
	return
}

// GetID will return the ID stored by PostID.
// For testing purposes.
func GetID(w http.ResponseWriter, r *http.Request) {
	fmt.Print("GET ")
	fmt.Println(storedID)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, storedID.ID)
	return
}

// LoginUser authenticates and logs in the verified user, if possible.
// Creates a JWT and puts it into the client's cookie.
func LoginUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data *UserModel
	err := decoder.Decode(data)
	checkError(err)

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password),
		bcrypt.MaxCost)
	checkError(err)

	// check that conversion back is successful
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.Password))
	checkError(err)

	const mySQLDateTime = "2006-01-02 15:04:05"
	currTime := time.Now().Format(mySQLDateTime)

	rows, err := webapp.DataBase.DB.Query("SELECT ID, PasswordHash FROM" +
		"User WHERE Username=" + data.Username)

	var rowData struct {
		ID           int
		PasswordHash string
	}
	err = rows.Scan(&rowData)
	checkError(err)

	if string(hash) != rowData.PasswordHash {
		var output struct {
			Success bool
		}
		output.Success = false
		out, err := json.Marshal(output)
		checkError(err)
		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}

	expireCookie := time.Now().Add(time.Hour * 1)

	key := make([]byte, 50)
	_, err = rand.Read(key)
	checkError(err)

	rows, err = webapp.DataBase.DB.Query(
		"INSERT INTO UserSession (SessionKey," +
			" UserID," + " LoginTime, LastSeenTime) VALUES(" + string(key) +
			", LAST_INSERT_ID(), " + currTime + ", " + currTime + ");")
	defer rows.Close()
	checkError(err)

	cookie := http.Cookie{
		Name:     "Auth",
		Value:    string(key),
		Expires:  expireCookie,
		HttpOnly: true,
	}
	var output struct {
		Success bool
	}
	output.Success = true
	out, err := json.Marshal(output)
	checkError(err)
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 307)
	w.Write(out)
}

// LogoutUser logs out the verified user, if possible.
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Auth")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	_, err = webapp.DataBase.DB.Query(
		"DELETE FROM User WHERE UserSession.SessionKey=" + cookie.Value)
	checkError(err)
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
	http.SetCookie(w, &deleteCookie)
	var output struct {
		Success bool
	}
	output.Success = true
	out, err := json.Marshal(output)
	checkError(err)
	w.WriteHeader(http.StatusOK)
	w.Write(out)
	return
}

// NewUser checks authentication of the given username and password
// and creates a new user with those credentials.
func NewUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data *UserModel
	err := decoder.Decode(data)
	checkError(err)

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MaxCost)
	checkError(err)

	// check that conversion back is successful
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.Password))
	checkError(err)

	expireCookie := time.Now().Add(time.Hour * 1)

	key := make([]byte, 50)
	_, err = rand.Read(key)
	checkError(err)

	const mySQLDateTime = "2006-01-02 15:04:05"
	currTime := time.Now().Format(mySQLDateTime)

	rows, err := webapp.DataBase.DB.Query(
		"INSERT INTO User (Username, PasswordHash) VALUES(" + data.Username +
			", " + string(hash) + "); INSERT INTO UserSession (SessionKey," +
			" UserID," + " LoginTime, LastSeenTime) VALUES(" + string(key) +
			", LAST_INSERT_ID(), " + currTime + ", " + currTime + ");")
	defer rows.Close()
	checkError(err)

	cookie := http.Cookie{Name: "Auth", Value: string(key), Expires: expireCookie,
		HttpOnly: true}
	http.SetCookie(w, &cookie)
	var output struct {
		Success bool
	}
	output.Success = true
	out, err := json.Marshal(output)
	checkError(err)
	http.Redirect(w, r, "/", 307)
	w.Write(out)
}

// PostID will store the string to ID. The ID can be fetched using GetID.
// For testing purposes.
func PostID(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&storedID)
	checkError(err)
	fmt.Print("POST ")
	fmt.Println(storedID)
	w.WriteHeader(http.StatusOK)
	return
}

// TestDB will grab test data from the database.
// For testing purposes.
func TestDB(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	rows, err := webapp.DataBase.DB.Query("SELECT * FROM list")
	checkError(err)
	defer rows.Close()
	fmt.Fprintf(w, "list results:\n")
	for rows.Next() {
		var line string
		err := rows.Scan(&line)
		checkError(err)
		fmt.Fprintf(w, "%s\n", line)
	}
	return
}

// Validate checks user authentication and runs the handler if successful.
func Validate(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Auth")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		// SessionKey
		rows, err := webapp.
			DataBase.DB.Query("SELECT User.ID, UserSession.SessionKey" +
			" FROM User LEFT JOIN UserSession ON (User.ID=UserSession.UserID)" +
			" WHERE UserSession.SessionKey=" + cookie.Value)
		checkError(err)
		var data struct {
			UserID     int
			SessionKey string
		}
		err = rows.Scan(&data)
		checkError(err)
		rows.Close()
		if data.SessionKey == cookie.Value {
			const mySQLDateTime = "2006-01-02 15:04:05"
			currTime := time.Now().Format(mySQLDateTime)
			_, err = webapp.DataBase.DB.Query("UPDATE UserSession SET LastSeenTime=" +
				currTime + " WHERE UserSession.UserID=" + string(data.UserID))
			handler(w, r)
		} else {
			deleteCookie := http.Cookie{Name: "Auth", Value: "none",
				Expires: time.Now()}
			http.SetCookie(w, &deleteCookie)
			w.WriteHeader(http.StatusOK)
		}
		return
	})
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
