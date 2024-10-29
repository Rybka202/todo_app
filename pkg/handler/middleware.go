package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const(
	authorizationHeader = "Authorization"
	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context){
	header := c.GetHeader(authorizationHeader)
	if header == ""{
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2{
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[0] != "Bearer"{
		newErrorResponse(c, http.StatusUnauthorized, "invalid bearer header")
		return
	}

	userID, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil{
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)
}

func getUserID(c *gin.Context) (int, error){
	id, ok := c.Get(userCtx)
	if !ok{
		return 0, errors.New("user is not found")
	}

	idInt, ok := id.(int)
	if !ok{
		return 0, errors.New("user id is not valid type")
	}

	return idInt, nil
}