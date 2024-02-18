FROM node:latest
WORKDIR webapp
ENTRYPOINT ["npm", "start"]
