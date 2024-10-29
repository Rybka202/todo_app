package handler

import (
	"net/http"
	"strconv"
	"todo_app"

	"github.com/gin-gonic/gin"
)

//	@Summary		Create Todo List
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	create todo list
//	@ID				create-list
//	@Accept			json
//	@Produce		json
//	@Param			input	body		todo.TodoList	true	"list info"
//	@Success		200		{integer}	integer			1
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/api/lists [post]
func (h *Handler) createList(c *gin.Context){
	userId, err := getUserID(c)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoList.Create(userId, input)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

type getAllListsResponse struct{
	Data []todo.TodoList `json:"data"`
}

//	@Summary		Get All Lists 
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	get all lists
//	@ID				get-all-lists
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	getAllListsResponse
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/api/lists [get]
func (h *Handler) getAllLists(c *gin.Context){
	userId, err := getUserID(c)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.service.TodoList.GetAll(userId)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}


//	@Summary		Get List By Id
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	get list by id
//	@ID				get-list-by-id
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"Account ID"
//	@Success		200		{object}	todo.TodoList
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context){
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

	list, err := h.service.TodoList.GetById(userId, id)
	if err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

//	@Summary		Update List
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	update list
//	@ID				update_list
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Account ID"
//	@Param			input	body		todo.UpdateListInput	true	"update list info"
//	@Success		200		{object}	StatusResp
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context){
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

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.TodoList.Update(userId, id, input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResp{
		Status: "ok",
	})
}


//	@Summary		Delete List
//	@Security		ApiKeyAuth
//	@Tags			lists
//	@Description	delete list
//	@ID				delete_list
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int	true	"Account ID"
//	@Success		200		{object}	StatusResp
//	@Failure		400,404	{object}	ErrorResp
//	@Failure		500		{object}	ErrorResp
//	@Failure		default	{object}	ErrorResp
//	@Router			/api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context){
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

	if err := h.service.TodoList.Delete(userId, id); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResp{
		Status: "ok",
	})
}