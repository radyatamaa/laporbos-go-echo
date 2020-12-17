package usecase

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/master_customer"
	"github.com/models"
)

type MasterCustomer struct {
	MasterCustomerRepo master_customer.Repository
	contextTimeout     time.Duration
}

func NewMasterCustomer(f master_customer.Repository, timeout time.Duration) master_customer.Usecase {
	return &MasterCustomer{
		MasterCustomerRepo: f,
		contextTimeout:     timeout,
	}
}

func (f MasterCustomer) GetAll(ctx context.Context, page, limit, offset int) (*models.MasterCustomerDtoWithPagination, error) {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	getFacilities, err := f.MasterCustomerRepo.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	facilitiesDtos := make([]*models.MasterCustomerDto, 0)

	for _, element := range getFacilities {
		dto := models.MasterCustomerDto{
			Id:            element.Id,
			MANDT:          element.MANDT,
			KodeCustomers:  element.KodeCustomers,
			LAND1:          element.LAND1,
			NAME1:          element.NAME1,
		}
		facilitiesDtos = append(facilitiesDtos, &dto)
	}
	totalRecords, _ := f.MasterCustomerRepo.GetCount(ctx)

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

	response := &models.MasterCustomerDtoWithPagination{
		Data: facilitiesDtos,
		Meta: meta,
	}
	return response, nil
}

func (f MasterCustomer) Import(ctx context.Context, fileLocation string) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	xlsx, err := excelize.OpenFile(fileLocation)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rows, err := xlsx.Rows("KNA1-Table Customer")
	if err != nil {
		log.Fatal(err)
	}
	masterCOA := make([]*models.MasterCustomer, 0)
	index := 0

	for rows.Next() {
		row := rows.Columns()
		if index != 0 && len(row) > 0 {
			master := models.MasterCustomer{
				Id:           0,
				CreatedBy:    "admin",
				CreatedDate:  time.Now(),
				ModifiedBy:   nil,
				ModifiedDate: nil,
				DeletedBy:    nil,
				DeletedDate:  nil,
				IsDeleted:    0,
				IsActive:     0,
				MANDT:         row[0],
				KodeCustomers:  row[1],
				LAND1:          row[2],
				NAME1:         row[3],
			}
			masterCOA = append(masterCOA, &master)
			//fmt.Printf("%s\t%s\n", row[1], row[3]) // Print values in columns B and D
		}
		index = index + 1
	}
	for _, element := range masterCOA {
		f.MasterCustomerRepo.Insert(ctx, element)
	}
	//errRemove := os.Remove(fileLocation)
	//if errRemove != nil {
	//	return models.ErrInternalServerError
	//}
	return nil
}
