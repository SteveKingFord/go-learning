package config

import "fmt"

type Mysql struct {
	Server       string `structure:"server" json:"server" yaml:"server"`
	Config       string `structure:"config" json:"config" yaml:"config"`
	Database     string `structure:"database" json:"database" yaml:"database"`
	Username     string `structure:"username" json:"username" yaml:"username"`
	Password     string `structure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `structure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `structure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `structure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `structure:"log-zap" json:"logZap" yaml:"log-zap"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Server + ")/" + m.Database + "?" + m.Config
}

// PostgreSQL config
type PostgreSQL struct {
	Server       string `structure:"server" json:"server" yaml:"server"`
	Port         int    `structure:"port" json:"port" yaml:"port"`
	User         string `structure:"user" json:"user" yaml:"user"`
	Database     string `structure:"database" json:"database" yaml:"database"`
	Password     string `structure:"password" json:"password" yaml:"password"`
	Config       string `structure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `structure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `structure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `structure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `structure:"log-zap" json:"logZap" yaml:"log-zap"`
}

func (m *PostgreSQL) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s",
		m.Server, m.User, m.Password, m.Database, m.Port, m.Config)
}

// SQLServer config
type SQLServer struct {
	Server       string `structure:"server" json:"server" yaml:"server"`
	Database     string `structure:"database" json:"database" yaml:"database"`
	User         string `structure:"user" json:"user" yaml:"user"`
	Password     string `structure:"password" json:"password" yaml:"password"`
	MaxIdleConns int    `structure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `structure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `structure:"log-mode" json:"logMode" yaml:"log-mode"`
	LogZap       string `structure:"log-zap" json:"logZap" yaml:"log-zap"`
}

func (m *SQLServer) Dsn() string {
	return fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
		m.User, m.Password, m.Server, m.Database)
}
