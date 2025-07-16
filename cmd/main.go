package main

import (
	// "fmt"

	"github.com/labstack/echo/v4"
	"github.com/takumifahri/RESTful-API-GO/internal/database"
	strRepo "github.com/takumifahri/RESTful-API-GO/internal/repository/catalog"
	strUsecase"github.com/takumifahri/RESTful-API-GO/internal/usecase/store"
	"github.com/takumifahri/RESTful-API-GO/internal/delivery/rest"

)

// Method kosongan untuk enum
// type TypeFood uint8
// const (
// 	APPETIZER TypeFood = iota
// 	MAIN_COURSE
// 	DESSERT
// 	BEVERAGE
// )

// type TypeDirnk uint8
// const (
// 	ALCOHOLIC TypeDirnk = iota
// 	NON_ALCOHOLIC
// 	COFFEE
// 	TEA
// )
const (
 	dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
)



func main() {
	e := echo.New()

	db := database.ConnectDB(dbAddress)
	
	catalogRepo := strRepo.GetRepository(db)
	
	storeUsecase := strUsecase.GetUsecase(catalogRepo)
	
	handler := rest.NewHandler(storeUsecase)

	rest.LoadRoutes(e, handler)
	//Loger uttk port nya
	e.Logger.Fatal(e.Start(":8081"))
}
