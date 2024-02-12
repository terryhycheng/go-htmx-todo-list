FROM node:20-alpine as tailwind-builder

WORKDIR /app

COPY . .

RUN npm install
RUN npm run tailwindcss:build


FROM golang:1.22 as go-builder

WORKDIR /app

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go mod tidy

COPY --from=tailwind-builder /app/build/css ./build/css
RUN templ fmt . && templ generate && go build -o ./build ./cmd/go-todo-list/main.go


FROM golang:1.22 as go-runner

WORKDIR /app

COPY --from=go-builder /app/build ./build
COPY --from=go-builder /app/assets ./assets
COPY .env.prod .env

EXPOSE 3000
CMD ["./build/main"]

