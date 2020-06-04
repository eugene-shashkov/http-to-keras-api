package methods

import "github.com/gin-gonic/gin"

type Interface interface {
	IsModelExist(c *gin.Context)
	TrainModel(c *gin.Context)
	IsPathExist(path string) bool
	GoPython(file string, j string) string
	GetStatus(c *gin.Context)
}

type Method struct{}

func InitMethods() Interface {
	return &Method{}
}
