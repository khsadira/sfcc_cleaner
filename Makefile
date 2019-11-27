NAME	:=	sfcc_clean
GO		:=	go
FMT		=	gofmt
pkgs	=	$(shell env GO111MODULE=on $(GO) list -m)

FILE	=	main.go\



DOCKER_IMAGE_NAME       ?= sfcc_clean

all: format module build

test:
	@echo ">> running tests"
	@go test -short $(pkgs)

format:
	@echo ">> formatting code"
	@$(FMT) -w $(FILE)

module:
	@if [ ! -d "./vendor" ]; then \
	    echo ">> creating module"; \
	    $(GO) mod init "github.com/khsadira/cleaner"; \
	    $(GO) mod vendor; \
	fi

build: 
	@echo ">> building binaries"
	@$(GO) build -o $(NAME)

docker:
	@echo ">> building docker image"
	@docker build -t $(DOCKER_IMAGE_NAME) .

clean:
	rm -rf ${NAME}

fclean: clean
	rm -rf go.mod go.sum vendor

re: fclean module all test

.PHONY: all clean fclean format build test
