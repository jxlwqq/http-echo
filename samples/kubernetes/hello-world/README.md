# Use it in Kubernetes

To run the server in Kubernetes:

```yaml
kubectl apply -f samples/kubernetes/hello-world/deployment.yaml
kubectl apply -f samples/kubernetes/hello-world/service.yaml
kubectl port-forward services/http-echo 8080:8080
```

Then visit http://localhost:8080/ in your browser.