package config

type Server struct {
	Zap Zap `structure:"zap" json:"zap" yaml:"zap"`
	System  System  `structure:"system" json:"system" yaml:"system"`
	// gorm
	Mysql Mysql `structure:"mysql" json:"mysql" yaml:"mysql"`
	PostgreSQL PostgreSQL `structure:"postgreSQL" json:"postgreSQL" yaml:"postgreSQL"`
	SQLServer SQLServer `structure:"SQLServer" json:"SQLServer" yaml:"SQLServer"`
}
