GOCMD=go
GOBUILD=$(GOCMD) build
GOBUILDPROD=$(GOCMD) build -ldflags "-linkmode external -extldflags -static" 
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
SODA=buffalo db
BUFFALO=buffalo
APPNAME=gbfm

build:
	$(GOBUILD) -v -o gbfm

buildprod:
	$(GOBUILDPROD) -v -o gbfm

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/gbfm
	rm -rf gbfm

test:
	$(GOTEST) -v ./grifts -race
	$(GOTEST) -v ./models -race
	$(GOTEST) -v ./actions -race

db-up:
	docker run --name=gbfm_db -d -p 5432:5432 -e POSTGRES_DB=gbfm_development postgres
	sleep 10
	$(SODA) migrate up
	docker ps | grep gbfm_db

db-down:
	docker stop gbfm_db
	docker rm gbfm_db


teardown-dev: clean
	$(DOCKERCOMPOSE) down

dev:
	$(BUFFALO) dev

