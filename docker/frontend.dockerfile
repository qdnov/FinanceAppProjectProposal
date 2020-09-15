FROM node:12 AS base

COPY frontend /usr/app

WORKDIR /usr/app

RUN npm install

CMD ["npm", "start"]