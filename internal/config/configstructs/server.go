package configstructs

type Server struct {
	Listen int  `yaml:"listen"`
	Debug  bool `yaml:"debug"`
}
