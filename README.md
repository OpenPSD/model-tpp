# model-tpp
## Model implementation for TPPs as AISP and PISP

- [x] Download the PSD2-server implementation 
- [x] Download the model-tpp implementation 
- [x] Call `go run main.go` of PSD2-server and model-tpp on a unix shell
- [x] Call `http://localhost:8080/test/payments/embedded` from Postman e.g. 

This will exceute a Payment initiation on the PSD2 server, the PSU authentification and the OTP verification. 

## A successfull call will return: 

### Test Create and authorise Payment with Embedded SCA

`Payment created:{
		"transactionStatus": "RCVD",
		"paymentId": "axa-pay-paymentid-1234",
		"_links": {
	"startAuthenticationWithPsuAuthentication": {"href": "/payments/axa-pay-paymentid-1234/authorisations"},
	"self": {"href": "/payments/axa-pay-paymentid-1234"} }
	}`

`Payment PSU authenticated:{
		"scaStatus": "psuAuthenticated",
	  "_links":{
	  "authoriseTransaction": {"href": "/payments/axa-pay-paymentid-1234/authorisations/123axa-auth456"}
		}
	  }`
	  
`Payment OTP verified:{
		"scaStatus": "finalised",
		"_links":{
	  "scaStatus": {"href":"/payments/axa-pay-paymentid-1234/authorisations/123axa-auth456"}
	  } }`
