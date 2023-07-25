FROM golang:alpine AS builder
WORKDIR /app
ARG COMMIT=v0.1.0
LABEL COMMIT=${COMMIT}
ENV GOPROXY=https://goproxy.cn
COPY . .
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache --virtual .build-deps ca-certificates gcc g++
RUN go build -o apis

FROM alpine
WORKDIR /app
COPY --from=builder /app/apis .
COPY --from=builder /app/config.toml .
RUN chmod +x ./apis
EXPOSE 9999

CMD ["/app/apis"]

