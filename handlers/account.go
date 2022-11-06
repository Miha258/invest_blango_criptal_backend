package handlers

import (
	"invest_blango_criptal_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type newUserPassword struct {
	NewPassword string `json:"password" binding:"required"`
}

type updateBalance struct {
	Sum int64 `json:"sum" binding:"required"`
}

func (h *Handler) changePasssord(c *gin.Context) {
	var input newUserPassword


	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid json parameters")
		return
	}

	userId, _ := c.Get(userCtx)
	userPassword, _ := c.Get(passwordCtx)

	
	if err := h.services.ChangePassword(userId.(int), userPassword.(string)); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": userId,
	})
}


func (h *Handler) getAccountData(c *gin.Context) {
	userLogin := c.Query("login")
	userPassword := c.Query("user_password")

	user, err := h.services.GetUser(models.SingIn{Login: userLogin, Password: userPassword})
	if err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"data": user,		
	})		
}


func (h *Handler) updateAccountDocs(c *gin.Context) {
	var docs models.UserDocs

	if err := c.BindJSON(&docs); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid json parameters")
		return
	}

	userId, _ := c.Get(userCtx)

	if err := h.services.EditUserData(userId.(int), docs); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,		
	})		
}


func (h *Handler) updateAccountBalance(c *gin.Context) {
	var newBalance updateBalance

	if err := c.BindJSON(&newBalance); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid json parameters")
		return
	}

	userId, _ := c.Get(userCtx)
	if err := h.services.UpdateBalance(userId.(int), newBalance.Sum); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,		
	})		
}


func (h *Handler) createPromocode(c *gin.Context) {
	var newPromocode models.UserPromo
	userId, _ := c.Get(userCtx)

	
	if err := c.BindJSON(&newPromocode); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, "Invalid json parameters")
		return
	}

	if err := h.services.CreatePromocode(userId.(int), newPromocode.Name); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,		
	})		
}


func (h *Handler) acceptPromocodeUsage(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	
	if err := h.services.AcceptPromocodeUsage(userId.(int)); err != nil {
		logrus.Error(err)
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,		
	})		
}