package controllers

import (
	"net/http"

	"web/models"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)

}

func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []models.Person
		result  gin.H
	)

	idb.DB.Find(&persons)

	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)

	id := c.Param("id")

	err := idb.DB.Where("id = ?", id).First(&person).Error

	if err != nil {
		result = gin.H{"result": err.Error()}
	} else {
		idb.DB.Delete(&person)
		result = gin.H{"result": "Successfully deleted person with id " + id}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePerson(c *gin.Context) {
	var (
		person    models.Person
		newPerson models.Person
	)

	id := c.Param("id")

	err := idb.DB.First(&person, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"result": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = idb.DB.Model(&person).Update(newPerson).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "update failed"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"result": "update success"})
	}

}
