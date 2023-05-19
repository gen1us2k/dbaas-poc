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

	"github.com/AlekSi/pointer"
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
	return &EverestServer{v: client}, nil
}
func (e *EverestServer) ListKubernetesClusters(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, nil)
}
func (e *EverestServer) RegisterKubernetes(c echo.Context) error {
	var k KubernetesCluster
	if err := c.Bind(&k); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, BadRequest{Message: pointer.ToString(err.Error())})
	}
	log.Println(k.Kubeconfig)
	_, err := clientcmd.BuildConfigFromKubeconfigGetter("", NewConfigGetter(*k.Kubeconfig).loadFromString)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, BadRequest{Message: pointer.ToString(err.Error())})
	}
	m := map[string]interface{}{
		"kubeconfig": k.Kubeconfig,
	}

	_, err = e.v.KVv2("secret").Put(context.TODO(), *k.Name, m)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, BadRequest{Message: pointer.ToString(err.Error())})
	}
	return c.JSON(http.StatusOK, k)
}
func (e *EverestServer) ListDatabases(c echo.Context, kubernetesName string) error {
	return e.proxyKubernetes(c, kubernetesName)
}
func (e *EverestServer) CreateDatabaseCluster(c echo.Context, kubernetesName string) error {
	return c.JSON(http.StatusNotImplemented, nil)
}
func (e *EverestServer) proxyKubernetes(c echo.Context, kubernetesName string) error {
	kConfig, err := e.v.KVv2("secret").Get(context.TODO(), kubernetesName)
	kubeconfig, ok := kConfig.Data["kubeconfig"].(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, BadRequest{Message: pointer.ToString(err.Error())})
	}
	config, err := clientcmd.BuildConfigFromKubeconfigGetter("", NewConfigGetter(kubeconfig).loadFromString)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, BadRequest{Message: pointer.ToString(err.Error())})
	}
	data, err := json.Marshal(kConfig)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, BadRequest{Message: pointer.ToString(err.Error())})
	}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, BadRequest{Message: pointer.ToString(err.Error())})
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Host:   strings.TrimPrefix(config.Host, "https://"),
		Scheme: "https",
	})
	transport, err := rest.TransportFor(config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, BadRequest{Message: pointer.ToString(err.Error())})
	}
	reverseProxy.Transport = transport
	req := c.Request()
	req.URL.Path = fmt.Sprintf("/apis/dbaas.percona.com/v1/namespaces/%s/databaseclusters", "default")
	reverseProxy.ServeHTTP(c.Response(), req)
	return nil
}
