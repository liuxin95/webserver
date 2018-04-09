package main

import (
	"encoding/json"
	"github.com/emicklei/go-restful"
	"io"
	"log"
	"net/http"
	"github.com/woqutech/k8s"
)

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/deploymentinfo").To(deploymentinfo))
	ws.Route(ws.GET("/deploymentcreate").To(deploymentcreate))
	ws.Route(ws.GET("/deploymentdelete").To(deploymentdelete))
	ws.Route(ws.GET("/deploymentupdate").To(deploymentupdate))
	ws.Route(ws.GET("/deploymentlist").To(deploymentlist))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func deploymentinfo(req *restful.Request, resp *restful.Response) {
	deployment := k8s.DeployMentInfo()
	tmp, _ := json.Marshal(*deployment)
	io.WriteString(resp, string(tmp))
}
func deploymentcreate(req *restful.Request, resp *restful.Response) {
	k8s.Create()
	io.WriteString(resp, "ok")
}
func deploymentdelete(req *restful.Request, resp *restful.Response) {
	k8s.Delete()
	io.WriteString(resp, "ok")
}
func deploymentupdate(req *restful.Request, resp *restful.Response) {
	k8s.Update()
	io.WriteString(resp, "ok")
}
func deploymentlist(req *restful.Request, resp *restful.Response) {
	deployments := k8s.List()
	tmp, _ := json.Marshal(deployments)
	io.WriteString(resp, string(tmp))
}



