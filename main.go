package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_articleHttpDeliver "github.com/bxcodec/go-clean-arch/article/delivery/http"
	_articleRepo "github.com/bxcodec/go-clean-arch/article/repository"
	_articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	_authorRepo "github.com/bxcodec/go-clean-arch/author/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	_echoMiddleware "github.com/labstack/echo/middleware"

	_masterCOAHttpDeliver "github.com/master_coa/delivery/http"
	_masterCOARepo "github.com/master_coa/repository"
	_masterCOAUcase "github.com/master_coa/usecase"

	_masterVendorHttpDeliver "github.com/master_vendor/delivery/http"
	_masterVendorRepo "github.com/master_vendor/repository"
	_masterVendorUcase "github.com/master_vendor/usecase"

	_masterCustomerHttpDeliver "github.com/master_customer/delivery/http"
	_masterCustomerRepo "github.com/master_customer/repository"
	_masterCustomerUcase "github.com/master_customer/usecase"

	_salesOrderHttpDeliver "github.com/sales_order/delivery/http"
	_salesOrderRepo "github.com/sales_order/repository"
	_salesOrderUcase "github.com/sales_order/usecase"
)

func main() {
	//accountStorage := "cgostorage"
	//accessKeyStorage := "OwvEOlzf6e7QwVoV0H75GuSZHpqHxwhYnYL9QbpVPgBRJn+26K26aRJxtZn7Ip5AhaiIkw9kH11xrZSscavXfQ=="
	//dbHost := viper.GetString(`database.host`)
	//dbPort := viper.GetString(`database.port`)
	//dbUser := viper.GetString(`database.user`)
	//dbPass := viper.GetString(`database.pass`)
	//dbName := viper.GetString(`database.name`)
	//dev db

	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "laporbos_db"

	////dev IS
	//baseUrlis := "http://identity-server-asparnas.azurewebsites.net"
	////dev URL Forgot Password
	//urlForgotPassword := "http://cgo-web-api-dev.azurewebsites.net/account/change-password"
	//basicAuth := "cm9jbGllbnQ6c2VjcmV0"
	//redirectUrlGoogle := "http://cgo-web-api.azurewebsites.net/account/callback"
	//clientIDGoogle := "422978617473-acff67dn9cgbomorrbvhqh2i1b6icm84.apps.googleusercontent.com"
	//clientSecretGoogle := "z_XfHM4DtamjRmJdpu8q0gQf"

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	// if err != nil && viper.GetBool("debug") {
	// 	fmt.Println(err)
	// }
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	//middL := middleware.InitMiddleware()
	//e.Use(middL.CORS)
	e.Use(_echoMiddleware.CORSWithConfig(_echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	authorRepo := _authorRepo.NewMysqlAuthorRepository(dbConn)
	ar := _articleRepo.NewMysqlArticleRepository(dbConn)
	masterCOARepo := _masterCOARepo.NewMasterCOARepository(dbConn)
	masterCustomerRepo := _masterCustomerRepo.NewMasterCustomerRepository(dbConn)
	masterVendorRepo := _masterVendorRepo.NewMasterVendorRepository(dbConn)
	salesOrderRepo := _salesOrderRepo.NewSalesOrderRepository(dbConn)

	timeoutContext := 30 * time.Second
	au := _articleUcase.NewArticleUsecase(ar, authorRepo, timeoutContext)
	masterCOAUsecase := _masterCOAUcase.NewMasterCOA(masterCOARepo, timeoutContext)
	masterCustomerUsecase := _masterCustomerUcase.NewMasterCustomer(masterCustomerRepo, timeoutContext)
	masterVendorUsecase := _masterVendorUcase.NewMasterVendor(masterVendorRepo, timeoutContext)
	salesOrderUsecase :=  _salesOrderUcase.NewSalesOrder(salesOrderRepo, timeoutContext)

	_salesOrderHttpDeliver.NewSalesOrderHandler(e,salesOrderUsecase)
	_masterVendorHttpDeliver.NewMasterVendorHandler(e,masterVendorUsecase)
	_masterCustomerHttpDeliver.NewMasterVendorHandler(e,masterCustomerUsecase)
	_masterCOAHttpDeliver.NewMasterCOAHandler(e,masterCOAUsecase)
	_articleHttpDeliver.NewArticleHandler(e, au)
	log.Fatal(e.Start(":9090"))
}
