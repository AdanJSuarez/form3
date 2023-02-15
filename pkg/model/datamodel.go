package model

import (
	"encoding/json"
	"io"
)

// Ref: https://www.api-docs.form3.tech/api/schemes/fps-direct/introduction/message-body-structure/data-section

type DataModel struct {
	Data Data `json:"data"`
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

func NewDataModel(data Data) DataModel {
	return DataModel{
		Data: data,
	}
}

func NewData(ID, organizationID, resourceType string, version int64, attributes Attributes) Data {
	return Data{
		ID:             ID,
		OrganizationID: organizationID,
		Type:           resourceType,
		Version:        version,
		Attributes:     attributes,
	}
}

// func NewAttributes(
// 	AccountClassification   string
// 	AccountMatchingOptOut   bool
// 	AccountNumber           string
// 	AlternativeNames        []string
// 	BankID                  string
// BankIDCode              string
// BaseCurrency            string
// Bic                     string
// Country                 string
// Iban                    string
// JointAccount            bool
// Name                    []string
// SecondaryIdentification string
// Status                  string
// Switched                bool     )

func (d *DataModel) Unmarshal(responseBody io.ReadCloser) error {
	if err := json.NewDecoder(responseBody).Decode(d); err != nil {
		return err
	}
	return nil
}
