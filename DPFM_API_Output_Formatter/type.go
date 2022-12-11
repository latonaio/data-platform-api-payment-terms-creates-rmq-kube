package dpfm_api_output_formatter

type PaymentTerms struct {
	PaymentTerms                string           `json:"PaymentTerms"`
	BaseDate                    int              `json:"BaseDate"`
	BaseDateCalcAddMonth        int              `json:"BaseDateCalcAddMonth"`
	BaseDateCalcFixedDate       int              `json:"BaseDateCalcFixedDate"`
	PaymentDueDateCalcAddMonth  int              `json:"PaymentDueDateCalcAddMonth"`
	PaymentDueDateCalcFixedDate int              `json:"PaymentDueDateCalcFixedDate"`
	PaymentTermsText            PaymentTermsText `json:"PaymentTermsText"`
}

type PaymentTermsText struct {
	PaymentTerms     string `json:"PaymentTerms"`
	Language         string `json:"Language"`
	PaymentTermsName string `json:"PaymentTermsName"`
}
