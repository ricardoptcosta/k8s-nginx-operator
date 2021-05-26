#!/usr/bin/env bash

kubectl get --watch --output-watch-events wateringalarm \
-o=custom-columns=type:type,name:object.metadata.name,plant:object.spec.plant,timeinterval:object.spec.timeinterval --no-headers	| \
	while read next; do
		PLANT=$(echo $next | cut -d' ' -f3)
    EVENT=$(echo $next | cut -d' ' -f1)
	  TIMEINTERVAL=$(echo $next | cut -d' ' -f4)
		case $EVENT in
                  ADDED|MODIFIED)

			  kubectl apply -f - << EOF
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: $PLANT
spec:
  jobTemplate:
    metadata:
      name: $PLANT
    spec:
      backoffLimit: 5
      activeDeadlineSeconds: 30
      template:
        spec:
          restartPolicy: Never
          containers:
          - image: ubuntu
            name: $PLANT
            command:
            - /bin/sh
            - -c
            - apt update && apt install mailutils && mail -s 'Time to water $PLANTs' ricardoptcosta@gmail.com
  schedule: "* * * * *"

EOF
			   ;;
					DELETED)
                    kubectl delete cronjob $PLANT
                    ;;
          esac
done
