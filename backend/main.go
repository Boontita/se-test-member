package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Boontita/se-64-example/controller"
	"github.com/Boontita/se-64-example/entity"
	"github.com/Boontita/se-64-example/middlewares"

)

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{

			// BookOrder Routes
			protected.GET("/book_orders", controller.ListBookOrders)
			protected.GET("/book_order/:id", controller.GetBookOrder)
			protected.POST("/book_orders", controller.CreateBookOrder)
			protected.PATCH("/book_orders", controller.UpdateBookOrder)
			protected.DELETE("/book_orders/:id", controller.DeleteBookOrder)

			// BookType Routes
			protected.GET("/book_types", controller.ListBookTypes)
			protected.GET("/book_type/:id", controller.GetBookType)
			protected.POST("/book_types", controller.CreateBookType)
			protected.PATCH("/book_types", controller.UpdateBookType)
			protected.DELETE("/book_types/:id", controller.DeleteBookType)

			// Company Routes
			protected.GET("/companies", controller.ListCompanies)
			protected.GET("/company/:id", controller.GetCompany)
			protected.POST("/compamies", controller.CreateCompany)
			protected.PATCH("/companies", controller.UpdateCompany)
			protected.DELETE("/companies/:id", controller.DeleteCompany)

			//OrderStatus Routes
			protected.GET("/order_statuses", controller.ListOrderStatuses)
			protected.GET("/order_status/:id", controller.GetOrderStatus)
			protected.POST("/order_statuses", controller.CreateOrderStatus)
			protected.PATCH("/order_statuses", controller.UpdateOrderStatus)
			protected.DELETE("/order_statuses/:id", controller.DeleteOrderStatus)
		}
	}

	// Member Routes
	r.POST("/members", controller.CreateMember)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
