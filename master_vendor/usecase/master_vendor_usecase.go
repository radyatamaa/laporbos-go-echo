package usecase

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/master_vendor"
	"github.com/models"
)

type MasterVendor struct {
	MasterVendorRepo master_vendor.Repository
	contextTimeout   time.Duration
}

func NewMasterVendor(f master_vendor.Repository, timeout time.Duration) master_vendor.Usecase {
	return &MasterVendor{
		MasterVendorRepo: f,
		contextTimeout:   timeout,
	}
}

func (f MasterVendor) GetAll(ctx context.Context, page, limit, offset int) (*models.MasterVendorDtoWithPagination, error) {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	getFacilities, err := f.MasterVendorRepo.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	facilitiesDtos := make([]*models.MasterVendorDto, 0)

	for _, element := range getFacilities {
		dto := models.MasterVendorDto{
			Id:    element.Id,
			KodeVendor: element.KodeVendor,
			NAME1: element.NAME1,
		}
		facilitiesDtos = append(facilitiesDtos, &dto)
	}
	totalRecords, _ := f.MasterVendorRepo.GetCount(ctx)

	totalPage := int(math.Ceil(float64(totalRecords) / float64(limit)))
	prev := page
	next := page
	if page != 1 {
		prev = page - 1
	}

	if page != totalPage {
		next = page + 1
	}

	meta := &models.MetaPagination{
		Page:          page,
		Total:         totalPage,
		TotalRecords:  totalRecords,
		Prev:          prev,
		Next:          next,
		RecordPerPage: len(facilitiesDtos),
	}

	response := &models.MasterVendorDtoWithPagination{
		Data: facilitiesDtos,
		Meta: meta,
	}
	return response, nil
}

func (f MasterVendor) Import(ctx context.Context, fileLocation string) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	xlsx, err := excelize.OpenFile(fileLocation)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rows, err := xlsx.Rows("LFA1-Master vendor")
	if err != nil {
		log.Fatal(err)
	}
	masterCOA := make([]*models.MasterVendor, 0)
	index := 0

	for rows.Next() {
		row := rows.Columns()
		if index != 0 && len(row) > 0 {
			master := models.MasterVendor{
				Id:           0,
				CreatedBy:    "admin",
				CreatedDate:  time.Now(),
				ModifiedBy:   nil,
				ModifiedDate: nil,
				DeletedBy:    nil,
				DeletedDate:  nil,
				IsDeleted:    0,
				IsActive:     0,
				KodeVendor:        row[0],
				NAME1:        row[1],
			}
			masterCOA = append(masterCOA, &master)
			//fmt.Printf("%s\t%s\n", row[1], row[3]) // Print values in columns B and D
		}
		index = index + 1
	}
	for _, element := range masterCOA {
		f.MasterVendorRepo.Insert(ctx, element)
	}
	//errRemove := os.Remove(fileLocation)
	//if errRemove != nil {
	//	return models.ErrInternalServerError
	//}
	return nil
}
