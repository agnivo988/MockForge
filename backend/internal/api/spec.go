package api

import (

	"github.com/gin-gonic/gin"

	"github.com/agnivo988/MockForge/backend/internal/parser"
	"github.com/agnivo988/MockForge/backend/internal/store"
)

func UploadSpec(c *gin.Context) {
	env := c.PostForm("env")
	file, err := c.FormFile("spec")
	if err != nil {
		c.JSON(400, gin.H{"error": "file required"})
		return
	}

	path := "specs/" + env + ".yaml"
	_ = c.SaveUploadedFile(file, path)

	spec, err := parser.LoadSpec(path)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	store.Manager.Set(env, spec)
	c.JSON(200, gin.H{"status": "uploaded"})
}
