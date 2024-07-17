package routes

import (
	"SQLTest/config"
	"SQLTest/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// function to get user
func User(env *config.Env, c *gin.Engine) {
	user := c.Group("/user")
	UserModel := env.M
	user.GET("/get-user/:id", func(c *gin.Context) {
		rawID := c.Param("id")
		id, err := strconv.Atoi(rawID)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, err := UserModel.SelectRow(id)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, user)
	})

	user.POST("/create", func(c *gin.Context) {
		var user models.CreateUser
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = UserModel.CreateRow(user)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "User created",
		})
	})
}
