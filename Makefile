.PHONY: documentation

clean:
	@echo "Cleaning up workspace"
	@rm -rf public

documentation:
	@echo "Building documentation"
	@mkdir -p dist/docs
	@hugo
