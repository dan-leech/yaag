package yaag

type Config struct {
	On bool `json:"on,omitempty" yaml:"On" toml:"On"`

	BaseUrls map[string]string `json:"baseUrls,omitempty" yaml:"BaseUrls" toml:"BaseUrls"`

	DocTitle string `json:"docTitle,omitempty" yaml:"DocTitle" toml:"DocTitle"`
	DocPath  string `json:"docPath,omitempty" yaml:"DocPath" toml:"DocPath"`
}
