# k8s-nginx-operator

This repo is a step by step approach to creating a kubernetes operator following a natural self discovery path. The end goal here is to create a bad ass operator!

## Requirements:

Minikube v1.22+  
Kubectl


## Attempt 1 

We heard about 


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


# Attempt 2

Using an existing kubernetes resources like a Deployment is cool, but what I would really like to do is to create my own custom resource. What's simple and useful?

I love plants. I have many plants. I forget to water them when I should. They die. I get sad. I need to address this so that I can get my life in order. Could a Kubernetes CustomResourceDefinition come to the rescue?

- Create CRD and deploy it
  1. Create CRD wateringalarm-crd.yml
  2. Run `$ kubectl apply -f attempt-2/wateringalarm-crd.yml`

Boom! We have a new CRD. Check $ kubectl api-resources to double check we have `wateringalarms` as a new api resource

- Create custom resources of type WateringAlarm. In this case we will build an alarm for orchids and succulents
  1. Create orchid-wateringalarm.yml and succulent-wateringalarm.yml
  2. Run  `$ kubectl apply -f attemp-2/orchid-wateringalarm.yml && kubectl apply -f attempt2/succulent-wateringalarm.yml`

Now, I have the resources created, how do I take advantage of the operator pattern to leverage them?
- Create a watcher on the operator
- Do something when the the watcher identifies a change

This is exactly the same as on attempt-1, now with my beautiful custom resource. See attempt-2/operator.sh for my implementation.


# Attempt 3

Bash is nice, but I hear Kubernetes is written in Go. Can I write an operator in Go?

# Attempt 4

I'm running these operators locally which is fine for development and tests. But if I want to deploy to a massive Kubernetes cluster in production, I need to containerize it.

# Attempt 5

My buddies from my Spring Boot days don't feel like learning Go but are interested in taking advantage of Operator pattern in k8s. I ask them why and they tell me that having to manually spin up new MySql databases is driving them nuts. They heard operators can automate this away: not just create the databases but also initialize them.

# Attempt 6

I hear I can manipulate objects outside the Kubernetes cluster. Huh?

---

Work in Progress

Convert Bash operator to Golang 
(reading Kubernetes Operators book)



