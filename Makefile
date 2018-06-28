DOCKERHUB_USERNAME ?= kevinhaselhoff
DOCKERHUB_PASSWORD ?= 
SLACK_TOKEN ?=
IMAGE_NAME := gobot
DOCKER_REGISTRY := kevinhaselhoff
TAG ?=local-dev
DOCKER_IMAGE := ${DOCKER_REGISTRY}/${IMAGE_NAME}

.PHONY: login
login:
		docker login -u ${DOCKERHUB_USERNAME} -p ${DOCKERHUB_PASSWORD}

.PHONY: build
build:
		docker build -t ${DOCKER_IMAGE}:${TAG} .

.PHONY: push
push:
		docker push ${DOCKER_IMAGE}:${TAG}

.PHONY: k8s-deploy
k8s-deploy:
		helm upgrade --install gobot ./deploy/gobot \
		--namespace gobot \
		--set env.slackToken="${SLACK_TOKEN}"