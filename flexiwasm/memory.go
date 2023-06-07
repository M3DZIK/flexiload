package flexiwasm

import "github.com/wasmerio/wasmer-go/wasmer"

// GetOffset gets the offset of the buffer
func GetOffset(instance *wasmer.Instance, functionName string) (int32, error) {
	// Get `bufAddr` function
	f, err := instance.Exports.GetFunction(functionName)
	if err != nil {
		return 0, err
	}

	// Call `bufAddr` function
	result, err := f()
	if err != nil {
		return 0, err
	}

	// Get offset of the buffer
	offset := result.(int32)

	return offset, nil
}

// WriteMemory writes the given bytes to the wasm memory
func WriteMemory(instance *wasmer.Instance, offset int32, bytes []byte) (int32, error) {
	memory, err := instance.Exports.GetMemory("memory")
	if err != nil {
		return 0, err
	}

	memoryData := memory.Data()

	for i, b := range bytes {
		memoryData[offset+int32(i)] = b
	}

	return int32(len(bytes)), nil
}
