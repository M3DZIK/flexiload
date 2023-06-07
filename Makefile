all: build-examples

build: build-examples
build-examples: example-modules example-loader

example-modules:
	@mkdir -p build/modules
	@echo "Building example module..."
	@cd ./examples/wasm/module && tinygo build -o ./../../../build/modules/module.wasm -target wasi .

example-loader:
	@mkdir -p build
	@echo "Building example loader..."
	@go build -o ./build/wasm-loader ./examples/wasm/loader

clean:
	@echo "Cleaning up..."
	@rm -rf ./build
