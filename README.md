# klustered-pancakes

## build
docker build -t peterjbishop/klustered-pancakes:v1 .

## push 
docker push peterjbishop/klustered-pancakes:v1

## deploy in minikube
minikube start
kubectl apply -f deployment.yaml

## redeploy in minikube 
kubectl rollout restart deployment/klustered-pancakes-deployment

## tunnel into the cluster
minikube service klustered-pancakes-app-service

## load local docker image into minikube
minikube image load peterjbishop/klustered-pancakes:v1

## update { reminder: change version! }
- docker build -t peterjbishop/klustered-pancakes:v1 .
- minikube image load peterjbishop/klustered-pancakes:v1
- kubectl rollout restart deployment/klustered-pancakes-deployment


