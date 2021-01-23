FROM nginx:1.19.6-alpine
COPY index.html /usr/share/nginx/html
COPY fingerprint2.js /usr/share/nginx/html