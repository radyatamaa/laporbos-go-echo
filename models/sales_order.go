package models

import "time"

type SalesOrder struct {
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
	SaTy string `json:"sa_ty"`
	Item int `json:"item"`
	Material string `json:"material"`
	Description string `json:"description"`
	OrderQty int `json:"order_qty"`
	NetPrice float64 `json:"net_price"`
	NetValue float64 `json:"net_value"`
	Curr string `json:"curr"`
	UoM string `json:"uo_m"`
	DlvDate string `json:"dlv_date"`
	Plant string `json:"plant"`
	SalesOffice string `json:"sales_office"`
	SalesGroup string `json:"sales_group"`
	SalesOrg string `json:"sales_org"`
	DistributionChanel string `json:"distribution_chanel"`
	StorageLocation string `json:"storage_location"`
}

type SalesOrderDto struct {
	Id           int    `json:"id"`
	DocNumber string `json:"doc_number"`
	DocDate string `json:"doc_date"`
	SaTy string `json:"sa_ty"`
	Item int `json:"item"`
	Material string `json:"material"`
	Description string `json:"description"`
	OrderQty int `json:"order_qty"`
	NetPrice float64 `json:"net_price"`
	NetValue float64 `json:"net_value"`
	Curr string `json:"curr"`
	UoM string `json:"uo_m"`
	DlvDate string `json:"dlv_date"`
	Plant string `json:"plant"`
	SalesOffice string `json:"sales_office"`
	SalesGroup string `json:"sales_group"`
	SalesOrg string `json:"sales_org"`
	DistributionChanel string `json:"distribution_chanel"`
	StorageLocation string `json:"storage_location"`
}
type NewCommandSalesOrder struct {
	Id           int    `json:"id"`
	DocNumber string `json:"doc_number"`
	DocDate string `json:"doc_date"`
	SaTy string `json:"sa_ty"`
	Item int `json:"item"`
	Material string `json:"material"`
	Description string `json:"description"`
	OrderQty int `json:"order_qty"`
	NetPrice float64 `json:"net_price"`
	NetValue float64 `json:"net_value"`
	Curr string `json:"curr"`
	UoM string `json:"uo_m"`
	DlvDate string `json:"dlv_date"`
	Plant string `json:"plant"`
	SalesOffice string `json:"sales_office"`
	SalesGroup string `json:"sales_group"`
	SalesOrg string `json:"sales_org"`
	DistributionChanel string `json:"distribution_chanel"`
	StorageLocation string `json:"storage_location"`
}
type SalesOrderDtoWithPagination struct {
	Data []*SalesOrderDto `json:"data"`
	Meta *MetaPagination    `json:"meta"`
}
