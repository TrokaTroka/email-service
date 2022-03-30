package email

import (
	"github.com/gin-gonic/gin"
	"mail-sender/models"
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

	err = SaveEmail(&email)
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	err = SendEmail(&email)
	
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"email": email,
	})

}
