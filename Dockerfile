FROM node:alpine
EXPOSE 3000
VOLUME ["/opt/app"]
WORKDIR /opt/app
CMD [ "npm", "start" ]