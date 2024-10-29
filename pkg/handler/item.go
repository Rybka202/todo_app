package handler

import (
	"net/http"
	"strconv"

	"todo_app"

	"github.com/gin-gonic/gin"
)

//	@Summary		Create Todo Item
//	@Security		ApiKeyAuth
//	@Tags			items
//	@Description	create todo item
//	@ID				create-item
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Account ID"
//	@Param			input	body		todo.TodoItem	true	"item info"
//	@Success		200		{integer}	integer			1
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context){
	userId, err := getUserID(c)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoItem.Create(userId, listId, input)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

type getAllItemsResponse struct{
	Data []todo.TodoItem `json:"data"`
}

//	@Summary		Get All Items 
//	@Security		ApiKeyAuth
//	@Tags			items
//	@Description	get all items
//	@ID				get-all-items
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"Account ID"
//	@Success		200		{object}	getAllItemsResponse
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/api/lists/{id}/items [get]
func (h *Handler) getAllItems(c *gin.Context){
	userId, err := getUserID(c)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.service.TodoItem.GetAll(userId, listId)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

//	@Summary		Get Item By Id
//	@Security		ApiKeyAuth
//	@Tags			items
//	@Description	get item by id
//	@ID				get-item-by-id
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"Account ID"
//	@Success		200		{object}	todo.TodoItem
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context){
	userId, err := getUserID(c)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	item, err := h.service.TodoItem.GetById(userId, id)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

//	@Summary		Update Item
//	@Security		ApiKeyAuth
//	@Tags			items
//	@Description	update item
//	@ID				update_item
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Account ID"
//	@Param			input	body		todo.UpdateItemInput	true	"update item info"
//	@Success		200		{object}	StatusResp
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context){
	userId, err := getUserID(c)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, "no update values")
		return
	}

	if err := h.service.TodoItem.Update(userId, id, input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResp{
		Status: "ok",
	})
}

//	@Summary		Delete Item
//	@Security		ApiKeyAuth
//	@Tags			items
//	@Description	delete item
//	@ID				delete_item
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"Account ID"
//	@Success		200		{object}	StatusResp
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context){
	userId, err := getUserID(c)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.service.TodoItem.Delete(userId, id); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResp{
		Status: "ok",
	})
}