package routes

import (
	"SQLTest/persistence"
	"strconv"

	"github.com/gin-gonic/gin"
)

// function to get user
func User(c *gin.Engine) {

	user := c.Group("/user")

	user.GET("/:id", func(c *gin.Context) {
		rawID := c.Param("id")
		id, err := strconv.Atoi(rawID)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, err := persistence.SelectRow(id)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, user)
	})
}
