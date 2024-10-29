package handler

import (
	"net/http"
	"todo_app"

	"github.com/gin-gonic/gin"
)

//	@Summary		SignUp
//	@Tags			auth
//	@Description	create account
//	@ID				create-account
//	@Accept			json
//	@Produce		json
//	@Param			input	body		todo.User	true	"account info"
//	@Success		200		{integer}	integer		1
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context){
	var input todo.User
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, "error invalid input data")
		return
	}

	id, err := h.service.Authorization.CreateUser(input)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]any{
		"id": id,
	})
}

type signInInput struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


//	@Summary		SignIn
//	@Tags			auth
//	@Description	login
//	@ID				login
//	@Accept			json
//	@Produce		json
//	@Param			input	body		signInInput	true	"credentials"
//	@Success		200		{string}	string		"token"
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context){
	var input signInInput
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}
