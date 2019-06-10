#!/bin/sh
export HOST_FQDN=$(kubectl get route knative-eventing-in-mem-demo --output jsonpath='{.status.domain}' -n knative-eventing-in-mem-domain)
echo $HOST_FQDN $KNATIVE_INGRESS