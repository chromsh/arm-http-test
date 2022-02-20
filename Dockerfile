FROM golang:latest AS build
WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build  -o /bin/main

FROM scratch
COPY --from=build /bin/main /bin/main
ENTRYPOINT [ "/bin/main" ]