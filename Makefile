include .env

create:
	yc serverless container create --name $(SERVERLESS_CONTAINER_NAME)
	yc serverless container allow-unauthenticated-invoke --name  $(SERVERLESS_CONTAINER_NAME)

create_gw_spec:
	$(shell sed "s/SERVERLESS_CONTAINER_ID/${SERVERLESS_CONTAINER_ID}/;s/SERVICE_ACCOUNT_ID/${SERVICE_ACCOUNT_ID}/" api-gw.yaml.example > api-gw.yaml)
create_gw: create_gw_spec
	yc serverless api-gateway create --name $(SERVERLESS_CONTAINER_NAME) --spec api-gw.yaml
webhook_info:
	curl --request POST --url "https://api.telegram.org/bot$(TELEGRAM_API_TOKEN)/getWebhookInfo"

webhook_delete:
	curl --request POST --url "https://api.telegram.org/bot$(TELEGRAM_API_TOKEN)/deleteWebhook"

webhook_create: webhook_delete
	curl --request POST --url "https://api.telegram.org/bot$(TELEGRAM_API_TOKEN)/setWebhook" --header 'content-type: application/json' --data "{\"url\": \"$(SERVERLESS_APIGW_URL)\"}"

webhook_drop_updates:
	curl --request POST --url "https://api.telegram.org/bot$(TELEGRAM_API_TOKEN)/setWebhook?url=https://postman-echo.com/post"
	make webhook_delete

build: webhook_create
	go test -v ./...
	docker build -t cr.yandex/$(YC_IMAGE_REGISTRY_ID)/$(SERVERLESS_CONTAINER_NAME) .

push: build
	docker push cr.yandex/$(YC_IMAGE_REGISTRY_ID)/$(SERVERLESS_CONTAINER_NAME)

deploy: push
	$(shell sed 's/=.*/=/' .env > .env.example)
	yc serverless container revision deploy --container-name $(SERVERLESS_CONTAINER_NAME) --image 'cr.yandex/$(YC_IMAGE_REGISTRY_ID)/$(SERVERLESS_CONTAINER_NAME):latest' --service-account-id $(SERVICE_ACCOUNT_ID)  --environment='$(shell tr '\n' ',' < .env)DUMP=0' --core-fraction 50 --execution-timeout $(SERVERLESS_CONTAINER_EXEC_TIMEOUT)

all: deploy

cover:
	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out
