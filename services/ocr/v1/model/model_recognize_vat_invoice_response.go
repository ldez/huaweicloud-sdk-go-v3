package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeVatInvoiceResponse struct {
	Result         *VatInvoiceResult `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o RecognizeVatInvoiceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeVatInvoiceResponse struct{}"
	}

	return strings.Join([]string{"RecognizeVatInvoiceResponse", string(data)}, " ")
}
