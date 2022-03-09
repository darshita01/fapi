package handlers

import (
	"backend/business/services/user"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func middleware1(m echo.HandlerFunc) echo.HandlerFunc{
	// fmt.Println("middlerware1")
	// return func(ctx echo.Context) error {
	// 	fmt.Println("return of middleware1")
	// 	ctx.Response().Header().Set("context-type","middleware1")
	// 	ctx.Response().After(func() {
	// 		fmt.Println("After response in middleware1")
	// 	})
	// 	return m(ctx)
	// }

	log.Println("inside middleware 1 ")

	return m
}

// func middleware2(m echo.HandlerFunc) echo.HandlerFunc{
// 	fmt.Println("middleware2")
// 	return func(ctx echo.Context) error {
// 		fmt.Println("return of middleware2")
// 		ctx.Response().After(func() {
// 			fmt.Println("After response in middleware2")
// 		})
// 		return m(ctx)
// 	}
// }
// func response(m echo.HandlerFunc) echo.HandlerFunc{
// 	fmt.Println("Inside response middleware")
// 	return func(ctx echo.Context) error {
// 		fmt.Println("Setting header")
// 		ctx.Response().Header().Set("context-type","asdfghj")
// 		ctx.Response().After(func() {
// 			fmt.Println("After response in response")
// 		})
// 		return m(ctx)
// 	}
// }

func Api(db *sqlx.DB , e *echo.Echo){

	h := userHandler{
		user: user.New(db),
	}

// 	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
//         Format: `[${time_rfc3339}]  ${method}  ${status}  ${latency_human} $` + "\n",
// }))

	// fileXYZ, err := os.OpenFile("test.log", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666 )
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// defer fileXYZ.Close()

	// log.SetOutput(fileXYZ)

	// log.Println("First line in the file....")

	
	e.GET("/Users/:id",h.GetData, middleware1) // ,middleware1,middleware2, response
	e.GET("/Users/",h.GetData)
	e.POST("/Users/",h.EnterData)
	e.PUT("/Users/:id",h.EditData)
	e.DELETE("/Users/:id",h.DeleteData)
	
	
}

