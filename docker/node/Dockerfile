FROM node:14.10

WORKDIR /usr/src/app

COPY ./front/package.json ./
COPY ./front/yarn.lock ./

RUN yarn set version berry
RUN yarn set version 2.4.1

COPY ./front/.yarnrc.yml ./

RUN yarn install

COPY . ./
