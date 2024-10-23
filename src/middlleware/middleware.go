package middlleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthRequired(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.Split(auth, "Bearer ")[1]
	fmt.Println(token)
	return
}
