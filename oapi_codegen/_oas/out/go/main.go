package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// CreateDatabaseCluster - Create a database cluster on given kubernetes cluster
	e.POST("/v1/kubernetes/:kubernetesName/database-cluster", c.CreateDatabaseCluster)

	// ListDatabases - List of created database clusters on provided kubernetes cluster
	e.GET("/v1/kubernetes/:kubernetesName/database-cluster", c.ListDatabases)

	// ListKubernetesClusters - List of registered kubernetes clusters
	e.GET("/v1/kubernetes", c.ListKubernetesClusters)

	// RegisterKubernetes - Register Kubernetes cluster in the control plane
	e.POST("/v1/kubernetes", c.RegisterKubernetes)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}