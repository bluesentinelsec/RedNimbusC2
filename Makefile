all: implant redOperator lambdaC2

implant:
	go build -o release/redNimbusC2-implant -ldflags="-s -w" cmd/implant/main.go

redOperator:
	go build -o release/redNimbusC2-operator -ldflags="-s -w" cmd/redOperator/main.go

lambdaC2:
	go build -o release/redNimbusC2-lambda -ldflags="-s -w" cmd/lambdaC2/main.go

test:
	go test ./pkg/shellexec/
	go test ./pkg/cli

# intended to be run locally so that
# GitHub Actions does not rack up
# expenses under my AWS account
test-aws:
	go test ./...