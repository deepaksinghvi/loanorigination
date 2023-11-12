/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"github.com/deepaksinghvi/loanorigination/common"
	"github.com/deepaksinghvi/loanorigination/controller"
	_ "github.com/deepaksinghvi/loanorigination/docs"
	"github.com/deepaksinghvi/loanorigination/loan_worker"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main1() {
	/*loan_worker.StartWorker()
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)
	fmt.Println("Loan Origination loan_worker started, press ctrl+c to terminate...")
	<-done */
}

func main() {
	var h common.CadenceHelper
	h.SetupServiceConfig()
	common.NewLOCadenceHelper(&h)
	loan_worker.StartWorker(&h)
	fmt.Println("Loan Origination loan_worker started, press ctrl+c to terminate...")
	router := setupRouter()
	router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	// Swagger setup
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/loan-application", controller.CreateLoanApplication)
	router.GET("/loan-application/:workflow_id/:run_id", controller.GetLoanApplication)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	return router
}
