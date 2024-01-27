FROM golang:1.21 as build

ENV BIN_FILE /opt/apka
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS

RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} ./cmd

FROM alpine:3.9

ENV WORK_DIR /opt

ENV BIN_FILE ${WORK_DIR}/apka
ENV CONFIG /go/src/config/.env

COPY --from=build ${BIN_FILE} ${BIN_FILE}
COPY --from=build ${CONFIG} ${WORK_DIR}/config/.env

WORKDIR ${WORK_DIR}
CMD ${BIN_FILE}
