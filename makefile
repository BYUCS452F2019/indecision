NAME := $(shell basename "$(PWD)")

# go
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# angular
NPM=npm
NPM_INSTALL=$(NPM) install
NG_BUILD=ng build --prod --aot --build-optimizer 
NG1=web

all: build

build: build-x86  build-web

build-x86:
	env GOOS=linux CGO_ENABLED=0 $(GOBUILD) -o $(NAME)-bin -v

build-web: $(NG1)
	# ng1
	cd $(NG1) && $(NPM_INSTALL) && $(NG_BUILD) --base-href="/"
	rm -rf $(NG1)-dist
	mv -f $(NG1)/dist/$(NG1) $(NG1)-dist

clean: 
	$(GOCLEAN)
	rm -f $(NAME)-bin
	rm -rf $(NG1)-dist

run: $(NAME)-bin $(NG1)-dist
	./$(NAME)-bin

deps:
	$(GOGET) -d -v

$(NAME)-bin:
	$(MAKE) build-x86

$(NG1)-dist:
	$(MAKE) build-web