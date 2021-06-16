# k8s-wateringalarm-operator

This repo is a step by step approach to creating a kubernetes operator following a natural self discovery path. The end goal here is to create a bad ass operator!

## Requirements:

Minikube v1.22+  
Kubectl


## Iteration 1 

What is a Kubernetes operator? People at the office keep talking about it. There are a few concepts I hear that make sense. I hear it takes advantage of Kubernetes control-loop. The control-loop is a mechanism by which Kubernetes compares a desired state of the world with the real state of the world.Take for instance a deployment with 3 pod replicas. Whenever I manually destroy a pod in my local Minikube cluster, another one will be spun off if that pod. Here, initially, the desired state of the world matched the real state of the world, ie, 3 pods. However, once I destroy a pod, the real state of the world is now only 2 pods. As such, the control loop needs to reconcile this mismatch what is desired and what is real. Now, this concept is implemented via a Kubernetes component called the Deployment Controller. This controller watches for any changes both on the resource specifications and on the cluster deployed resources. This is it! Ignoring for now the role of the ReplicaSet controller, the Deployment controller informs the scheduler that it needs to assign a new pod to a suitable node. The scheduler then lets the api server which kubelet it has to instruct to create a new pod.

This is a bit intense. Maybe a simple drawing would help clear it out.

DEPLOYMENT CONTROLLER FLOW SKETCH HERE

I don't know how this deployment controller is implemented, but I think that maybe I could create one myself. This is a big challenge. Where shall I put it? 

I will create a script which will watch for any updates on how many pods of Nginx I want running. How do I pass this information? Configmaps are used to pass information around, so maybe I can use that. I will create a sample configmap yaml which will convey the number of nginx replicas that I want. 

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: ricardosdeployment
data:
  ricardosReplicas: "3" 
```

Awesome stuff, now that I have the piece of information that I want and the transmission vessel, let's get onto the script. I will use `bash` as I need to use `kubectl` and I'm not particularly familiar with language specific [Kubernetes client libraries](https://kubernetes.io/docs/reference/using-api/client-libraries/). 

I need to start the script by telling the api server to look for new configmaps, as well as updates and deletions. 

`iteration-0/operator.sh`
```bash
kubectl get --watch --output-watch-events configmap -o=custom-columns=TYPE:type,NAME:object.metadata.name,REPLICAS:object.data.ricardosReplicas 
```

With two terminals open, I can run the `kubectl` in one and then I can apply the `sample-configmap.yml` on the other one. On the former, I see it delivers a table in our familiar `kubectl` output format and there it is, a line which states that the configmap regarding ricardosdeployment was added to the configmap resources.

  | TYPE  | NAME               | REPLICAS |
  |--|--|--|
  | ADDED | kube-root-ca.crt   | \<none>   |
  | ADDED | ricardosdeployment | 3        |
  |       |                    |          |

This is great. I now want to transpose this information to a replicaset which will create or delete pods in order to make my wishes come true.

To do so, I add a few line to my script which do the following: 

1. fetch the pieces of data I care about and store them in the variables EVENT, NAME and REPLICAS. With such information, whenever an EVENT takes place to any configmap, for instance whenever I modify the number of replicas from 3 to 2, this script will run and will apply the changes to the real world via a heredoc. If the replicaset already exists it will be updated or deleted. If it doesn't it will be created. 

`iteration-0/operator.sh`
```bash
kubectl get --watch --output-watch-events configmap -o=custom-columns=type:type,name:object.metadata.name,replicas=object.data.customReplicas --no-headers | \
	while read next; do
    EVENT=$(echo $next | cut -d' ' -f1)
		NAME=$(echo $next | cut -d' ' -f2)
    REPLICAS=$(echo $next | cut -d' ' -f3)
	
		case $EVENT in
                  ADDED|MODIFIED)
			  kubectl apply -f - << EOF
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: $NAME
  labels:
    app: $NAME
spec:
  replicas: $REPLICAS
  selector:
    matchLabels:
      app: $NAME
  template:
    metadata:
      labels:
        app: $NAME
    spec:
      containers:
      - name: nginx-webserver
        image: nginx

EOF
			   ;;
					DELETED)
                    kubectl delete replicaset $NAME
                    ;;
          esac
done
``` 


Let's test our baby. 




------iteration 2

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

------ This is powerful, but I still don't get how can I use this to my advantage. All I learnt is that Deployments take care of making sure the number of pods running match what I want. What I came to realise is that the key thing is, that I can apply this same rationale to any other type of objects ---


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

The above command creates the go.mod file which will contain the project dependencies among other details.

Then I initialize the project. 

`kubebuilder init --domain ricardoptcosta.github.io`

At this point I have the following folder structure:

	.
	├── bin
	│   └── manager
	├── config
	│   ├── certmanager
	│   ├── default
	│   ├── manager
	│   ├── prometheus
	│   ├── rbac
	│   └── webhook
	├── Dockerfile
	├── go.mod
	├── go.sum
	├── hack
	│   └── boilerplate.go.txt
	├── main.go
	├── Makefile
	└── PROJECT

Then I ask Kubebuilder to scaffold a Kubernetes API by creating a Custom Resource Definition and the Controller.

`kubebuilder create api --resource --controller --group alarm --version v1alpha1 --kind WateringAlarm `  

This command creates the `api` and `controllers` folders. At this point I have the following folder strcuture:

    .
    ├── api
    │   └── v1alpha1
    │       ├── groupversion_info.go
    │       ├── wateringalarm_types.go
    │       └── zz_generated.deepcopy.go
    ├── bin
    │   └── manager
    ├── config
    │   ├── certmanager
    │   │   ├── certificate.yaml
    │   │   ├── kustomization.yaml
    │   │   └── kustomizeconfig.yaml
    │   ├── crd
    │   │   ├── kustomization.yaml
    │   │   ├── kustomizeconfig.yaml
    │   │   └── patches
    │   │       ├── cainjection_in_wateringalarms.yaml
    │   │       └── webhook_in_wateringalarms.yaml
    │   ├── default
    │   │   ├── kustomization.yaml
    │   │   ├── manager_auth_proxy_patch.yaml
    │   │   ├── manager_webhook_patch.yaml
    │   │   └── webhookcainjection_patch.yaml
    │   ├── manager
    │   │   ├── kustomization.yaml
    │   │   └── manager.yaml
    │   ├── prometheus
    │   │   ├── kustomization.yaml
    │   │   └── monitor.yaml
    │   ├── rbac
    │   │   ├── auth_proxy_client_clusterrole.yaml
    │   │   ├── auth_proxy_role_binding.yaml
    │   │   ├── auth_proxy_role.yaml
    │   │   ├── auth_proxy_service.yaml
    │   │   ├── kustomization.yaml
    │   │   ├── leader_election_role_binding.yaml
    │   │   ├── leader_election_role.yaml
    │   │   ├── role_binding.yaml
    │   │   ├── wateringalarm_editor_role.yaml
    │   │   └── wateringalarm_viewer_role.yaml
    │   ├── samples
    │   │   └── alarm_v1alpha1_wateringalarm.yaml
    │   └── webhook
    │       ├── kustomization.yaml
    │       ├── kustomizeconfig.yaml
    │       └── service.yaml
    ├── controllers
    │   ├── suite_test.go
    │   └── wateringalarm_controller.go
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── hack
    │   └── boilerplate.go.txt
    ├── main.go
    ├── Makefile
    └── PROJECT
      

Following Kubebuilder help page, I then edit the API scheme on `api/v1alpha1/wateringalarm_types.go`. In the struct WateringAlarmSpec I replace the Foo field with the fields Plant and TimeInterval.
Change this section of the code  
`$ vim api/v0alpha1/wateringalarm_types.go`
```golang
27 type WateringAlarmSpec struct {
28         // INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
29         // Important: Run "make" to regenerate code after modifying this file
30 
31         // Foo is an example field of WateringAlarm. Edit WateringAlarm_types.go to remove/updat    e
32         Foo string `json:"foo,omitempty"`
33 }
```
into
```golang
 27 type WateringAlarmSpec struct {
 28         Plant string `json:"plant,omitempty"`
 29
 30         //+kubebuilder:validation:Minimum=0
 31         TimeInterval int32 `json:"timeinterval,omitempty"`
 32 }

```
Note that on that file, `wateringalarm_types.go`, we only edit the Spec of the custom resource. The actual WateringAlarm struct, further down on the file, only makes use of the WateringAlarmSpec and WateringAlarmStatus structs. Also, I add a marker that ensures that TimeInterval is bounded at zero.

Run `make manifests` to create the CRD.

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



