.PHONY: documentation

build: static/img/graphs
	@echo "Building documentation"
	@mkdir -p dist/docs
	@hugo

clean:
	@echo "Cleaning up workspace"
	@rm -rf public
	@rm -rf static/img/graphs

static/img/graphs:
	@mkdir static/img/graphs
	@dot -Tpng -o static/img/graphs/architecture.png docs/architecture.dot

