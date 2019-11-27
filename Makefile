NAME	:=	sfcc_clean
GO		:=	go
FMT		=	gofmt
pkgs	=	$(shell env GO111MODULE=on $(GO) list -m)

FILE	=	main.go\
            auth/manage_auth.go\
            auth/scrape_auth.go\
            auth/type_struct_auth.go\
            blacklist/blacklist_add_module.go\
            blacklist/blacklist_del_module.go\
            blacklist/blacklist_module.go\
            blacklist/blacklist_save_module.go\
            blacklist/blacklist_send_module.go\
            blacklist/blacklist_show_module.go\
            blacklist/type_struct_blacklist.go\
            clean/cust_grp/clean_del_data_module.go\
            clean/cust_grp/clean_get_data_module.go\
            clean/cust_grp/clean_module.go\
            clean/cust_grp/type_struct.go\
            clean/prom_camp/clean_delete_module.go\
            clean/prom_camp/clean_module.go\
            clean/prom_camp/collector_campaign.go\
            clean/prom_camp/collector_promotion.go\
            clean/prom_camp/collector_sites.go\
            clean/prom_camp/data_to_struct.go\
            clean/prom_camp/get_hosts_opts.go\
            clean/prom_camp/get_sites.go\
            clean/prom_camp/manage_camp_struct.go\
            clean/prom_camp/manage_metrics.go\
            clean/prom_camp/manage_prom_struct.go\
            clean/prom_camp/script_delete_module.go\
            clean/prom_camp/script_info_module.go\
            clean/prom_camp/sfcc_query.go\
            clean/prom_camp/type_struct.go\
            clean/prom_camp/type_struct_script.go\
            clean/utils/create_info.go\
            clean/utils/generate_forms.go\
            clean/utils/get_data.go\
            clean/utils/get_hosts.go\
            clean/utils/get_opts.go\
            clean/utils/get_sites.go\
            clean/utils/get_token.go\
            clean/utils/rework_blacklist.go\
            clean/utils/send_format_delete_module.go\
            clean/utils/sfcc_stuff.go\
            clean/utils/type_struct.go\


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
