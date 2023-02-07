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

func main() {
	f3 := form3.New()
	if err := f3.ConfigurationByValue(localhost, accountPath, organizationID); err != nil {
		log.Printf("Error on ConfigurationByValue: %v", err)
		return
	}
	account := f3.Account()

	data, err := account.Create(model.DataModel{
		Data: model.Data{
			ID:             generateUUID(),
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

	dataFetch, err := account.Fetch("020cf7d8-01b9-461d-89d4-89d57fd0d998")
	if err != nil {
		log.Printf("Error on Fetch: %v", err)
	}
	log.Println("Data Fetched: ", dataFetch)
}

func generateUUID() string {
	id := uuid.New()
	return id.String()
}

/*
- Include vendor or not: https://blog.boot.dev/golang/should-you-commit-the-vendor-folder-in-go/
- Has to have complete fields.
- Try twice the same message gives us a 409: conflict with existing state. ID has to uniq.



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
*/
