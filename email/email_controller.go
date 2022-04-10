package email

import (
	"mail-sender/models"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {

	var email models.Email

	err := c.ShouldBindJSON(&email)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = SendEmail(&email)
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	err = SaveEmail(&email)
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	
	c.JSON(200, gin.H{
		"email": email,
	})

}
