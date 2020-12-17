package models

import "time"

type Ar struct {
	Id               int     `json:"id" validate:"required"`
	CreatedBy        string     `json:"created_by":"required"`
	CreatedDate      time.Time  `json:"created_date" validate:"required"`
	ModifiedBy       *string    `json:"modified_by"`
	ModifiedDate     *time.Time `json:"modified_date"`
	DeletedBy        *string    `json:"deleted_by"`
	DeletedDate      *time.Time `json:"deleted_date"`
	IsDeleted        int        `json:"is_deleted" validate:"required"`
	IsActive         int        `json:"is_active" validate:"required"`
	DocNumber string `json:"doc_number"`
	DocDate string `json:"doc_date"`
	PostingDate string `json:"posting_date"`
	TglJatuhTempo string `json:"tgl_jatuh_tempo"`
	Customer string `json:"customer"`
	DC string `json:"dc"`
	Amount float64 `json:"amount"`
	DocCurrency string `json:"doc_currency"`
	DocHeader string `json:"doc_header"`
	Assignment string `json:"assignment"`
	SalesDocument string `json:"sales_document"`
	BillingDocument string `json:"billing_document"`
}

type ArDto struct {
	Id           int    `json:"id"`
	DocNumber string `json:"doc_number"`
	DocDate string `json:"doc_date"`
	PostingDate string `json:"posting_date"`
	TglJatuhTempo string `json:"tgl_jatuh_tempo"`
	Customer string `json:"customer"`
	DC string `json:"dc"`
	Amount float64 `json:"amount"`
	DocCurrency string `json:"doc_currency"`
	DocHeader string `json:"doc_header"`
	Assignment string `json:"assignment"`
	SalesDocument string `json:"sales_document"`
	BillingDocument string `json:"billing_document"`
}
type NewCommandAr struct {
	Id           int    `json:"id"`
	DocNumber string `json:"doc_number"`
	DocDate string `json:"doc_date"`
	PostingDate string `json:"posting_date"`
	TglJatuhTempo string `json:"tgl_jatuh_tempo"`
	Customer string `json:"customer"`
	DC string `json:"dc"`
	Amount float64 `json:"amount"`
	DocCurrency string `json:"doc_currency"`
	DocHeader string `json:"doc_header"`
	Assignment string `json:"assignment"`
	SalesDocument string `json:"sales_document"`
	BillingDocument string `json:"billing_document"`
}
type ArDtoWithPagination struct {
	Data []*ArDto `json:"data"`
	Meta *MetaPagination    `json:"meta"`
}
