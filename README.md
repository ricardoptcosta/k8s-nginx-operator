# k8s-nginx-operator

Requirements:

Minikube v1.22+
Kubectl


This Kubernetes operator creates, updates and deletes Nginx Deployments based on configmaps that are created, updated and deleted.

In order to test have three terminal windows open:
- Window 1 - run operator:
  - $ bash operator.sh

- Window 2 - create, delete or update configmaps:
  - $ kubectl apply -f sample-configmap.yml
  - $ kubectl delete -f sample-configmap.yml

- Window 3 - check deployments being created, deleted or updated:
  - $ kubectl get deployments


To see the text written on the configmap displayed on the browser, do port-forwarding


---

Work in Progress

Convert Bash operator to Golang 
(reading Kubernetes Operators book)



