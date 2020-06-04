package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/fonjeekay/gameai-py/methods"
)

func main() {

	r := gin.Default()

	m := methods.InitMethods()

	r.POST("/trainmodel", m.TrainModel)

	r.POST("/ismodelexist", m.IsModelExist)

	r.GET("/prediction", m.TrainModel)

	r.GET("/status", m.GetStatus)

	r.Run(":64128")
}
