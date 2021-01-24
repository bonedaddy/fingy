.PHONY: build
build: docker-build-nginx docker-build-api

.PHONY: run
run: build
	docker-compose up -d

.PHONY: download-fingerprintjs2
download-fingerprintjs2:
	wget -O master.zip https://github.com/fingerprintjs/fingerprintjs/archive/2.1.4.zip
	unzip master.zip
	mv fingerprintjs-2.1.4 fpjs2
	rm master.zip
	mv fpjs2/index.html index.html
	mv fpjs2/fingerprint2.js fingerprint2.js
	rm -rf fpjs2


.PHONY: docker-build-nginx
docker-build-nginx:
	docker build -f nginx.Dockerfile . -t fingy-nginx

.PHONY: docker-build-api
docker-build-api:
	docker build -f fingy.Dockerfile . -t fingy-api