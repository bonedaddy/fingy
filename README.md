# Fingy

Fingy is a CLI tool that provides an API for collecting browser fingerprint information using fingerprintjs, and stores this information in a sqlite database. It is intended for use with penetration testing, and allowing for easy mass fingerprint collection. It collects the following information persisting it into a database:

* murmur fingerprint hash
* user agent
* platform
* language
* ip address

# Usage

Fingy is designed to be used with the included docker images and docker compose file. 

# Todo

* make sure nginx correctly passes the client ip address information as we are proxying the request