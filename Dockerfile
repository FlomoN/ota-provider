FROM golang:latest AS golang
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a src/main.go
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -a src/main.go

FROM node:latest as node
WORKDIR /app
COPY --from=golang /app .
WORKDIR /app/frontend
RUN npm ci
RUN npm run build

FROM scratch
WORKDIR /app
COPY --from=node /app/main ./
COPY --from=node /app/frontend/build/ ./frontend/build/
EXPOSE 3001
CMD ["/app/main"]
#RUN bash
