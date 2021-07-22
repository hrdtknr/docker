FROM golang:1.16.5-alpine3.14

WORKDIR /user/src/app

COPY main.go ./

EXPOSE 3000

CMD ["go", "run", "main.go"]
# ダメだった書き方
# CMD ["app", "go run main.go"]
# CMD ["go run", "main.go"]

## 参考コマンド
# docker run -p 8080:3000 -d express-hello-world