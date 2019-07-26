

default:
	@echo "run 'make build' for build"

build: link go_build unlink

go_build:
	go build

link:
	@mkdir vendor; true
	@cd vendor && ln -s ../restapi || cd restapi && cd ..
	@cd ..;

unlink:
	@cd vendor;unlink restapi;cd ..;rm -d vendor