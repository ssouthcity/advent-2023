FROM golang:alpine

WORKDIR /usr/src/advent2023

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /advent2023

ENV PORT=80

EXPOSE 80

CMD [ "/advent2023" ]
