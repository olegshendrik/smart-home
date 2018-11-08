package server

import (
	"github.com/e154/smart-home/system/swaggo/gin-swagger/swaggerFiles"
	"github.com/gin-gonic/gin"
)

func (s *Server) setControllers() {

	r := s.engine
	basePath := r.Group("/api")

	v1 := basePath.Group("/v1")
	v1.GET("/", s.ControllersV1.Index.Index)
	v1.GET("/swagger", func(context *gin.Context) {
		context.Redirect(302, "/api/v1/swagger/index.html")
	})
	v1.GET("/swagger/*any", s.ControllersV1.Swagger.WrapHandler(swaggerFiles.Handler))

	// nodes
	v1.POST("/node", s.ControllersV1.Node.AddNode)
	v1.GET("/node/:id", s.ControllersV1.Node.GetNodeById)
	v1.PUT("/node/:id", s.ControllersV1.Node.UpdateNode)
	v1.DELETE("/node/:id", s.ControllersV1.Node.DeleteNodeById)
	v1.GET("/node", s.ControllersV1.Node.GetNodeList)

	// scripts
	v1.POST("/script", s.ControllersV1.Script.AddScript)
	v1.GET("/script/:id", s.ControllersV1.Script.GetScriptById)
	v1.PUT("/script/:id", s.ControllersV1.Script.UpdateScript)
	v1.DELETE("/script/:id", s.ControllersV1.Script.DeleteScriptById)
	v1.GET("/script", s.ControllersV1.Script.GetScriptList)
	v1.GET("/script/:id/exec", s.ControllersV1.Script.Exec)
}