#!/bin/sh

echo "############ EVENTING DEMO ##############"

kubectl create ns knative-eventing-in-mem-domain
kubectl apply -f https://github.com/knative/eventing/releases/download/v0.4.0/release.yaml
kubectl apply -f https://github.com/knative/eventing-sources/releases/download/v0.4.0/release.yaml
kubectl apply -f https://raw.githubusercontent.com/knative/build-templates/master/kaniko/kaniko.yaml -n knative-eventing-in-mem-domain
kubectl apply -f secret.yml
kubectl apply -f sa-build-bot.yml
kubectl apply -f sa-k8events.yml
kubectl apply -f service.yml
kubectl apply -f channel.yml
kubectl apply -f source.yml
kubectl apply -f subscription.yml