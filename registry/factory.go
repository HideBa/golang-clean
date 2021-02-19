package registry

var FactorySingleton *Factory

type Factory struct {
	Envs  *Envs
	cache map[string]interface{}
}

func ClearFactory() {
	FactorySingleton = nil
}

func GetFactory() *Factory {
	if FactorySingleton == nil {
		FactorySingleton = &Factory{
			Envs: NewEnvs(),
		}
	}
	return FactorySingleton
}
