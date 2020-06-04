package methods

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IsModelExist method gives possibility to check from client side
// event when model is created
func (m *Method) IsModelExist(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	strJSON := buf.String()

	type Request struct {
		Model int `json:"modelid"`
	}

	var request Request
	json.Unmarshal([]byte(strJSON), &request)
	modelid := strconv.Itoa(request.Model)

	type Response struct {
		ModelStatus int `json:"status"`
	}

	var response Response

	response.ModelStatus = 0

	if m.IsPathExist(`ml/models/` + modelid) {
		response.ModelStatus = 1
	}

	c.JSON(http.StatusOK, response)
}
