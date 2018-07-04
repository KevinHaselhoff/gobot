DOCKERHUB_USERNAME ?= kevinhaselhoff
DOCKERHUB_PASSWORD ?= 
SLACK_TOKEN ?=
IMAGE_NAME := gobot
DOCKER_REGISTRY := kevinhaselhoff
TAG ?=local-dev
DOCKER_IMAGE := ${DOCKER_REGISTRY}/${IMAGE_NAME}

.PHONY: login
login:
		echo "${DOCKERHUB_PASSWORD}" | docker login -u ${DOCKERHUB_USERNAME} --password-stdin

.PHONY: build
build:
		docker build -t ${DOCKER_IMAGE}:${TAG} .

.PHONY: push
push:
		docker push ${DOCKER_IMAGE}:${TAG}

.PHONY: k8sconfig
k8sconfig:
		gcloud auth activate-service-account --key-file=${HOME}/gcloud-service-key.json
		gcloud --quiet config set project ${GOOGLE_PROJECT_ID}
		gcloud --quiet config set compute/zone ${GOOGLE_COMPUTE_ZONE}
		gcloud config set project ${GOOGLE_PROJECT_ID}
		gcloud --quiet container clusters --region ${GOOGLE_COMPUTE_ZONE} get-credentials ${GOOGLE_CLUSTER_NAME}

.PHONY: k8s-deploy
k8s-deploy:
		helm upgrade --install gobot ./deploy/gobot \
		--namespace gobot \
		--set env.slackToken="${SLACK_TOKEN}" \
		--set image.tag="${TAG}"