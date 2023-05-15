package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	vault "github.com/hashicorp/vault/api"
	"github.com/labstack/echo/v4"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type EverestServer struct {
	v *vault.Client
}

func NewEverestServer() (*EverestServer, error) {
	config := vault.DefaultConfig()
	config.Address = "http://127.0.0.1:4321"
	client, err := vault.NewClient(config)
	if err != nil {
		return nil, err
	}
	client.SetToken("myroot")
	return &EverestServer{v: client}
}
func (e *EverestServer) ListKubernetesClusters(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}
func (e *EverestServer) RegisterKubernetes(c echo.Context) error {
	var k KubeCluster
	if err := c.BindJSON(&k); err != nil {
		log.Println(err)
		return
	}
	log.Println(k.Kubeconfig)
	_, err := clientcmd.BuildConfigFromKubeconfigGetter("", NewConfigGetter(k.Kubeconfig).loadFromString)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	m := map[string]interface{}{
		"kubeconfig": k.Kubeconfig,
	}

	_, err = s.v.KVv2("secret").Put(context.TODO(), k.Name, m)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	return c.JSON(http.StatusNotImplemented, nil)
}
func (e *EverestServer) ListDatabases(c echo.Context, kubernetesName string) error {
	return c.JSON(http.StatusNotImplemented, nil)
}
func (e *EverestServer) CreateDatabaseCluster(c echo.Context, kubernetesName string) error {
	return c.JSON(http.StatusNotImplemented, nil)
}
func (e *EverestServer) proxyKubernetes(c echo.Context, kubernetesName string) error {
	kConfig, err := s.v.KVv2("secret").Get(context.TODO(), kubernetesName)
	kubeconfig, ok := kConfig.Data["kubeconfig"].(string)
	if !ok {
		return
	}
	config, err := clientcmd.BuildConfigFromKubeconfigGetter("", NewConfigGetter(kubeconfig).loadFromString)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	data, err := json.Marshal(kConfig)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Host:   strings.TrimPrefix(config.Host, "https://"),
		Scheme: "https",
	})
	transport, err := rest.TransportFor(config)
	if err != nil {
		return
	}
	reverseProxy.Transport = transport
	req := c.Request
	req.URL.Path = strings.TrimLeft(req.URL.Path, fmt.Sprintf("/proxy/%s", name))
	reverseProxy.ServeHTTP(c.Writer, req)
}
