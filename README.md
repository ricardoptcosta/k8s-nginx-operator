# k8s-wateringalarm-operator

This repo is a step by step approach to creating a kubernetes operator following a natural self discovery path. The end goal here is to create a bad ass operator!

## Requirements:

Minikube v1.22+  
Kubectl


## Iteration 1 

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


## Iteration 2

Using an existing kubernetes resources like a ConfigMap is cool, but what I would really like to do is to create my own custom resource. What's simple and useful?

I love plants. I have many plants. I forget to water them when I should. They die. I get sad. I need to address this so that I can get my life in order. Could a Kubernetes CustomResourceDefinition come to the rescue?

- Create CRD and deploy it
  1. Create CRD wateringalarm-crd.yml
  2. Run `$ kubectl apply -f iteration-2/wateringalarm-crd.yml`

Boom! We have a new CRD. Check `$ kubectl api-resources` to double check we have `wateringalarms` as a new api resource

- Create custom resources of type WateringAlarm. In this case we will build an alarm for orchids and succulents
  1. Create orchid-wateringalarm.yml and succulent-wateringalarm.yml
  2. Run  `$ kubectl apply -f iteration-2/orchid-wateringalarm.yml && kubectl apply -f iteration-2/succulent-wateringalarm.yml`

Now, I have the resources created, how do I take advantage of the operator pattern to leverage them?
- Create a watcher on the operator
- Do something when the the watcher identifies a change

This is exactly the same as on iteration-1, now with my beautiful custom resource. See iteration-2/operator.sh for my implementation.

After launching the watcher I create a loop that deploys or updates a cronjob whenever a WateringAlarm is created/updated. It deletes the cronjob whenever a WateringAlarm resource is deleted.

Since I need the container created by the cronjob to send me an email reminding me about watering the plants, I used an ubuntu image and installed ssmpt. There is a bunch of configuration required for the email protocol. See this [link](https://www.havetheknowhow.com/Configure-the-server/Install-ssmtp.html). Following that link, I created a configmap from the files `/etc/ssmtp/revaliases` and ` /etc/ssmtp/ssmtp.conf` and pasted it on the `operator.sh file` with the name `ssmtp-conf`. These files were then mounted on the containter via the configmap. I used a Gmail account to act as my email server.

Sweet lord, my watering alarms are working and my operator is making sure any creation, update or deletion of a WateringAlarm is reconciled with reality! To be more specific, if I realize my cactus need watering once a month, after I create the cactus-wateringalarm custom resource, the operator will create a cronjob that will spin up a pod which sends me a friendly email reminder to water my cacti.


TODO -------------- For the credentials I used a secret. IF HAVE TIME REVISE THIS LAST BIT

## Iteration 3 (to be confirmed)
Same as iteration 2 but to implement a mysql database

## Iteration 4

Bash is nice, but I hear Kubernetes is written in Go. Can I write an operator in Go?

--- USE GO SDK

## Iteration 5

My friend told me he used *Kubebuilder* to scaffold his operator. Shall we try? 
First think, we have to install kubebuilder and Kustomize. Why both?

I am starting the project by initializating an empty project. I used iteration5/ as the root of the project.

`go mod init wateringalarm`

Then I initialize the controller project. What is the controller?

`kubebuilder init --domain ricardoptcosta.github.io`

Then I ask Kubebuilder to scaffold a Kubernetes API by creating a Custom Resource Definition and the Controller.

`kubebuilder create api --resource --controller --group alarm --version v1alpha1 --kind WateringAlarm `
Following Kubebuilder help page, I then edit the API scheme on `api/v1alpha1/wateringalarm_types.go. In the struct WateringAlarmSpec I replace the Foo field with the fields Plant and TimeInterval.

Then, edit the controller in `controllers/wateringalarm_controller.go`. In particular, implement the operator's logic in the Reconcile function



## Iteration 6

My friend's dad told us nobody ain't got time for that and he simply uses the operator-sdk. Could it be that great?

--- USE OPERATOR SDK

## Iteration 7


I'm running these operators locally which is fine for development and tests. But if I want to deploy to a massive Kubernetes cluster in production, I need to containerize it.

## Iteration 8

My buddies from my Spring Boot days don't feel like learning Go but are interested in taking advantage of Operator pattern in k8s. I ask them why and they tell me that having to manually spin up new MySql databases is driving them nuts. They heard operators can automate this away: not just create the databases but also initialize them.

--- USE JAVA OPERATOR SDK

## Iteration 9

I hear I can manipulate objects outside the Kubernetes cluster. Huh?

---

Work in Progress

Convert Bash operator to Golang 
(reading Kubernetes Operators book)



