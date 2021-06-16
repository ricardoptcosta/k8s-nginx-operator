#!/usr/bin/env bash
kubectl get --watch --output-watch-events configmap -o=custom-columns=type:type,name:object.metadata.name,replicas:object.data.ricardoReplicas --no-headers | \
	while read next; do
    NAME=$(echo $next | cut -d' ' -f2)
    EVENT=$(echo $next | cut -d' ' -f1)
    REPLICAS=$(echo $next | cut -d' ' -f3) | `sed -e 's/^"//' -e 's/"$//'` 
		case $EVENT in ADDED|MODIFIED)
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
