.PHONY: clean graphs

build:
	@mkdocs build

clean:
	@echo "Cleaning up workspace"

graphs:
	@echo "Generating graphs"
	@rm -rf docs/static/img/graphs/*.png
	@rm -rf docs/static/img/network-server/graphs/*.png
	@cd docs/static/img/graphs && dot -Tpng -O *.dot
	@cd docs/static/img/network-server/graphs && dot -Tpng -O *.dot
