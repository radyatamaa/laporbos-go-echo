package usecase

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ap"
	"github.com/models"
)

type Ap struct {
	ApRepo         ap.Repository
	contextTimeout time.Duration
}

func NewAp(f ap.Repository, timeout time.Duration) ap.Usecase {
	return &Ap{
		ApRepo:         f,
		contextTimeout: timeout,
	}
}

func (f Ap) GetAll(ctx context.Context, page, limit, offset int) (*models.ApDtoWithPagination, error) {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	getFacilities, err := f.ApRepo.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	facilitiesDtos := make([]*models.ApDto, 0)

	for _, element := range getFacilities {
		dto := models.ApDto{
			Id:                 element.Id,
			DocNumber:          element.DocNumber,
			DocDate:            element.DocDate,
			PostingDate:        element.PostingDate,
			TglJatuhTempo:      element.TglJatuhTempo,
			Vendor:             element.Vendor,
			DC:                 element.DC,
			Amount:             element.Amount,
			DocCurrency:        element.DocCurrency,
			DocHeader:          element.DocHeader,
			Assignment:         element.Assignment,
			SalesDocument:      element.SalesDocument,
			PurchasingDocument:  element.PurchasingDocument,
		}
		facilitiesDtos = append(facilitiesDtos, &dto)
	}
	totalRecords, _ := f.ApRepo.GetCount(ctx)

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

	response := &models.ApDtoWithPagination{
		Data: facilitiesDtos,
		Meta: meta,
	}
	return response, nil
}

func (f Ap) Import(ctx context.Context, fileLocation string) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	xlsx, err := excelize.OpenFile(fileLocation)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rows := xlsx.GetRows("FBL1N - SAP")

	Ap := make([]*models.Ap, 0)
	index := 0

	for _, row := range rows {
		if index != 0 && len(row) > 0 {
			amount, _ := strconv.ParseFloat(row[6], 64)

			master := models.Ap{
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
				Vendor:             row[4],
				DC:                 row[5],
				Amount:             amount,
				DocCurrency:        row[7],
				DocHeader:          row[8],
				Assignment:          row[9],
				SalesDocument:      row[10],
				PurchasingDocument:   row[11],
			}
			Ap = append(Ap, &master)
			//fmt.Printf("%s\t%s\n", row[1], row[3]) // Print values in columns B and D
		}
		index = index + 1
	}
	for _, element := range Ap {
		f.ApRepo.Insert(ctx, element)
	}
	//errRemove := os.Remove(fileLocation)
	//if errRemove != nil {
	//	return models.ErrInternalServerError
	//}
	return nil
}
