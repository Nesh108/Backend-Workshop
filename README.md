# Backend Trends Workshop
### Workshop at Agile Search for Smarp [12.05.2016]
##### Steps
- Create your code (see example: server.go)
- Create Docker file (see example: Dockerfile)
- Push it to Github/Bitbucket
- Create Automated Build on https://hub.docker.com/
- Either push some changes or manually trigger the build
- Check build status
- Pull docker image `docker pull <username>/<name_docker>`
- Set port forwarding
###### Load Balancer with NGINX
- Install nginx
- Setup nginx (see example: LB/nginx.conf)
###### Network of Pods with Kubernetes
- Create service (SVC)
	`kubectl create -f backendWS_service.yaml`
- Create Replication Controller (RC)
	`kubectl create -f backendWS_replicationCtrl.yaml`
- Setup Horizontal Pod Autoscaler (HPA)
	`kubectl autoscale rc backend-ws --min=1 --max=10 --cpu-percent=50`
- Success!