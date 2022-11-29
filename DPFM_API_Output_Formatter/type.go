package dpfm_api_output_formatter

type PaymentTerms struct {
	PaymentTerms             string           `json:"PaymentTerms"`
	DueDate                  int              `json:"DueDate"`
	BaseDateCalcFixedDate    int              `json:"BaseDateCalcFixedDate"`
	BaseDateCalcAddMonth     int              `json:"BaseDateCalcAddMonth"`
	PaymentDateCalcFixedDate int              `json:"PaymentDateCalcFixedDate"`
	PaymentDateAddMonth      int              `json:"PaymentDateAddMonth"`
	PaymentTermsText         PaymentTermsText `json:"PaymentTermsText"`
}

type PaymentTermsText struct {
	PaymentTerms     string `json:"PaymentTerms"`
	Language         string `json:"Language"`
	PaymentTermsName string `json:"PaymentTermsName"`
}
