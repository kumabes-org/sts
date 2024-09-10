# STS - Security Token Service

```golang
go mod init
go mod tidy
go build main.go
./main
go run main
```

```
docker build -t sts:0.0.1 .

docker run --name sts \
    -d \
    -p 7788:7788 \
    sts:0.0.1 
```

## Links
- https://pkg.go.dev/google.golang.org/api/sts/v1beta
- https://medium.com/@dd_de_b/aws-sts-assumerole-in-go-441712307491
- https://permify.co/post/jwt-authentication-go/ 
- https://pascalallen.medium.com/jwt-authentication-with-go-242215a9b4f8