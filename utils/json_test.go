package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/* -------------------------------------------------------------------------- */
/*                             Return JSON Struct                             */
/* -------------------------------------------------------------------------- */
func TestReturnJSONStruct(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	sampleStruct := struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}{
		Name:  "test",
		Count: 123,
	}
	assert.NotNil(t,sampleStruct)
	ReturnJSONStruct(c, &sampleStruct)
}
