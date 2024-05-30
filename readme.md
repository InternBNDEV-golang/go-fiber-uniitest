<div align="center">

# go-fiber-unitest

![img](https://cdn.discordapp.com/attachments/372372440334073859/1216723412244893738/go_1.gif?ex=66016cfb&is=65eef7fb&hm=76c9538616341d981031c54ade56cbeee7eadc96c71d64bbaa6f85cc60436b40&)

</div>

## Install golang package

```bash
go mod tidy
```

# Start APP

```sh
go run . || go run main.go
```

## Air

```sh
go install github.com/cosmtrek/air@latest
```

```sh
air
```

# Podman

```
podman build -t fiber-test .
```

```
podman run --rm -it -p 3000:3000 fiber-test
```

# Write Unittest
- create test case function in `unittest/test` folder
- then call function test case in `unittest/unittest.go`

# Run Unittest
```
go test -v unittest_test.go
```
