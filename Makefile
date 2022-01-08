.PHONY : 

include .env
export

run:    
	go run main.go 

rund:
	docker build -t syedomair/ex-paygate-void:latest \
	 --build-arg service_name=$(SERVICE_NAME) \
         --build-arg log_level=$(LOG_LEVEL) \
         --build-arg port=$(PORT) \
         --build-arg database_url=$(DATABASE_URL_DOCKER) \
         --build-arg signingkey=$(SIGNINGKEY)  .
	docker container run  -p 8322:8322 --name ex-paygate-void syedomair/ex-paygate-void:latest

test_v:    
	go test ./... -v

test_r:    
	go test ./... -race

