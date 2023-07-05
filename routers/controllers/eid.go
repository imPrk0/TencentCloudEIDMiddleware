package controllers

import (
	model "TencentCloudEIDMiddleware/models"
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm4"
	"github.com/tjfoc/gmsm/x509"
	"net/http"
)

func EidDecode(c *gin.Context) {
	privateKey := c.Query("key")
	DesKey := c.Query("des_key")
	UserInfo := c.Query("user_info")

	pb, err := x509.ReadPrivateKeyFromHex(privateKey)
	if nil != err {
		c.JSON(http.StatusInternalServerError, model.ResJSON{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	dek, err := base64.StdEncoding.DecodeString(DesKey)
	if nil != err {
		c.JSON(http.StatusInternalServerError, model.ResJSON{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	sessionKey, err := sm2.Decrypt(pb, dek, 0)
	if nil != err {
		c.JSON(http.StatusInternalServerError, model.ResJSON{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	info, _ := base64.StdEncoding.DecodeString(UserInfo)
	out, err := sm4.Sm4Ecb(sessionKey, info, false)
	if nil != err {
		c.JSON(http.StatusInternalServerError, model.ResJSON{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	var userInfo *model.EidData
	err = json.Unmarshal([]byte(string(out)), &userInfo)
	if nil != err {
		c.JSON(http.StatusInternalServerError, model.ResJSON{
			Code:    500,
			Message: "Failed to parse user JSON",
		})
		return
	}

	c.JSON(http.StatusOK, model.ResJSON{
		Code:    0,
		Message: "success",
		Data:    userInfo,
	})
	return
}
