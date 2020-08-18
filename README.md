# cache
--------------------------------------------
Cache is the microservice deployed on kubernetes. It stores data in memory and has a persistent backup from data base.
- The microservice provides functionality of get and post to cache memory. The memory make sure that the same is persisted in mysql database.
- On redeployment of service the rabbitmq fetches all data from db and inserts it into cache memory.

API exposed -
- GET - /cache
- POST - /cache
- GET - /swagger 
---------------------------------------------

Steps to run it locally - 

kubectl create -f kubernetes/my-sql-secret.yml
kubectl apply -f kubernetes/my-sql-pv.yml
kubectl apply -f kubernetes/my-sql-pvc.yml
kubectl apply -f kubernetes/my-sql-deployment.yml
kubectl apply -f kubernetes/my-sql-service.yml 

kubectl apply -f kubernetes/rabbitmq-deployment.yml
kubectl apply -f kubernetes/rabbitmq-service.yml 

kubectl apply -f cache-app.yml
kubectl apply -f cache-app-service.yml

Since we are running it on minikube we can get the app url as- 
 minikube service cache-app --url 
 We can use above url to get the results from APIs.
 

