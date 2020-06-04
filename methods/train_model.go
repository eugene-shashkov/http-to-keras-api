package methods

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TrainModel is methods where we sending json to our ml python script
// to train and create model *.h5 file
func (m *Method) TrainModel(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	JSONStr := buf.String()
	m.GoPython(`ml.py`, JSONStr)

	var response Response = Response{Status: `model training started`}
	c.JSON(http.StatusOK, response)
}
