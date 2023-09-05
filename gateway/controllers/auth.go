package controllers

import (
	"github.com/hashicorp/go-hclog"
	"net/http"

	account "github.com/Feruz666/serve-n/gateway/protos/account/protos"
	"github.com/gin-gonic/gin"
)

func (cn *Controllers) Register(ctx *gin.Context) {
	var user account.Person
	log := hclog.Default()

	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Error("[controllers.Register] ctx.ShouldBindJSON", "error", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})

		return
	}

	err := cn.client.CreatePerson(ctx.Request.Context(), &user)
	if err != nil {
		log.Error("[controllers.Register] cn.client.CreatePerson", "error", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "CreatePerson()",
		})
	}

	// TODO GET RESPONCE FROM KAFKA ...

}
