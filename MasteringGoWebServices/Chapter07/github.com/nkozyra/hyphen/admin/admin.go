package admin

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	DB "github.com/nkozyra/hyphen/data"
	"html/template"
	"log"
	Password "nathankozyra.com/api/password"
	"net/http"
	"sync"
	"time"
)

var Router *mux.Router
var templates = template.Must(template.ParseGlob("templates/admin/*"))
var store = sessions.NewCookieStore([]byte("hyphen-session"))
var Session *sessions.Session
var AdminUser User

const serverName = "localhost"
const SSLport = ":443"
const HTTPport = ":8080"
const SSLprotocol = "https://"
const HTTPprotocol = "http://"
const templatesPath = "templates"

type User struct {
	LoggedIn bool
	UID      int
}

type Error struct {
	Code int
	Text string
}

type Page struct {
	Title  string
	Errors []Error
}

func Init() {
	Router = mux.NewRouter()
	Router.HandleFunc("/admin", ModuleDashboard).Methods("GET")
	Router.HandleFunc("/admin/login", ModuleLogin).Methods("GET")
	Router.HandleFunc("/admin/login", ModuleLoginProcess).Methods("POST")
	Router.HandleFunc("/register", ModuleRegister).Methods("GET")
	Router.HandleFunc("/register", ModuleRegisterProcess).Methods("POST")
	Router.HandleFunc("/welcome", ModuleAuthWelcome).Methods("GET")

	Router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))
}

func ValidateSession(r *http.Request) bool {
	S, _ := store.Get(r, "hyphen-session")

	Session = S
	if val, ok := Session.Values["sid"]; ok {
		fmt.Println(val)
		fmt.Println("WE HAVE COOKIE!")
		fmt.Println(Session.Values["sid"])
		var checkUID int
		DB.Database.QueryRow("SELECT groups_users_id from sessions where session_id=?", Session.Values["sid"]).Scan(&checkUID)
		if checkUID > 0 {
			AdminUser.LoggedIn = true
			AdminUser.UID = checkUID
			return true
		} else {
			return false
		}

		return true
	} else {

		fmt.Println("no such cookie")
		return false
	}

	return true
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UH")
	http.ServeFile(w, r, r.URL.Path[1:])

}

func ModuleDashboard(w http.ResponseWriter, r *http.Request) {

	if ValidateSession(r) == true {
		fmt.Println("ADMIN USER", AdminUser.UID)
		err := templates.ExecuteTemplate(w, "admin_dashboard.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Redirect(w, r, "/admin/login", http.StatusFound)
		return
	}
}

func ModuleAuthWelcome(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "auth_welcome.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ModuleRegisterProcess(w http.ResponseWriter, r *http.Request) {
	registerEmail := r.FormValue("registerEmail")
	//registerPublication := r.FormValue("registerPublication")
	registerPassword := r.FormValue("registerPassword")
	registerSalt, registerHash := Password.ReturnPassword(registerPassword)
	registerTime := time.Now().Unix()
	q, connectErr := DB.Database.Exec("INSERT INTO groups_users SET groups_users_guid=?, groups_organizations_id=?, groups_publications_id=?, groups_users_email=?, groups_users_joined=?, groups_users_password=?, groups_users_salt=?", "registerEmail", 0, 0, registerEmail, registerTime, registerHash, registerSalt)

	fmt.Println(q)
	if connectErr != nil {
		derr, dbError := DB.GetError(connectErr)
		if derr != nil {

		}
		fmt.Println(dbError)
		p := Page{}
		p.Errors = append(p.Errors, Error{Code: 1, Text: connectErr.Error()})
		err := templates.ExecuteTemplate(w, "auth_reg_error.html", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Redirect(w, r, "/welcome", http.StatusFound)
	}
}

func ModuleRegister(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func ModuleLoginProcess(w http.ResponseWriter, r *http.Request) {
	loginEmail := r.FormValue("loginEmail")
	loginPassword := r.FormValue("loginPassword")
	fmt.Println(loginPassword)

	var salt string
	var expectedPassword string
	var uid string
	DB.Database.QueryRow("SELECT groups_users_salt,groups_users_password,groups_users_id from groups_users WHERE groups_users_email=?", loginEmail).Scan(&salt, &expectedPassword, &uid)

	generatedPassword := Password.GenerateHash(salt, loginPassword)

	if generatedPassword == expectedPassword {
		fmt.Println("passwords match!")
		SID := Password.GenerateSessionID(64)
		Session.Values["sid"] = SID
		now := time.Now().Unix()

		DB.Database.Exec("INSERT into sessions set session_id=?,session_authenticated=1,session_start=?,groups_users_id=?", SID, now, uid)
		Session.Save(r, w)

	} else {
		fmt.Println("no password matcho")
		fmt.Println("generated", generatedPassword)
		fmt.Println("expected", expectedPassword)
	}

	fmt.Println(salt)
}

func ModuleLogin(w http.ResponseWriter, r *http.Request) {

	if ValidateSession(r) == true {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}

	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func redirectNonSecure(w http.ResponseWriter, r *http.Request) {
	log.Println("Non-secure request initiated, redirecting.")
	redirectURL := SSLprotocol + serverName + r.RequestURI
	http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
}

func StartServer() {
	DB.Init()
	fmt.Println("hey")

	wg := sync.WaitGroup{}

	log.Println("Starting redirection server, try to access @ http:")

	wg.Add(1)
	go func() {
		http.ListenAndServe(HTTPport, http.HandlerFunc(redirectNonSecure))
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		http.ListenAndServeTLS(SSLport, "cert.pem", "key.pem", Router)
		wg.Done()
	}()

	wg.Wait()

}
