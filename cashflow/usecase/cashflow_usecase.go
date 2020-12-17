package usecase

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/models"
	"github.com/cashflow"
)

type Cashflow struct {
	CashflowRepo   cashflow.Repository
	contextTimeout time.Duration
}

func NewCashflow(f cashflow.Repository, timeout time.Duration) cashflow.Usecase {
	return &Cashflow{
		CashflowRepo:   f,
		contextTimeout: timeout,
	}
}

func (f Cashflow) GetAll(ctx context.Context, page, limit, offset int) (*models.CashflowDtoWithPagination, error) {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	getFacilities, err := f.CashflowRepo.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	facilitiesDtos := make([]*models.CashflowDto, 0)

	for _, element := range getFacilities {
		dto := models.CashflowDto{
			Id:              element.Id,
			DocNumber:      element.DocNumber,
			DocDate:        element.DocDate,
			PostingDate:    element.PostingDate,
			COA:           element.COA,
			DC:             element.DC,
			Amount:         element.Amount,
			DocCurrency:    element.DocCurrency,
			DocHeader:      element.DocHeader,
			Assignment:    element.Assignment,
			SalesDocument:  element.SalesDocument,
			BillingDocument: element.BillingDocument,
		}
		facilitiesDtos = append(facilitiesDtos, &dto)
	}
	totalRecords, _ := f.CashflowRepo.GetCount(ctx)

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

	response := &models.CashflowDtoWithPagination{
		Data: facilitiesDtos,
		Meta: meta,
	}
	return response, nil
}

func (f Cashflow) Import(ctx context.Context, fileLocation string) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	xlsx, err := excelize.OpenFile(fileLocation)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rows, err := xlsx.Rows("Template SO")
	if err != nil {
		log.Fatal(err)
	}
	Cashflow := make([]*models.Cashflow, 0)
	index := 0

	for rows.Next() {
		row := rows.Columns()
		if index != 0 && len(row) > 0 {
			amount, _ := strconv.ParseFloat(row[5], 64)
			master := models.Cashflow{
				Id:                 0,
				CreatedBy:          "admin",
				CreatedDate:        time.Now(),
				ModifiedBy:         nil,
				ModifiedDate:       nil,
				DeletedBy:          nil,
				DeletedDate:        nil,
				IsDeleted:          0,
				IsActive:           0,
				DocNumber:     row[0],
				DocDate:        row[1],
				PostingDate:    row[2],
				COA:          row[3],
				DC:             row[4],
				Amount:       amount,
				DocCurrency:    row[6],
				DocHeader:     row[7],
				Assignment:   row[8],
				SalesDocument: row[9],
				BillingDocument: row[10],
			}
			Cashflow = append(Cashflow, &master)
			//fmt.Printf("%s\t%s\n", row[1], row[3]) // Print values in columns B and D
		}
		index = index + 1
	}
	for _, element := range Cashflow {
		f.CashflowRepo.Insert(ctx, element)
	}
	//errRemove := os.Remove(fileLocation)
	//if errRemove != nil {
	//	return models.ErrInternalServerError
	//}
	return nil
}
