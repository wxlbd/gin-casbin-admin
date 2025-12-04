package ginx

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestParamError(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Execute
	err := errors.New("invalid parameter")
	ParamError(c, err)

	// Verify
	assert.Equal(t, http.StatusBadRequest, w.Code) // The implementation sets HTTP 400
	assert.Contains(t, w.Body.String(), "\"code\":400")
	assert.Contains(t, w.Body.String(), "\"message\":\"invalid parameter\"")
}
