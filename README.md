# Fingy

Fingy is a CLI tool that provides an API for collecting browser fingerprint information using fingerprintjs, and stores this information in a sqlite database. It is intended for use with penetration testing, and allowing for easy mass fingerprint collection. It collects the following information persisting it into a database:

* murmur fingerprint hash
* user agent
* platform
* language
* ip address

# Dependencies

* Docker
* Docker compose

# Usage

1) Build docker images with `make build`
2) Start docker images using `make run`

# Todo

* make sure nginx correctly passes the client ip address information as we are proxying the request