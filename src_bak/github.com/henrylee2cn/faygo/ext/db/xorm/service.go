package xorm

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	 _ "github.com/denisenkom/go-mssqldb" //mssql
	 _ "github.com/go-sql-driver/mysql" //mysql
	// _ "github.com/lib/pq"              //postgres
	// _ "github.com/mattn/go-oci8"         //oracle(need to install the pkg-config utility)
	// _ "github.com/mattn/go-sqlite3"      //sqlite

	"github.com/henrylee2cn/faygo"
	"github.com/henrylee2cn/faygo/utils"
	"bytes"
)

// DBService is a database engine object.
type DBService struct {
	Default *xorm.Engine            // the default database engine
	List    map[string]*xorm.Engine // database engine list
}

var dbService = func() (serv *DBService) {
	serv = &DBService{
		List: map[string]*xorm.Engine{},
	}
	var errs []string
	defer func() {
		if len(errs) > 0 {
			panic("[xorm] " + strings.Join(errs, "\n"))
		}
		if serv.Default == nil {
			faygo.Panicf("[xorm] the `default` database engine must be configured and enabled")
		}
	}()

	err := loadDBConfig()
	if err != nil {
		faygo.Panicf("[xorm]", err.Error())
		return
	}

	for _, conf := range dbConfigs {
		if !conf.Enable {
			continue
		}
		if conf.Driver == "mssql" {
			conf.Connstring = formatConnString(conf.Connstring)
		}
		engine, err := xorm.NewEngine(conf.Driver, conf.Connstring)
		if err != nil {
			faygo.Critical("[xorm]", err.Error())
			errs = append(errs, err.Error())
			continue
		}
		err = nil //engine.Ping()
		if err != nil {
			faygo.Critical("[xorm]", err.Error())
			errs = append(errs, err.Error())
			continue
		}
		engine.SetLogger(iLogger)
		engine.SetMaxOpenConns(conf.MaxOpenConns)
		engine.SetMaxIdleConns(conf.MaxIdleConns)
		engine.SetDisableGlobalCache(conf.DisableCache)
		engine.ShowSQL(conf.ShowSql)
		engine.ShowExecTime(conf.ShowExecTime)

		if (conf.TableFix == "prefix" || conf.TableFix == "suffix") && len(conf.TableSpace) > 0 {
			var impr core.IMapper
			if conf.TableSnake {
				impr = core.SnakeMapper{}
			} else {
				impr = core.SameMapper{}
			}
			if conf.TableFix == "prefix" {
				engine.SetTableMapper(core.NewPrefixMapper(impr, conf.TableSpace))
			} else {
				engine.SetTableMapper(core.NewSuffixMapper(impr, conf.TableSpace))
			}
		}

		if (conf.ColumnFix == "prefix" || conf.ColumnFix == "suffix") && len(conf.ColumnSpace) > 0 {
			var impr core.IMapper
			if conf.ColumnSnake {
				impr = core.SnakeMapper{}
			} else {
				impr = core.SameMapper{}
			}
			if conf.ColumnFix == "prefix" {
				engine.SetTableMapper(core.NewPrefixMapper(impr, conf.ColumnSpace))
			} else {
				engine.SetTableMapper(core.NewSuffixMapper(impr, conf.ColumnSpace))
			}
		}

		if conf.Driver == "sqlite3" && !utils.FileExists(conf.Connstring) {
			os.MkdirAll(filepath.Dir(conf.Connstring), 0777)
			f, err := os.Create(conf.Connstring)
			if err != nil {
				faygo.Critical("[xorm]", err.Error())
				errs = append(errs, err.Error())
			} else {
				f.Close()
			}
		}

		serv.List[conf.Name] = engine
		if DEFAULTDB_NAME == conf.Name {
			serv.Default = engine
		}
	}
	return
}()

func formatConnString(conn string) string {
	result := conn
	mssql_symbol := "sqlserver://"
	if strings.Index(conn, mssql_symbol) == 0 { //uri's pattern
		var ret_buffer bytes.Buffer
		uri := strings.Split(conn, "//")[1]
		uri_tmp := strings.Split(uri, "@")
		if uri_tmp != nil {
			user_pwd := strings.Split(uri_tmp[0], ":")
			if user_pwd != nil {
				ret_buffer.WriteString("user id")
				ret_buffer.WriteString("=")
				ret_buffer.WriteString(user_pwd[0])
				ret_buffer.WriteString(";")
				ret_buffer.WriteString("password")
				ret_buffer.WriteString("=")
				ret_buffer.WriteString(user_pwd[1])
				ret_buffer.WriteString(";")
			}
			uri_tmp = strings.Split(uri_tmp[1], "?")
			if uri_tmp != nil {
				host_port_arr := strings.Split(uri_tmp[0], ":")
				if host_port_arr != nil {
					ret_buffer.WriteString("server")
					ret_buffer.WriteString("=")
					ret_buffer.WriteString(host_port_arr[0])
					ret_buffer.WriteString(";")
					ret_buffer.WriteString("port")
					ret_buffer.WriteString("=")
					ret_buffer.WriteString(host_port_arr[1])
					ret_buffer.WriteString(";")
				}
				uri_params := strings.Split(uri_tmp[1], "&")
				if uri_params != nil {
					for  _, value := range uri_params {
						uri_param_kv := strings.Split(value, "=")
						ret_buffer.WriteString(uri_param_kv[0])
						ret_buffer.WriteString("=")
						ret_buffer.WriteString(uri_param_kv[1])
						ret_buffer.WriteString(";")
					}
				}
			}
		}
		result = ret_buffer.String()
	}
	return result
}