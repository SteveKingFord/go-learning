package config

type System struct {
	Env           string `structure:"env" json:"env" yaml:"env"`
	Addr          int    `structure:"addr" json:"addr" yaml:"addr"`
	DbType        string `structure:"db-type" json:"dbType" yaml:"db-type"`
	OssType       string `structure:"oss-type" json:"ossType" yaml:"oss-type"`
	UseMultipoint bool   `structure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`

}
