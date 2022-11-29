package requests

type PaymentTerms struct {
	PaymentTerms             *string `json:"PaymentTerms"`
	DueDate                  *int    `json:"DueDate"`
	BaseDateCalcFixedDate    *int    `json:"BaseDateCalcFixedDate"`
	BaseDateCalcAddMonth     *int    `json:"BaseDateCalcAddMonth"`
	PaymentDateCalcFixedDate *int    `json:"PaymentDateCalcFixedDate"`
	PaymentDateAddMonth      *int    `json:"PaymentDateAddMonth"`
}
