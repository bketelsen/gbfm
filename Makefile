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
	docker-compose -p dev up -d dev
	sleep 10
	$(SODA) migrate up

db-down:
	docker-compose -p dev down

dev:
	CMS_URL=http://localhost:8080 $(BUFFALO) dev

