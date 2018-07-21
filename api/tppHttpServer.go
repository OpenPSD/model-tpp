package api

import (
	"html/template"
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
func (s TppHTTPServer) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := LoginPageData{PageTitle: "AXA Pay Bank Login", BankID: s.BankID, PIN: s.PIN}
	//fmt.Fprint(w, "TPP server reference PISP Login\n")
	tmpl := template.Must(template.ParseFiles("api/login.html"))
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

	return routes
}
