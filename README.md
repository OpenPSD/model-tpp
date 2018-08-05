# model-tpp
Model implementation for TPPs as AISP and PISP

1. Download the PSD2-serverimplementation 
2. Download the model-tpp implementation 
3. go run main.go of PSD2-server and model-tpp 
4. Call http://localhost:8080/test/payments/embedded 

A successfull call will return: 

Test Create and authorise Payment with Embedded SCA
Payment created:{
		"transactionStatus": "RCVD",
		"paymentId": "axa-pay-paymentid-1234",
		"_links": {
	"startAuthenticationWithPsuAuthentication": {"href": "/payments/axa-pay-paymentid-1234/authorisations"},
	"self": {"href": "/payments/axa-pay-paymentid-1234"} }
	}
Payment PSU authenticated:{
		"scaStatus": "psuAuthenticated",
	  "_links":{
	  "authoriseTransaction": {"href": "/payments/axa-pay-paymentid-1234/authorisations/123axa-auth456"}
		}
	  }
Payment OTP verified:{
		"scaStatus": "finalised",
		"_links":{
	  "scaStatus": {"href":"/payments/axa-pay-paymentid-1234/authorisations/123axa-auth456"}
	  } }
