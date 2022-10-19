# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

ARG GO_VER=1.16
ARG ALPINE_VER=3.12

FROM golang:${GO_VER}-alpine${ALPINE_VER}

WORKDIR /go/src/github.com/kmilodenisglez/cc-identity-go
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 9999
CMD ["cc-identity-go"]
