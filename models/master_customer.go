package models

import "time"

type MasterCustomer struct {
	Id               int     `json:"id" validate:"required"`
	CreatedBy        string     `json:"created_by":"required"`
	CreatedDate      time.Time  `json:"created_date" validate:"required"`
	ModifiedBy       *string    `json:"modified_by"`
	ModifiedDate     *time.Time `json:"modified_date"`
	DeletedBy        *string    `json:"deleted_by"`
	DeletedDate      *time.Time `json:"deleted_date"`
	IsDeleted        int        `json:"is_deleted" validate:"required"`
	IsActive         int        `json:"is_active" validate:"required"`
	MANDT string `json:"mandt"`
	KodeCustomers string `json:"kode_customers"`
	LAND1 string `json:"land_1"`
	NAME1 string `json:"name_1"`
}

type MasterCustomerDto struct {
	Id           int    `json:"id"`
	MANDT string `json:"mandt"`
	KodeCustomers string `json:"kode_customers"`
	LAND1 string `json:"land_1"`
	NAME1 string `json:"name_1"`
}
type NewCommandMasterCustomer struct {
	Id           int    `json:"id"`
	MANDT string `json:"mandt"`
	KodeCustomers string `json:"kode_customers"`
	LAND1 string `json:"land_1"`
	NAME1 string `json:"name_1"`
}
type MasterCustomerDtoWithPagination struct {
	Data []*MasterCustomerDto `json:"data"`
	Meta *MetaPagination    `json:"meta"`
}
