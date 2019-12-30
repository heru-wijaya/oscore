package controller

import (
	"encoding/json"
	"io/ioutil"
	"ovo-score/dto"
	"ovo-score/model"

	"github.com/gin-gonic/gin"
)

// CreateEntityHandler for handler api create entity
func CreateEntityHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, &gin.H{
			"message": err.Error(),
		})
	}

	objEntity := &dto.Entity{}
	err = json.Unmarshal(body, objEntity)

	if err != nil {
		c.JSON(500, &gin.H{
			"message": err.Error(),
		})
	}

	err = model.InsertEntity(c.Request.Context(), objEntity)

	if err != nil {
		c.JSON(500, &gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(201, &gin.H{
		"message": "object created",
	})

	return
}
