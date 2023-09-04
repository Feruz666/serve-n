package controllers

import (
	"log"
	"net/http"

	account "github.com/Feruz666/serve-n/gateway/protos/account/protos"
	"github.com/gin-gonic/gin"
)

func (cn *Controllers) Register(ctx *gin.Context) {
	var user account.Person

	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Fatalf("ctx.ShouldBind() failed %e", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	err := cn.client.CreatePerson(ctx.Request.Context(), &user)
	if err != nil {
		log.Fatalf("сreatePerson() failed %e", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "CreatePerson()",
		})
	}


	

}
