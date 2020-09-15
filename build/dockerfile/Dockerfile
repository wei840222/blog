FROM golang:alpine AS build
ADD . /src
WORKDIR /src
RUN cd cmd && go build -o app

FROM alpine
COPY --from=build /src/cmd/app /app
CMD /app