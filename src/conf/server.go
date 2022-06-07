package conf

type Server struct {
	Web struct {
		Address  string `yaml:"address"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"web"`
}
