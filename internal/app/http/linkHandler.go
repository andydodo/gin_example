package http

import (
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
		"Name": link.UserName,
		"Url":  link.Url,
	})

}

func (a *AppServer) UpdateLinkHandler(c *gin.Context) {
	type request struct {
		UserName string `json:"username"`
		Url      string `json:"url"`
	}

	var (
		req  request
		link types.Link
	)

	id := c.Param("id")
	if id == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	_, err := a.LinkService.GetLink(id)
	if err != nil {
		a.Logger.Printf("error getting link %v", err)
		c.Status(http.StatusNotFound)
		return
	}

	err = c.BindJSON(&req)
	if err != nil || req.UserName == "" || req.Url == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	link.UserName = req.UserName
	link.Url = req.Url

	err = a.LinkService.UpdateLink(&link)
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
		UserName string `json:"username"`
		Url      string `json:"url"`
	}

	var (
		linkModel types.Link
		req       request
	)

	err := c.BindJSON(&req)
	if err != nil || req.UserName == "" || req.Url == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	linkModel.UserName = req.UserName

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

	link, err := a.LinkService.DeleteLink(id)
	if err != nil {
		a.Logger.Printf("error deleting link: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"UserName": link.UserName,
		"Url":      link.Url,
	})
}
