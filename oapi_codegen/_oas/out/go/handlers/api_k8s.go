package handlers

import (
	"net/http"

	"github.com/GIT_USER_ID/GIT_REPO_ID/models"
	"github.com/labstack/echo/v4"
)

// CreateDatabaseCluster - Create a database cluster on given kubernetes cluster
func (c *Container) CreateDatabaseCluster(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// ListDatabases - List of created database clusters on provided kubernetes cluster
func (c *Container) ListDatabases(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// ListKubernetesClusters - List of registered kubernetes clusters
func (c *Container) ListKubernetesClusters(ctx echo.Context) error {
	// The line should present here
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}

// RegisterKubernetes - Register Kubernetes cluster in the control plane
func (c *Container) RegisterKubernetes(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.HelloWorld{
		Message: "Hello World",
	})
}
