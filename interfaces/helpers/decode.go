package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Decode(c *gin.Context, v interface{}) error {
	if err := c.ShouldBindJSON(v); err != nil {
		return fmt.Errorf("unable to decode: %v", err)
	}
	return nil
}
