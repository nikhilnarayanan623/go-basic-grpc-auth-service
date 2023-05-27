FROM golang:1.20.4-alpine3.18 AS build-stage
WORKDIR /home/app
COPY ./ /home/app
RUN cd /home/app && go mod download
RUN mkdir -p /home/app/bulid
RUN go build -o ./build/api ./cmd/main.go

FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /home/app/build/api /api
CMD ["/api"]