# knative-eventing-in-mem-demo


##  Get ENV for cluster/node

    kubectl get node --namespace istio-system  --output 'jsonpath={.items[0].status.addresses[0].address}'

    kubectl get svc knative-ingressgateway --namespace istio-system --output 'jsonpath={.spec.ports[?(@.port==80)].nodePort}'

    export HOST_IP=$(kubectl get node --namespace istio-system  --output 'jsonpath={.items[0].status.addresses[0].address}')

    export HOST_PORT=$(kubectl get svc knative-ingressgateway --namespace istio-system --output 'jsonpath={.spec.ports[?(@.port==80)].nodePort}')

    export KNATIVE_INGRESS=$HOST_IP:$HOST_PORT


## Deployment

    kubectl create ns knative-eventing-in-mem-domain

    ----kubectl apply -f https://github.com/knative/eventing/releases/download/v0.4.0/release.yaml----

    kubectl apply -f https://storage.googleapis.com/knative-releases/eventing/latest/release.yaml

    kubectl apply -f https://github.com/knative/eventing-sources/releases/download/v0.4.0/release.yaml

    kubectl apply -f https://raw.githubusercontent.com/knative/build-templates/master/kaniko/kaniko.yaml

    kubectl apply -f secret.yaml

    kubectl apply -f sa-build-bot.yaml

    kubectl apply -f sa-k8events.yaml

    kubectl apply -f service.yaml

    kubectl apply -f channel.yaml

    kubectl apply -f source.yaml

    kubectl apply -f subscription.yaml


## Get ENV for curl

    kubectl get route knative-eventing-in-mem-demo --output=custom-columns=NAME:.metadata.name,DOMAIN:.status.domain

    kubectl get route knative-eventing-in-mem-demo --output jsonpath='{.status.domain}'

    kubectl get route knative-eventing-in-mem-demo --output go-template --template='{{.status.domain}}{{"\n"}}'

    export HOST_FQDN=$(kubectl get route knative-eventing-in-mem-demo --output jsonpath='{.status.domain}' )

    echo $HOST_FQDN $KNATIVE_INGRESS

## Watch

    watch -n 1 -t kubectl get pod,deploy,eventing,serving

## Verify service

    curl -H "Host: $HOST_FQDN" http://$KNATIVE_INGRESS/ -w "\n"


## Create/Generate Events


    kubectl run -i --tty busybox --image=busybox --restart=Never   -- sh
    Once the shell comes up, just exit it and kill the pod.

    kubectl delete pod busybox

## Verify Events


    kubectl logs  pod/knative-eventing-in-mem-demo-00001-deployment-5849fb7d69-qz4hj   -c user-container --tail=10 -f
    kubectl logs  pod/knative-eventing-demo-00001-deployment-5849fb7d69-qz4hj   -c user-container --tail=10 -f

    kubectl logs -l serving.knative.dev/service=knative-eventing-in-mem-demo -c user-container
    kubectl logs -l serving.knative.dev/service=knative-eventing-demo -c user-container


 Based on the official Knative documentation, licensed under the Creative Commons Attribution 4.0 License,
 and code samples are licensed under the Apache 2.0 License