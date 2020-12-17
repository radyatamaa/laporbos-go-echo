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
	"github.com/sales_order"
)

type SalesOrder struct {
	SalesOrderRepo sales_order.Repository
	contextTimeout time.Duration
}

func NewSalesOrder(f sales_order.Repository, timeout time.Duration) sales_order.Usecase {
	return &SalesOrder{
		SalesOrderRepo: f,
		contextTimeout: timeout,
	}
}

func (f SalesOrder) GetAll(ctx context.Context, page, limit, offset int) (*models.SalesOrderDtoWithPagination, error) {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	getFacilities, err := f.SalesOrderRepo.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	facilitiesDtos := make([]*models.SalesOrderDto, 0)

	for _, element := range getFacilities {
		dto := models.SalesOrderDto{
			Id:                 element.Id,
			DocNumber:          element.DocNumber,
			DocDate:            element.DocDate,
			SaTy:                element.SaTy,
			Item:                element.Item,
			Material:            element.Material,
			Description:         element.Description,
			OrderQty:           element.OrderQty,
			NetPrice:           element.NetPrice,
			NetValue:            element.NetValue,
			Curr:                element.Curr,
			UoM:                 element.UoM,
			DlvDate:             element.DlvDate,
			Plant:              element.Plant,
			SalesOffice:        element.SalesOffice,
			SalesGroup:          element.SalesGroup,
			SalesOrg:           element.SalesOrg,
			DistributionChanel:  element.DistributionChanel,
			StorageLocation:    element.StorageLocation,
		}
		facilitiesDtos = append(facilitiesDtos, &dto)
	}
	totalRecords, _ := f.SalesOrderRepo.GetCount(ctx)

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

	response := &models.SalesOrderDtoWithPagination{
		Data: facilitiesDtos,
		Meta: meta,
	}
	return response, nil
}

func (f SalesOrder) Import(ctx context.Context, fileLocation string) error {
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
	SalesOrder := make([]*models.SalesOrder, 0)
	index := 0

	for rows.Next() {
		row := rows.Columns()
		if index != 0 && len(row) > 0 {
			item ,_:=strconv.Atoi(row[3])
			orderqty,_:=strconv.Atoi(row[6])
			netprice,_:= strconv.ParseFloat(row[7],64)
			netvalue,_:=strconv.ParseFloat(row[8],64)
			master := models.SalesOrder{
				Id:           0,
				CreatedBy:    "admin",
				CreatedDate:  time.Now(),
				ModifiedBy:   nil,
				ModifiedDate: nil,
				DeletedBy:    nil,
				DeletedDate:  nil,
				IsDeleted:    0,
				IsActive:     0,
				DocNumber:         row[0],
				DocDate:            row[1],
				SaTy:               row[2],
				Item:               item,
				Material:           row[4],
				Description:         row[5],
				OrderQty:          orderqty,
				NetPrice:           netprice,
				NetValue:           netvalue,
				Curr:                row[9],
				UoM:                row[10],
				DlvDate:            row[11],
				Plant:              row[12],
				SalesOffice:        row[13],
				SalesGroup:         row[14],
				SalesOrg:          row[15],
				DistributionChanel:  row[16],
				StorageLocation:    row[17],
			}
			SalesOrder = append(SalesOrder, &master)
			//fmt.Printf("%s\t%s\n", row[1], row[3]) // Print values in columns B and D
		}
		index = index + 1
	}
	for _, element := range SalesOrder {
		f.SalesOrderRepo.Insert(ctx, element)
	}
	//errRemove := os.Remove(fileLocation)
	//if errRemove != nil {
	//	return models.ErrInternalServerError
	//}
	return nil
}
