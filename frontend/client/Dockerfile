FROM node:16.10-alpine AS build

WORKDIR /app

COPY package*.json ./

RUN npm install
COPY . .
EXPOSE 8081


CMD [ "npm", "start" ]