package http

import (
	"fmt"
	"net/http"

	"github.com/LIYINGZHEN/ginexample/internal/app/types"
	"github.com/gin-gonic/gin"
)

func (a *AppServer) GetLinkHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	link, err := a.LinkService.GetLink(id)
	if err != nil {
		a.Logger.Printf("error getting link %v", err)
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"UserName": link.UserName,
		"Url":      link.Url,
	})

}

func (a *AppServer) UpdateLinkHandler(c *gin.Context) {
	type request struct {
		UserName string `form:"username" json:"username" binding:"required"`
		Url      string `form:"url" json:"url" binding:"required"`
	}
	var (
		req request
	)

	id := c.Param("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	link, err := a.LinkService.GetLink(id)
	if err != nil {
		a.Logger.Printf("error getting link %v", err)
		c.Status(http.StatusNotFound)
		return
	}

	err = c.Bind(&req)
	if err != nil || req.UserName == "" || req.Url == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	link.UserName = req.UserName
	link.Url = req.Url
	err = a.LinkService.UpdateLink(link)
	if err != nil {
		a.Logger.Printf("error updatinging link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"UserName": link.UserName,
		"Url":      link.Url,
	})

}

func (a *AppServer) CreateLinkHandler(c *gin.Context) {
	type request struct {
		UserName string `form:"username" json:"username" binding:"required"`
		Url      string `form:"url" json:"url" binding:"required"`
	}
	var (
		req       request
		linkModel types.Link
	)

	err := c.Bind(&req)
	if err != nil || req.UserName == "" || req.Url == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	linkModel.UserName = req.UserName
	//linkModel.Url = req.Url

	link, err := a.LinkService.CreateLink(&linkModel, req.Url)
	if err != nil {
		a.Logger.Printf("error creating link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"UserName": link.UserName,
		"Url":      link.Url,
	})
}

func (a *AppServer) DeleteLinkHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	err := a.LinkService.DeleteLink(id)
	if err != nil {
		a.Logger.Printf("error deleting link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Id #" + id: "deleted",
	})
}

func (a *AppServer) GetAllLinkHandler(c *gin.Context) {
	type response struct {
		ID       string `form:"id" json:"id" binding:"required"`
		UserName string `form:"username" json:"username" binding:"required"`
		Url      string `form:"url" json:"url" binding:"required"`
	}

	links, err := a.LinkService.GetAllLink()
	if err != nil {
		a.Logger.Printf("error get all link: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var res = make([]response, len(links))
	for k, v := range links {
		res[k] = response{ID: fmt.Sprintf("%v", v.ID), UserName: v.UserName, Url: v.Url}
	}
	c.JSON(http.StatusOK, res)

}
