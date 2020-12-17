package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"

	model "github.com/models"
)

func main() {
	//dev
	db, err := gorm.Open("mysql", "root:@(localhost)/laporbos_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	mig := model.MigrationHistory{}
	db.AutoMigrate(&mig)
	migration1 := model.MigrationHistory{
		DescMigration: "Add table MigrationHistory",
		Date:          time.Now(),
	}

	db.Create(&migration1)

	ap := model.Ap{}
	db.AutoMigrate(&ap)
	migration2 := model.MigrationHistory{
		DescMigration: "Add table Ap",
		Date:          time.Now(),
	}

	db.Create(&migration2)


	Ar := model.Ar{}
	db.AutoMigrate(&Ar)
	migration3 := model.MigrationHistory{
		DescMigration: "Add table Ar",
		Date:          time.Now(),
	}

	db.Create(&migration3)

	MasterCOA := model.MasterCOA{}
	db.AutoMigrate(&MasterCOA)
	migration4 := model.MigrationHistory{
		DescMigration: "Add table MasterCOA",
		Date:          time.Now(),
	}

	db.Create(&migration4)

	MasterCustomer := model.MasterCustomer{}
	db.AutoMigrate(&MasterCustomer)
	migration5 := model.MigrationHistory{
		DescMigration: "Add table MasterCustomer",
		Date:          time.Now(),
	}

	db.Create(&migration5)

	MasterVendor := model.MasterVendor{}
	db.AutoMigrate(&MasterVendor)
	migration6 := model.MigrationHistory{
		DescMigration: "Add table MasterVendor",
		Date:          time.Now(),
	}

	db.Create(&migration6)

	SalesOrder := model.SalesOrder{}
	db.AutoMigrate(&SalesOrder)
	migration7 := model.MigrationHistory{
		DescMigration: "Add table SalesOrder",
		Date:          time.Now(),
	}

	db.Create(&migration7)

}
