package api

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
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

func callPSD2Server(psd2request []byte, psd2url string, psd2method string, w http.ResponseWriter, r *http.Request, p httprouter.Params) string {
	var htmlresponse string
	client := &http.Client{}
	req, err := http.NewRequest(
		psd2method, psd2url, bytes.NewBuffer(psd2request))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer resp.Body.Close()

	htmlresponsebyte, _ := ioutil.ReadAll(resp.Body)
	htmlresponse = string(htmlresponsebyte)
	fmt.Println("Antwort Body:", string(htmlresponse))
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return htmlresponse

}

//TestPaymentEmbedded tests a complete payment transaction with embedded approach
func (s TppHTTPServer) TestPaymentEmbedded(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Println("Test Create and authorise Payment with Embedded SCA")
	fmt.Fprintln(w, "Test Create and authorise Payment with Embedded SCA")
	iniatePayment := []byte(`{
		"instructedAmount": {"currency": "EUR", "amount": "123.50"},
		"debtorAccount": {"iban": "DE40100100103307118608"}, "creditorName": "Home24 via AXA Pay",
		"creditorAccount": {"iban": "DE02100100109307118603"}, "remittanceInformationUnstructured": "Home24.com"
		}`)
	htmlBody := callPSD2Server(iniatePayment, "http://127.0.0.1:8000/payments/sepa-credit-transfer", "POST", w, r, p)
	fmt.Fprintln(w, "Payment created:"+string(htmlBody))

	authorisePSU := []byte(`{
		"psuData": {
		  "password": "start12"
		} }`)
	htmlBody = callPSD2Server(authorisePSU, "http://localhost:8000/payments/axa-pay-paymentid-1234/authorisations/123axa-auth456", "PUT", w, r, p)
	fmt.Fprintln(w, "Payment PSU authenticated:"+string(htmlBody))

	authoriseTransaction := []byte(`{
		"scaAuthenticationData": "123456"
	   }`)
	htmlBody = callPSD2Server(authoriseTransaction, "http://localhost:8000/payments/axa-pay-paymentid-1234/authorisations/123axa-auth456", "PUT", w, r, p)
	fmt.Fprintln(w, "Payment OTP verified:"+string(htmlBody))

}

//NewMockedTppHTTPPServer creates an mocked TPP API server
func NewMockedTppHTTPServer(s TppHTTPServer) http.Handler {
	return TppHTTPServerFactory(s)
}

// TppHTTPServerFactory injects the required dependencies into the TPP API server
func TppHTTPServerFactory(s TppHTTPServer) http.Handler {
	routes := httprouter.New()
	routes.GET("/login", s.Login)
	routes.GET("/", s.AxaPay)
	routes.POST("/test/payments/embedded", s.TestPaymentEmbedded)

	return routes
}
