## default: run and build
default: clean build

## clean: clean the dist directory
clean:
	rm -rf ./dist/

## build: clean the dist directory and build web app
build: clean
	npm --prefix ui ci
	npm run --prefix ui build
	mv ./ui/dist/ ./dist/
	make -C engine
	mkdir -p ./dist/bin/
	mv ./engine/dist/* ./dist/bin/

## help: print this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
