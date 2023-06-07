package main

import (
	"os"

	"github.com/M3DZIK/flexiload/flexiwasm"
)

func main() {
	// Create a new instance of the wasm engine and store
	flexiLoadInstance := flexiwasm.NewFlexiLoad()

	// Open the wasm file
	wasmBytes, err := os.ReadFile("build/modules/module.wasm")
	if err != nil {
		panic(err)
	}

	// Load the wasm module into the store
	module, err := flexiLoadInstance.LoadModule(wasmBytes, "module")
	if err != nil {
		panic(err)
	}

	// Load the wasi environment and module into the store and return an instance
	instance, err := flexiLoadInstance.LoadWasi(module)
	if err != nil {
		panic(err)
	}

	// Get the offset of the buffer
	offset, err := flexiwasm.GetOffset(instance, "bufAddr")

	// Write the given bytes to the wasm memory
	len, err := flexiwasm.WriteMemory(instance, offset, []byte("Dance Macabre"))
	if err != nil {
		panic(err)
	}

	// Get the `test` function
	test, err := instance.Exports.GetFunction("test")
	if err != nil {
		panic(err)
	}

	// Call the `test` function
	_, err = test(offset, len)
}
