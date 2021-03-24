package initialize

import (
	"fmt"
	"kingford-backend/global"
	"kingford-backend/initialize/internal"
	"kingford-backend/model"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//@author: SliverHorn
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB
func Gorm() *gorm.DB {
	switch global.Config.System.DbType {
	case "PostgreSQL":
		return GormPostgreSQL()
	case "SQLServer":
		return GormSQLServer()
	default:
		return GormMysql()
	}
}

// MysqlTables
//@author: SliverHorn
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB
func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},

	)
	if err != nil {
		global.Log.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.Log.Info("register table success")
}

//@author: SliverHorn
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB
func GormMysql() *gorm.DB {
	m := global.Config.Mysql
	if m.Database == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Server + ")/" + m.Database + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.Log.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@description: 初始化PostgreSQL数据库
func GormPostgreSQL() *gorm.DB {
	m := global.Config.PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d %s",
		m.Server, m.User, m.Password, m.Database, m.Port, m.Config)

	postgreConfig := postgres.Config{
		DSN: dsn,
	}

	if db, err := gorm.Open(postgres.New(postgreConfig), gormConfig(m.LogMode)); err != nil {
		global.Log.Error("PostgreSQL启动异常", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@description: 初始化SQLServer数据库
func GormSQLServer() *gorm.DB {
	m := global.Config.SQLServer
	// "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
		m.User, m.Password, m.Server, m.Database)

	sqlserverConfig := sqlserver.Config{
		DSN: dsn,
	}

	if db, err := gorm.Open(sqlserver.New(sqlserverConfig), gormConfig(m.LogMode)); err != nil {
		global.Log.Error("SQLServer启动异常", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@author: SliverHorn
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.Config.Mysql.LogZap {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = internal.Default.LogMode(logger.Info)
			break
		}
		config.Logger = internal.Default.LogMode(logger.Silent)
	}
	return config
}
