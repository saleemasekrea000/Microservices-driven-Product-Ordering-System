VERBOSE =1 

.PHONY: run
run:
	docker-compose -f order/docker-compose.yml  up -d
	docker-compose -f  product/docker-compose.yml up -d
	docker-compose -f product_order/docker-compose.yml up -d

.PHONY: stop
stop:
	docker-compose -f order/docker-compose.yml -f product/docker-compose.yml -f product_order/docker-compose.yml down

.PHONY: destroy
destroy:
	docker-compose -f order/docker-compose.yml -f product/docker-compose.yml down --rmi all

.PHONY: install-migration
install-migration:
		curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
		sudo apt-get update
		sudo apt-get install migrate