FROM node:alpine AS build-web
ADD ./web /src
WORKDIR /src
ENV VUE_APP_GRAPHQL_HTTP=https://fuck-word-blog.herokuapp.com/graphql
ENV VUE_APP_GRAPHQL_WS=wss://fuck-word-blog.herokuapp.com/graphql
RUN npm i
RUN npm run build

FROM golang:alpine AS build
RUN apk add bash ca-certificates git gcc g++ libc-dev
ADD . /src
WORKDIR /src
RUN cd cmd && go build -o app

FROM alpine
COPY --from=build-web /src/dist /web/dist
COPY --from=build /src/cmd/app /app
CMD /app