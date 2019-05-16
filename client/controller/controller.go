package controller

import (
	"html/template"
	"net/http"

	"github.com/gauravgahlot/dockerdoodle/client/rpc"
	"github.com/gauravgahlot/dockerdoodle/types"
)

var (
	homeController home
	hostController host
)

// Startup registers all the HTTP request handlers
func Startup(templates map[string]*template.Template, client *rpc.Client, hosts *[]types.Host) {
	homeController.homeTemplate = templates["home.html"]
	homeController.hosts = hosts
	homeController.client = client.DockerServiceClient
	homeController.registerRoutes()

	hostController.hostTemplate = templates["host.html"]
	hostController.hosts = hosts
	hostController.client = client.ContainerServiceClient
	hostController.registerRoutes()

	http.Handle("/js/", http.FileServer(http.Dir("client/public")))
	http.Handle("/vendor/", http.FileServer(http.Dir("client/public")))
	http.Handle("/css/", http.FileServer(http.Dir("client/public")))
}
