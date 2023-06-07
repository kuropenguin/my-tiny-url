FROM golang:1.19.5

RUN apt update && apt upgrade
RUN go install github.com/cosmtrek/air@v1.29.0


WORKDIR /app

COPY app/go.mod app/go.sum app/.air.toml ./ 

CMD ["air", "-c", ".air.toml"]
# CMD ["bash"]
