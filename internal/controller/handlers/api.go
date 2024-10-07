package handlers

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setUpApi(h *Handler) {
	h.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//h.engine.Use(token.JWTMiddleware())

	pureWash := h.engine.Group("/api")
	{
		pureWash.GET("/ping", h.Ping)
		pureWash.POST("/company", h.CreateCompanyHandler)
		pureWash.GET("/company/:id", h.GetCompanyHandler)
		pureWash.PUT("/company/:id", h.UpdateCompanyHandler)
		pureWash.DELETE("/company/:id", h.DeleteCompanyHandler)

		pureWash.POST("/order", h.CreateOrderHandler)
		pureWash.GET("/order/:id", h.GetOrderHandler)
		pureWash.GET("/orders", h.GetAllOrders)
		pureWash.GET("/courier_orders", h.GetAllForCourierOrders)
		pureWash.PUT("/order/:id", h.UpdateOrderHandler)
		pureWash.DELETE("/order/:id", h.DeleteOrderHandler)

		pureWash.POST("/address", h.CreateAddressHandler)
		pureWash.GET("/address/:id", h.GetAddressHandler)
		pureWash.PUT("/address/:id", h.UpdateAddressHandler)
		pureWash.DELETE("/address/:id", h.DeleteAddressHandler)

		pureWash.POST("/service", h.CreateServiceHandler)
		pureWash.GET("/service/:id", h.GetServiceHandler)
		pureWash.GET("/services", h.GetAllServices)
		pureWash.PUT("/service/:id", h.UpdateServiceHandler)
		pureWash.DELETE("/service/:id", h.DeleteServiceHandler)

		// pureWash.POST("/user_order", h.CreateOrderForUserHandler)
		// pureWash.GET("/user_order/:id", h.GetOrderForUserHandler)
		// pureWash.GET("/user_orders", h.GetAllOrdersForUser)
		// pureWash.PUT("/user_order_canceled/:id", h.UpdateOrderForUserHandler)

		// pureWash.POST("/user_address", h.CreateAddressForUserHandler)
		// pureWash.GET("/user_address/:id", h.GetAddressForUserHandler)
		// pureWash.PUT("/user_address/:id", h.UpdateAddressForUserHandler)

	}

}
