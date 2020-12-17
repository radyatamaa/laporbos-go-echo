package usecase

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/master_coa"
	"github.com/models"
)

type MasterCOA struct {
	MasterCOARepo   master_coa.Repository
	contextTimeout time.Duration
}



func NewMasterCOA(f master_coa.Repository, timeout time.Duration) master_coa.Usecase {
	return &MasterCOA{
		MasterCOARepo:   f,
		contextTimeout: timeout,
	}
}


func (f MasterCOA) GetAll(ctx context.Context, page, limit, offset int) (*models.MasterCOADtoWithPagination, error) {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	getFacilities, err := f.MasterCOARepo.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	facilitiesDtos := make([]*models.MasterCOADto, 0)

	for _, element := range getFacilities {
		dto := models.MasterCOADto{
			Id:    element.Id,
			SPRAS: element.SPRAS,
			KTOPL:element.KTOPL,
			COA:   element.COA,
			TXT20: element.TXT20,
			TXT50:element.TXT50,
			MCOD1: element.MCOD1,
		}
		facilitiesDtos = append(facilitiesDtos, &dto)
	}
	totalRecords, _ := f.MasterCOARepo.GetCount(ctx)

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

	response := &models.MasterCOADtoWithPagination{
		Data: facilitiesDtos,
		Meta: meta,
	}
	return response, nil
}

func (f MasterCOA) Import(ctx context.Context, fileLocation string) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()


	xlsx, err := excelize.OpenFile( fileLocation)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rows, err := xlsx.Rows("SKAT-Master Account")
	if err != nil {
		log.Fatal(err)
	}
	masterCOA := make([]*models.MasterCOA,0)
	index := 0

	for rows.Next() {
		row := rows.Columns()
		if index != 0 && len(row) > 0{
			master := models.MasterCOA{
				Id:           0,
				CreatedBy:    "admin",
				CreatedDate:  time.Now(),
				ModifiedBy:   nil,
				ModifiedDate: nil,
				DeletedBy:    nil,
				DeletedDate:  nil,
				IsDeleted:    0,
				IsActive:     0,
				SPRAS:        row[0],
				KTOPL:        row[1],
				COA:          row[2],
				TXT20:        row[3],
				TXT50:        row[4],
				MCOD1:        row[5],
			}
			masterCOA = append(masterCOA,&master)
			//fmt.Printf("%s\t%s\n", row[1], row[3]) // Print values in columns B and D
		}
		index = index + 1
	}
	for _,element := range masterCOA{
		f.MasterCOARepo.Insert(ctx,element)
	}
	//errRemove := os.Remove(fileLocation)
	//if errRemove != nil {
	//	return models.ErrInternalServerError
	//}
	return nil
}
