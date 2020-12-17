package models

import "time"

type MasterVendor struct {
	Id               int     `json:"id" validate:"required"`
	CreatedBy        string     `json:"created_by":"required"`
	CreatedDate      time.Time  `json:"created_date" validate:"required"`
	ModifiedBy       *string    `json:"modified_by"`
	ModifiedDate     *time.Time `json:"modified_date"`
	DeletedBy        *string    `json:"deleted_by"`
	DeletedDate      *time.Time `json:"deleted_date"`
	IsDeleted        int        `json:"is_deleted" validate:"required"`
	IsActive         int        `json:"is_active" validate:"required"`
	KodeVendor string `json:"kode_vendor"`
	NAME1 string `json:"name_1"`
}

type MasterVendorDto struct {
	Id           int    `json:"id"`
	KodeVendor string `json:"kode_vendor"`
	NAME1 string `json:"name_1"`
}
type NewCommandMasterVendor struct {
	Id           int    `json:"id"`
	KodeVendor string `json:"kode_vendor"`
	NAME1 string `json:"name_1"`
}
type MasterVendorDtoWithPagination struct {
	Data []*MasterVendorDto `json:"data"`
	Meta *MetaPagination    `json:"meta"`
}
