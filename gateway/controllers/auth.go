package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cn *Controllers) Register(ctx *gin.Context) {

	

	err := cn.client.CreatePerson()
	if err != nil {
		log.Fatalf("сreatePerson() failed %w", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "CreatePerson()",
		})
	}
}
