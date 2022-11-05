AWS_PROFILE = ""

deploy:
	./scripts/deploy.sh $$AWS_PROFILE

destroy:
	./scripts/destroy.sh $$AWS_PROFILE

clean:
	rm -f bootstrap bootstrap.zip