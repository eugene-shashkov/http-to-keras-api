package methods

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Method) GetStatus(c *gin.Context) {

	var response Response = Response{Status: `ok`}
	c.JSON(http.StatusOK, response)
}
