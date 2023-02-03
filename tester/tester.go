package main

import (
	"log"

	"github.com/AdanJSuarez/form3"
	"github.com/AdanJSuarez/form3/model"
	"github.com/google/uuid"
)

func main() {
	form3, err := form3.New("http://localhost:8080")
	if err != nil {
		log.Printf("Error on New: %v", err)
		return
	}
	account := form3.Account("/v1/organisation/accounts")

	data, err := account.Create(model.Data{
		Data: model.Account{
			ID:             generateUUID(),
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type:           "accounts",
			Version:        1,
			Attributes: model.AccountAttributes{
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
