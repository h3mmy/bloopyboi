package models

type ClientGenerator struct {
	Name         string
	ProviderName string
	Args         map[string]interface{}
	Verbose      bool
}

func (s *ClientGenerator) GetArgs() map[string]interface{} {
	return s.Args
}
func (s *ClientGenerator) SetArgs(args map[string]interface{}) {
	s.Args = args
}

func (s *ClientGenerator) SetProviderName(providerName string) {
	s.ProviderName = providerName
}

func (s *ClientGenerator) GetProviderName() string {
	return s.ProviderName
}

func (s *ClientGenerator) SetVerbose(verbose bool) {
	s.Verbose = verbose
}

func (s *ClientGenerator) SetName(name string) {
	s.Name = name
}
func (s *ClientGenerator) GetName() string {
	return s.Name
}
