NAME	:=	sfcc_clean
GO		:=	go
FMT		=	gofmt
pkgs	=	$(shell env GO111MODULE=on $(GO) list -m)

FILE	=	main.go\
            prom_camp/get_hosts_opts.go\
            prom_camp/get_sites.go\
            prom_camp/get_token.go\
            prom_camp/clean_module.go\
            prom_camp/clean_delete_module.go\
            prom_camp/type_struct.go\
            prom_camp/collector_sites.go\
            prom_camp/collector_campaign.go\
            prom_camp/collector_promotion.go\
            prom_camp/data_to_struct.go\
            prom_camp/manage_camp_struct.go\
            prom_camp/manage_metrics.go\
            prom_camp/manage_prom_struct.go\
            prom_camp/script_delete_module.go\
            prom_camp/script_info_module.go\
            prom_camp/sfcc_query.go\
            prom_camp/type_struct.go\
            auth/manage_auth.go\
            auth/scrape_auth.go\
            auth/type_struct_auth.go\
            blacklist/blacklist_module.go\
            blacklist/blacklist_show_module.go\
            blacklist/blacklist_add_module.go\
            blacklist/blacklist_del_module.go\
            blacklist/blacklist_send_module.go\
            blacklist/type_struct_blacklist.go


DOCKER_IMAGE_NAME       ?= sfcc_clean

all: format module build

test:
	@echo ">> running tests"
	@go test -short $(pkgs)

format:
	@echo ">> formatting code"
	@$(FMT) -w $(FILE)

module:
	@echo ">> creating module"
	@$(GO) mod init "github.com/khsadira/cleaner"
	@$(GO) mod vendor

build: 
	@echo ">> building binaries"
	@$(GO) build -o $(NAME)

docker:
	@echo ">> building docker image"
	@docker build -t $(DOCKER_IMAGE_NAME) .

fclean:
	rm -rf $(NAME)
	rm -rf go.mod
	rm -rf go.sum
	rm -rf vendor

re: fclean all test

.PHONY: all format build test
