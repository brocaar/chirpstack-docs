.PHONY: documentation

clean:
	@echo "Cleaning up workspace"
	@rm -rf public

documentation:
	@echo "Building documentation"
	@mkdir -p dist/docs
	@hugo
	@cd public/ && tar -pczf ../dist/docs/docs.loraserver.io.tar.gz .

