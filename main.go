package main

import (
	greet "myPOC/myPackage"
	webstuff "myPOC/myWebPackage"
)

func main() {
	// web server commands:
	// http://localhost:8080/up
	// http://localhost:8080/down
	// http://localhost:8080/status

	// build image and run locally to VS Code:
	// reset everything: docker system prune -a --volumes
	// docker rmi liveness
	// docker build --tag liveness .
	// docker image ls
	// docker run --publish 8080:8080 liveness
	// http://localhost:8080/users/bob

	// publish to docker hub:
	// docker login
	// find image:
	// docker ps
	// push image:
	// docker tag "6d29b5c37447" russellemergentai/liveness_1
	// docker push russellemergentai/liveness_1

	// From the Azure portal:
	// az aks get-credentials --resource-group K8SRG --name aksbetatest
	// kubectl apply -f deploy.yml
	// curl http://20.108.239.131:30475/users/bob
	// kubectl logs web-65f9f75fdf-dccct

	// THE IDEA of the liveness probe is that Kubernetes kills the unhealthy container and starts a new one.

	healthy := true

	greet.Hello()
	webstuff.IDWebcall(&healthy)
}
