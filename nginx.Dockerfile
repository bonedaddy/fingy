FROM nginx:1.19.6-alpine
COPY index.html /usr/share/nginx/html
COPY fingerprint2.js /usr/share/nginx/html
# this is currently work-in-progress
COPY default.conf /etc/nginx/conf.d/default.conf