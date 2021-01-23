# Fingy

Fingy is a CLI tool that provides an API for collecting browser fingerprint information using fingerprintjs, and stores this information in a sqlite database. It is intended for use with penetration testing, and allowing for easy mass fingerprint collection.

A corresponding HTML page in `index.html` is included that serves as a template to be included in your web pages.

# Usage

Before using this you will want to change the IP address the information is sent to defined near the bottom of `index.html`. It defaults to sending information to http://localhost:6969/submit