AWS_PROFILE = ""

deploy:
	GOOS=linux GOARCH=arm64 go build -o bootstrap -ldflags="-s -w" cmd/lambdaC2/main.go
	zip bootstrap.zip bootstrap
	./scripts/deploy.sh $$AWS_PROFILE

destroy:
	./scripts/destroy.sh $$AWS_PROFILE

clean:
	rm -f bootstrap bootstrap.zip