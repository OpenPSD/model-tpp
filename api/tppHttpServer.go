package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// LoginPageData holds all page data
type LoginPageData struct {
	PageTitle string
	BankID    string
	PIN       string
}

// TppHTTPServer holds all http handlers for the TPP API
type TppHTTPServer struct {
	BankID string // BankID
	PIN    string // PIN
}

// Login function for TPP PISP Login
func (s TppHTTPServer) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data := LoginPageData{PageTitle: "AXA Pay Bank Login", BankID: s.BankID, PIN: s.PIN}
	//fmt.Fprint(w, "TPP server reference PISP Login\n")
	tmpl := template.Must(template.ParseFiles("api/login.html"))
	r.ParseForm()
	log.Println("Form:", r.Form)
	tmpl.Execute(w, data)

}
func (s TppHTTPServer) AxaPay(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data := LoginPageData{PageTitle: "AXA Pay Bank Login", BankID: s.BankID, PIN: s.PIN}
	//fmt.Fprint(w, "TPP server reference PISP Login\n")
	tmpl := template.Must(template.ParseFiles("api/button.html"))
	tmpl.Execute(w, data)

}

// NewMockedTppHTTPPServer creates an mocked TPP API server
func NewMockedTppHTTPServer(s TppHTTPServer) http.Handler {
	return TppHTTPServerFactory(s)
}

// TppHTTPServerFactory injects the required dependencies into the TPP API server
func TppHTTPServerFactory(s TppHTTPServer) http.Handler {
	routes := httprouter.New()
	routes.GET("/login", s.Login)
	routes.GET("/", s.AxaPay)

	return routes
}
