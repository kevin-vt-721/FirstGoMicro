FROM golang:1.11beta2-alpine2.8 AS build-Getenv

#Allow Go to retrieve the dependencies for the build step
RUN apk add --no-cache git

#Secure against running as root
RUN adduser -D -u 10000 kmlowe
RUN mkdir /firstMicro/ && chown kmlowe /firstMicro/
USER kmlowe

WORKDIR /firstMicro/
ADD . /firstMicro/

#Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED-0 go build -0 /firstMicro/FirstGoMicroservice .

#final stage
FROM alpine:3.8

#Secure against running as root
RUN adduser -D -u 10000 kmlowe
USER kmlowe

WORKDIR /
COPY --from=build-env /firstMicro/certs/docker.localhost.*
COPY --from=build-env /firstMicro/FirstGoMicroservice /

EXPOSE 8080

CMD ["/FirstGoMicroservice"]
