
kubectl create -f kubernetes/my-sql-secret.yml
kubectl apply -f kubernetes/my-sql-pv.yml
kubectl apply -f kubernetes/my-sql-pvc.yml
kubectl apply -f kubernetes/my-sql-deployment.yml
kubectl apply -f kubernetes/my-sql-service.yml 

kubectl apply -f kubernetes/rabbitmq-deployment.yml
kubectl apply -f kubernetes/rabbitmq-service.yml 

kubectl apply -f cache-app.yml
kubectl apply -f cache-app-service.yml
