# Use it with Istio in Kubernetes

To run the server with Istio in Kubernetes:

```yaml
kubectl apply -f samples/istio/http-echo/echo-v1-deployment.yaml
kubectl apply -f samples/istio/http-echo/echo-v2-deployment.yaml
kubectl apply -f samples/istio/http-echo/echo-service.yaml

kubectl apply -f samples/istio/http-echo/istio-gateway.yaml
kubectl apply -f samples/istio/http-echo/istio-virtual-service.yaml
kubectl apply -f samples/istio/http-echo/istio-destination-rule.yaml
```

Then visit http://localhost in your browser.