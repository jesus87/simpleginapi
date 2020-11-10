package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/banwire/api-exam/models"
	"github.com/banwire/api-exam/store"
	"github.com/gin-gonic/gin"
)

var merchantRepository store.MerchantRepository

func CreateMerchantHandler(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%v - %v", "bytes", err))
		return
	}

	var merchant models.MerchantRequest
	err = json.Unmarshal(bytes, &merchant)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%v - %v", "serialized", err))
		return
	}

	err = merchantRepository.Create(merchant.MapCreate())
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v - %v", "connection", err))
		return
	}

	c.JSON(http.StatusOK, "")
}

func GetMerchantHandler(c *gin.Context) {
	values, err := merchantRepository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v - %v", "connection", err))
		return
	}

	c.JSON(http.StatusOK, values)
}

func EditMerchantHandler(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%v - %v", "bytes", err))
		return
	}

	var merchant models.MerchantRequest
	err = json.Unmarshal(bytes, &merchant)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("%v - %v", "serialized", err))
		return
	}

	err = merchantRepository.Edit(merchant.MapUpdate())
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("%v - %v", "connection", err))
		return
	}

	c.JSON(http.StatusOK, "")
}
