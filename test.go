package main

import (
	"fmt"
	"gin-data/app/models"
	"gin-data/app/services"
	"gin-data/utils/setting"
)

func init() {
	setting.Setup()
	models.Setup()
}

func maini() {
	result,_ :=services.GetTempFlagDistribution()
	fmt.Println(result)
}
