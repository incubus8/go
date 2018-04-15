package gin

import (
	"github.com/incubus8/go/pkg/errors"
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

func AbortWithAPIError(c *gin.Context, err *errors.APIError) {
	c.Abort()
	err.RecordedAt = time.Now().UTC().Format(time.RFC3339)

	if err.StatusCode != 0 {
		c.JSON(err.StatusCode, err)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}

	if err.Err != nil {
		c.Error(err.Err)
	} else {
		c.Error(err)
	}
}
