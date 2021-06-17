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
apiVersion: v1
data:
  revaliases: "# sSMTP aliases\n# \n# Format:\tlocal_account:outgoing_address:mailhub\n#\n#
    Example: root:your_login@your.domain:mailhub.your.domain[:port]\n# where [:port]
    is an optional port number that defaults to 25.\n\nroot:<YOUR-GMAIL-USERNAME>@gmail.com:smtp.gmail.com:587\n"
  ssmtp.conf: "#\n# Config file for sSMTP sendmail\n#\n# The person who gets all mail
    for userids < 1000\n# Make this empty to disable rewriting.\nroot=<YOUR-GMAIL-USERNAME>@gmail.com\n\n#
    The place where the mail goes. The actual machine name is required no \n# MX records
    are consulted. Commonly mailhosts are named mail.domain.com\n#mailhub=mail\nmailhub=smtp.gmail.com:587\n\nAuthUser=<YOUR-GMAIL-USERNAME>@gmail.com\nAuthPass=<YOUR-GMAIL-PASSWORD>\nUseTLS=YES\nUseSTARTTLS=YES\n\n#
    Where will the mail seem to come from?\n#rewriteDomain=\nrewriteDomain=gmail.com\n\n#
    The full hostname\n#hostname=ric-ThinkPad-X1-Carbon-6th\n#hostname=<YOUR-GMAIL-USERNAME>@gmail.com\nhostname=localhost\n\n#
    Are users allowed to set their own From: address?\n# YES - Allow the user to specify
    their own From: address\n# NO - Use the system generated From: address\nFromLineOverride=YES\n\n"
kind: ConfigMap
metadata:
  creationTimestamp: null
  name: ssmtp-conf

EOF

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
            - apt-get update && apt-get --assume-yes install ssmtp && cp /etc/ssmtp-temp/ssmtp.conf /etc/ssmtp/ssmtp.conf && cp /etc/ssmtp-temp/revaliases /etc/ssmtp/revaliases.conf && echo "Get your bum out of the sofa Ricardo, it's time to water the ${PLANT}s" | ssmtp <YOUR-GMAIL-USERNAME>@gmail.com
            volumeMounts:
            - name: ssmtp-conf
              mountPath: /etc/ssmtp-temp
          volumes:
            - name: ssmtp-conf
              configMap:
                name: ssmtp-conf
  schedule: "0 0 */$TIMEINTERVAL * *"

EOF
			   ;;
					DELETED)
                    kubectl delete cronjob $PLANT
                    kubectl delete configmap ssmtp-conf
                    ;;
          esac
done
