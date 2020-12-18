package usecase

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ar"
	"github.com/models"
)

type Ar struct {
	ArRepo         ar.Repository
	contextTimeout time.Duration
}

func NewAr(f ar.Repository, timeout time.Duration) ar.Usecase {
	return &Ar{
		ArRepo:         f,
		contextTimeout: timeout,
	}
}

func (f Ar) GetAll(ctx context.Context, page, limit, offset int) (*models.ArDtoWithPagination, error) {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	getFacilities, err := f.ArRepo.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	facilitiesDtos := make([]*models.ArDto, 0)

	for _, element := range getFacilities {
		dto := models.ArDto{
			Id:                 element.Id,
			DocNumber:          element.DocNumber,
			DocDate:            element.DocDate,
			PostingDate:        element.PostingDate,
			TglJatuhTempo:      element.TglJatuhTempo,
			Customer:             element.Customer,
			DC:                 element.DC,
			Amount:             element.Amount,
			DocCurrency:        element.DocCurrency,
			DocHeader:          element.DocHeader,
			Assignment:         element.Assignment,
			SalesDocument:      element.SalesDocument,
			BillingDocument: element.BillingDocument,
		}
		facilitiesDtos = append(facilitiesDtos, &dto)
	}
	totalRecords, _ := f.ArRepo.GetCount(ctx)

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

	response := &models.ArDtoWithPagination{
		Data: facilitiesDtos,
		Meta: meta,
	}
	return response, nil
}

func (f Ar) Import(ctx context.Context, fileLocation string) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	xlsx, err := excelize.OpenFile(fileLocation)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rows := xlsx.GetRows("FBL5N - SAP TMP")

	Ar := make([]*models.Ar, 0)
	index := 0

	for _, row := range rows {
		if index != 0 && len(row) > 0 {
			amount, _ := strconv.ParseFloat(row[6], 64)

			master := models.Ar{
				Id:                 0,
				CreatedBy:          "admin",
				CreatedDate:        time.Now(),
				ModifiedBy:         nil,
				ModifiedDate:       nil,
				DeletedBy:          nil,
				DeletedDate:        nil,
				IsDeleted:          0,
				IsActive:           0,
				DocNumber:          row[0],
				DocDate:            row[1],
				PostingDate:        row[2],
				TglJatuhTempo:      row[3],
				Customer:             row[4],
				DC:                 row[5],
				Amount:             amount,
				DocCurrency:        row[7],
				DocHeader:          row[8],
				Assignment:         row[9],
				SalesDocument:      row[10],
				BillingDocument: row[11],
			}
			Ar = append(Ar, &master)
			//fmt.Printf("%s\t%s\n", row[1], row[3]) // Print values in columns B and D
		}
		index = index + 1
	}
	for _, element := range Ar {
		f.ArRepo.Insert(ctx, element)
	}
	//errRemove := os.Remove(fileLocation)
	//if errRemove != nil {
	//	return models.ErrInternalServerError
	//}
	return nil
}
