package model

// Ref: https://www.api-docs.form3.tech/api/schemes/fps-direct/introduction/message-body-structure/data-section

type DataModel struct {
	Data  Data  `json:"data"`
	Links Links `json:"links,omitempty"`
}

type Data struct {
	ID             string     `json:"id"`
	OrganizationID string     `json:"organisation_id"`
	Type           string     `json:"type,omitempty"`
	Version        int64      `json:"version,omitempty"`
	Attributes     Attributes `json:"attributes,omitempty"`
}

type Attributes struct {
	AccountClassification   string   `json:"account_classification,omitempty"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 string   `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            bool     `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  string   `json:"status,omitempty"`
	Switched                bool     `json:"switched,omitempty"`
}

type Links struct {
	Self  string `json:"self"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
}
