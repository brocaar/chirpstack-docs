.PHONY: build clean graphs devshell

build: graphs
	@echo "Building documentation"
	mkdocs build

clean:
	@echo "Cleaning up workspace"
	rm -rf site

graphs:
	@echo "Generating graphs"
	@rm -rf docs/static/img/graphs/*.png
	@rm -rf docs/static/img/network-server/graphs/*.png
	@cd docs/static/img/graphs && dot -Tpng -O *.dot
	@cd docs/static/img/network-server/graphs && dot -Tpng -O *.dot

devshell:
	@echo "Starting devshell"
	docker-compose run --rm chirpstack-docs bash

serve:
	@echo "Serving documentation"
	docker-compose run --rm --service-ports chirpstack-docs mkdocs serve -a 0.0.0.0:3000
