### K8s Deployment Guide
1. Create a configmap that will be mounted on the pod to inject environment variables.
2. ```k create cm go-starter-stage-config --from-file=./.env.stage```. .env.stage file must exist before command execution, and must contain the environment variables that need to be injected.
3. Following commands should be executed from k8s/branch. Refer individual yml files to see resource specifications.
3. Create a deployment ```k create -f deploy.yml```.
4. Create a service to expose deployment externally - ```k create -f svc.yml```
5. Create a horizontal pod autoscaler - ```k create -f hpa.yml```