# Use it with Ingress in Kubernetes

The canary annotation enables the Ingress spec to act as an alternative service for requests to route to depending on the applied rules, and control the traffic splits. 


#### Deploying ingress-nginx Controller


See https://github.com/kubernetes/ingress-nginx/blob/main/docs/deploy/index.md


#### Deploying samples

```shell
kubectl apply -k samples/kubernetes/ingress-nginx-canary/
```

The ingress file is as follows:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: echo-primary
spec:
  rules:
    - host: canary.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: echo-v1
                port:
                  number: 8080
  ingressClassName: nginx
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: echo-canary
  annotations:
    nginx.ingress.kubernetes.io/canary: "true"                               # canary annotation
    nginx.ingress.kubernetes.io/canary-by-header: "Region"                   # canary by header
    nginx.ingress.kubernetes.io/canary-by-header-pattern: "London|Paris"     # canary by header pattern
spec:
  rules:
    - host: canary.example.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: echo-v2
                port:
                  number: 8080
  ingressClassName: nginx
```

#### Verifying the canary

```shell
curl -H "Host: canary.example.com" 127.0.0.1                           # return: v1
curl -H "Host: canary.example.com" -H "Region: Shanghai" 127.0.0.1     # return: v1
curl -H "Host: canary.example.com" -H "Region: Singapore" 127.0.0.1    # return: v1
curl -H "Host: canary.example.com" -H "Region: London" 127.0.0.1       # return: v2
curl -H "Host: canary.example.com" -H "Region: Paris" 127.0.0.1        # return: v2
```

