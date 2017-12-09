package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pass10, err := base64.StdEncoding.DecodeString(data.Password)
	data.Password = string(pass10)

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password),
		bcrypt.MaxCost)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// check that conversion back is successful
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	rows, err := webapp.DataBase.DB.Query(
		"SELECT User.PasswordHash, UserSession.SessionKey FROM User LEFT JOIN" +
			" UserSession ON (User.ID=UserSession.UserID) WHERE Username=" +
			data.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rowData struct {
		PasswordHash string
		SessionKey   string
	}
	err = rows.Scan(&rowData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var output struct {
		Success bool `json:"success"`
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
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	} else {
		output.Success = false
		out, err := json.Marshal(output)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pass10, err := base64.StdEncoding.DecodeString(data.Password)
	data.Password = string(pass10)

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password),
		bcrypt.MaxCost)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// check that conversion back is successful
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	const mySQLDateTime = "2006-01-02 15:04:05"
	currTime := time.Now().Format(mySQLDateTime)

	rows, err := webapp.DataBase.DB.Query("SELECT ID, PasswordHash FROM" +
		"User WHERE Username=" + data.Username)

	var rowData struct {
		ID           int
		PasswordHash string
	}
	err = rows.Scan(&rowData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if string(hash) != rowData.PasswordHash {
		var output struct {
			Success bool `json:"success"`
		}
		output.Success = false
		out, err := json.Marshal(output)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}

	expireCookie := time.Now().Add(time.Hour * 1)

	key := make([]byte, 50)
	_, err = rand.Read(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rows, err = webapp.DataBase.DB.Query(
		"INSERT INTO UserSession (SessionKey," +
			" UserID," + " LoginTime, LastSeenTime) VALUES(" + string(key) +
			", LAST_INSERT_ID(), " + currTime + ", " + currTime + ");")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	cookie := http.Cookie{
		Name:     "Auth",
		Value:    string(key),
		Expires:  expireCookie,
		HttpOnly: true,
	}
	var output struct {
		Success bool `json:"success"`
	}
	output.Success = true
	out, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	w.Write(out)
}

// LogoutUser logs out the verified user, if possible.
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Auth")
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	_, err = webapp.DataBase.DB.Query(
		"DELETE FROM User WHERE UserSession.SessionKey=" + cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		return
	}
	deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
	http.SetCookie(w, &deleteCookie)
	var output struct {
		Success bool `json:"success"`
	}
	output.Success = true
	out, err := json.Marshal(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(out)
	return
}

// NewUser checks authentication of the given username and password
// and creates a new user with those credentials.
func NewUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data UserModel
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	l := len(data.Username)
	if !(l < 20 && l > 6) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect username length."))
		return
	}

	l = len(data.Password)
	if !(l < 50 && l > 10) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect password length."))
		return
	}

	row := webapp.DataBase.DB.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM User WHERE Username='" + data.Username + "');")
	var exists int
	err = row.Scan(&exists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If attempting to create an already-existing user
	if exists == 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username already in use."))
		return
	}

	pass10, err := base64.StdEncoding.DecodeString(data.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	data.Password = string(pass10)

	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.MinCost)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// check that conversion back is successful
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(data.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	key := (func() []byte {
		const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
		seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
		k := make([]byte, 50)
		for i := range k {
			k[i] = charset[seededRand.Intn(len(charset))]
		}
		return k
	})()

	expireCookie := time.Now().Add(time.Hour * 1)

	const mySQLDateTime = "2006-01-02 15:04:05"
	currTime := time.Now().Format(mySQLDateTime)

	result, err := webapp.DataBase.DB.Exec(
		"INSERT INTO User (Username, PasswordHash) VALUES ('" + data.Username +
			"', '" + string(hash) + "');")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	lastID, _ := result.LastInsertId()
	result, err = webapp.DataBase.DB.Exec(
		"INSERT INTO UserSession (SessionKey," +
			" UserID," + " LoginTime, LastSeenTime) VALUES(\n'" + string(key) +
			"',\n " + strconv.FormatInt(lastID, 10) + ",\n '" + currTime +
			"',\n '" + currTime + "');")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{Name: "Auth", Value: string(key), Expires: expireCookie,
		HttpOnly: true}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// PostID will store the string to ID. The ID can be fetched using GetID.
// For testing purposes.
func PostID(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&storedID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Print("POST ")
	fmt.Println(storedID)
	w.WriteHeader(http.StatusOK)
	return
}

// TestDB will grab test data from the database.
// For testing purposes.
func TestDB(w http.ResponseWriter, r *http.Request) {
	rows, err := webapp.DataBase.DB.Query("SELECT * FROM list")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	fmt.Fprintf(w, "list results:\n")
	for rows.Next() {
		var line string
		err := rows.Scan(&line)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s\n", line)
	}
	w.WriteHeader(http.StatusOK)
	return
}

// Validate checks user authentication and runs the handler if successful.
func Validate(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Auth")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// SessionKey
		rows, err := webapp.
			DataBase.DB.Query("SELECT User.ID, UserSession.SessionKey" +
			" FROM User LEFT JOIN UserSession ON (User.ID=UserSession.UserID)" +
			" WHERE UserSession.SessionKey=" + cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var data struct {
			UserID     int
			SessionKey string
		}
		err = rows.Scan(&data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		rows.Close()
		if data.SessionKey == cookie.Value {
			const mySQLDateTime = "2006-01-02 15:04:05"
			currTime := time.Now().Format(mySQLDateTime)
			_, err = webapp.DataBase.DB.Query("UPDATE UserSession SET LastSeenTime=" +
				currTime + " WHERE UserSession.UserID=" + string(data.UserID))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			handler(w, r)
		} else {
			deleteCookie := http.Cookie{Name: "Auth", Value: "none",
				Expires: time.Now()}
			http.SetCookie(w, &deleteCookie)
			w.WriteHeader(http.StatusUnauthorized)
		}
		return
	})
}
