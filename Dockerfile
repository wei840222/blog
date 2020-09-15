FROM golang:alpine AS build
RUN apk add bash ca-certificates git gcc g++ libc-dev
ADD . /src
WORKDIR /src
RUN cd cmd && go build -o app

FROM alpine
COPY --from=build /src/cmd/app /app
CMD /app