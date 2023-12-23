FROM golang:1.19.5

RUN apt update -y && apt upgrade -y
RUN go install github.com/cosmtrek/air@v1.29.0


WORKDIR /app

COPY app/ /app/


CMD ["air", "-c", ".air.toml"]
# CMD ["bash"]
