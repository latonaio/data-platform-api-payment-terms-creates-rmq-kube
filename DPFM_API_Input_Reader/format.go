package dpfm_api_input_reader

import (
	"data-platform-api-payment-terms-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToPaymentTerms() *requests.PaymentTerms {
	data := sdc.PaymentTerms
	return &requests.PaymentTerms{
		PaymentTerms:             data.PaymentTerms,
		DueDate:                  data.DueDate,
		BaseDateCalcFixedDate:    data.BaseDateCalcFixedDate,
		BaseDateCalcAddMonth:     data.BaseDateCalcAddMonth,
		PaymentDateCalcFixedDate: data.PaymentDateCalcFixedDate,
		PaymentDateAddMonth:      data.PaymentDateAddMonth,
	}
}

func (sdc *SDC) ConvertToPaymentTermsText() *requests.PaymentTermsText {
	dataPaymentTerms := sdc.PaymentTerms
	data := sdc.PaymentTerms.PaymentTermsText
	return &requests.PaymentTermsText{
		PaymentTerms:     dataPaymentTerms.PaymentTerms,
		Language:         data.Language,
		PaymentTermsName: data.PaymentTermsName,
	}
}
