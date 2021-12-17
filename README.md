# Go GraphQL

🚀 GraphQL Server built by Go (Golang) + Dependencies Injection with Google Wire

**Flow Chart**

![flowchart](graphql-go.jpg)

## Project Structure

```bash
├── config                      # contains all env file for different environments
├── environment                 # dockerfile
└── server                      # contains go server
```

## Features:

- [x] DI - [Google Wire](https://github.com/google/wire)
- [x] Logging - [Logrus](https://github.com/Sirupsen/logrus)
- [x] Linter - [Go Linter](https://github.com/golangci/golangci-lint)
- [x] Docker Support
- [ ] Pre commit
- [ ] Dataloader
- [ ] Caching
- [ ] Testing

## Guidelines:

Run go lint to check syntax
```bash
make go-lint
```
