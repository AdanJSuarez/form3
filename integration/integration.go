package main

import (
	"log"

	"github.com/AdanJSuarez/form3/pkg/form3"
	"github.com/AdanJSuarez/form3/pkg/model"
	"github.com/google/uuid"
)

const (
	organizationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c"
	localhost      = "http://localhost:8080"
	accountPath    = "/v1/organisation/accounts"
)

var jitUUID = generateUUID()

func main() {
	f3 := form3.New()
	if err := f3.ConfigurationByValue(localhost, accountPath); err != nil {
		log.Printf("Error on ConfigurationByValue: %v", err)
		return
	}
	account := f3.Account()

	data, err := account.Create(model.DataModel{
		Data: model.Data{
			ID:             jitUUID,
			OrganizationID: organizationID,
			Type:           "accounts",
			Version:        1,
			Attributes: model.Attributes{
				Country: "GB",
				// BaseCurrency: "GBP",
				BankID:     "123456",
				BankIDCode: "GBDSC",
				Bic:        "EXMPLGB2XXX",
				Name:       []string{"a", "b"},
			},
		}})
	if err != nil {
		log.Printf("Error on Create: %v", err)
		return
	}
	log.Println("Data: ", data)

	dataFetch, err := account.Fetch(jitUUID) //"020cf7d8-01b9-461d-89d4-89d57fd0d998"
	if err != nil {
		log.Printf("Error on Fetch: %v", err)
	}
	log.Println("Data Fetched: ", dataFetch)

	if err := account.Delete(jitUUID, 1); err != nil {
		log.Printf("Error on Delete: %v", err)
	}
}

func generateUUID() string {
	id := uuid.New()
	return id.String()
}

/*
- Include vendor or not: https://blog.boot.dev/golang/should-you-commit-the-vendor-folder-in-go/
- Has to have complete fields. Check which field are forced by the Fake API
- Try twice the same message gives us a 409: conflict with existing state. ID has to uniq.
- Review the DataModel to remove omitempty of those that are not optional
- Headers are not check in the Fake API. Should the test fail in that case.


	  "country": "GB",
      "base_currency": "GBP",
      "bank_id": "123456",
      "bank_id_code": "GBDSC",
      "bic": "EXMPLGB2XXX",
      "user_defined_data": [
        {
          "key": "account_related_key",
          "value": "account_related_value"
        }
      ],
      "validation_type": "card",
      "reference_mask": "############",
      "acceptance_qualifier": "same_day",
      "switched_account_details": {
        "switched_effective_date": "2022-07-23",
        "account_number": "12345678",
        "account_with": {
          "bank_id": "123456",
          "bank_id_code": "GBDSC"
        },
        "account_number_code": "BBAN",
        "account_type": 0
      }

Comments:
Hi 👋

According to the documentation the following fields are deprecated:

title (superseded by name)
first_name
bank_account_name

alternative_bank_account_names (superseded by alternative_names)

However those changes don't seem to be implemented in the docker api provided - form3tech/interview-accountapi:v1.0.0-4-g63cf8434. The fields name and alternative_names are being ignored.

Additionally the switched field doesn't persist in the database nor is returned from the API on account creation.

-----------------

Hi,

During account create op for UK account with minimal payload(non CoP) is returning 400 error with message
"account_classification in body should be one of [Personal Business]"

But according to the doc it's a non required param and only used for the CoP request.

This is creating a bit of confusion on required and non required data.

--

I am not sure if this is the expected behaviour but I have figured what was the issue. I was testing the Create API without omitting empty fields and thus having "account_classification": "" in the payload.

Failed Payload: failed-payload.txt

According to the doc almost all the CoP fields are optional except the name field. Strangely, when I removed only the account_classification field it magically worked but I suspect that it shouldn't!

The required CoP name field that I have passed should have triggered a validation error as this is a required field and I have passed an array of empty strings!
"name": [ "", "", "", "" ]

Also, When I am passing the name and alternative_names correctly, they are not being set correctly.
Here is payload: req-resp-payload.txt

No account_number and iban are being generated. Also status field is missing in the response.



*/
