FROM golang:1.19.5

RUN apt update -y && apt upgrade -y
RUN go install github.com/cosmtrek/air@v1.29.0
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

WORKDIR /app

COPY app/go.mod app/go.sum app/.air.toml ./ 

CMD ["air", "-c", ".air.toml"]
# CMD ["bash"]
