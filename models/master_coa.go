package models

import "time"

type MasterCOA struct {
	Id               int     `json:"id" validate:"required"`
	CreatedBy        string     `json:"created_by":"required"`
	CreatedDate      time.Time  `json:"created_date" validate:"required"`
	ModifiedBy       *string    `json:"modified_by"`
	ModifiedDate     *time.Time `json:"modified_date"`
	DeletedBy        *string    `json:"deleted_by"`
	DeletedDate      *time.Time `json:"deleted_date"`
	IsDeleted        int        `json:"is_deleted" validate:"required"`
	IsActive         int        `json:"is_active" validate:"required"`
	SPRAS string `json:"spras"`
	KTOPL string `json:"ktopl"`
	COA string `json:"coa"`
	TXT20 string `json:"txt_20"`
	TXT50 string `json:"txt_50"`
	MCOD1 string `json:"mcod_1"`
}

type MasterCOADto struct {
	Id           int    `json:"id"`
	SPRAS string `json:"spras"`
	KTOPL string `json:"ktopl"`
	COA string `json:"coa"`
	TXT20 string `json:"txt_20"`
	TXT50 string `json:"txt_50"`
	MCOD1 string `json:"mcod_1"`
}
type NewCommandMasterCOA struct {
	Id           int    `json:"id"`
	SPRAS string `json:"spras"`
	KTOPL string `json:"ktopl"`
	COA string `json:"coa"`
	TXT20 string `json:"txt_20"`
	TXT50 string `json:"txt_50"`
	MCOD1 string `json:"mcod_1"`
}
type MasterCOADtoWithPagination struct {
	Data []*MasterCOADto `json:"data"`
	Meta *MetaPagination    `json:"meta"`
}
