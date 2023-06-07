package flexiwasm

import "github.com/wasmerio/wasmer-go/wasmer"

// FlexiLoad is a struct that holds the engine and store from wasmer
type FlexiLoad struct {
	engine *wasmer.Engine
	store  *wasmer.Store
}

// NewInstance creates a new instance of the wasm engine and store
func NewFlexiLoad() *FlexiLoad {
	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	return &FlexiLoad{
		engine: engine,
		store:  store,
	}
}

// LoadModule loads the wasm module into the store
func (i *FlexiLoad) LoadModule(wasmBytes []byte, programName string) (*wasmer.Module, error) {
	return wasmer.NewModule(i.store, wasmBytes)
}

// LoadWasi loads the wasi environment and module into the store and returns an instance
func (i *FlexiLoad) LoadWasi(module *wasmer.Module) (*wasmer.Instance, error) {
	wasiEnv, err := wasmer.NewWasiStateBuilder("wasi-program").
		// Choose according to your actual situation
		// Argument("--foo").
		// Environment("ABC", "DEF").
		// MapDirectory("./", ".").
		Finalize()
	if err != nil {
		return nil, err
	}

	// Instantiates the module
	importObject, err := wasiEnv.GenerateImportObject(i.store, module)
	if err != nil {
		return nil, err
	}

	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
