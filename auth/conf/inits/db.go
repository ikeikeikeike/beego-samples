package inits

import (
	"time"

	_ "github.com/ikeikeikeike/beego-samples/auth/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func init() {

	dbname := "default"
	runmode := beego.AppConfig.String("runmode")
	datasource := beego.AppConfig.String("datasource")

	switch runmode {
	case "prod":
		orm.RegisterDriver("postgres", orm.DR_Postgres)
		orm.RegisterDataBase(dbname, "postgres", datasource, 30)
		orm.SetMaxIdleConns(dbname, 100)
		orm.SetMaxOpenConns(dbname, 100)
	case "dev":
		orm.Debug = true
		fallthrough
	default:
		orm.RegisterDriver("sqlite3", orm.DR_Sqlite)
		orm.RegisterDataBase(dbname, "sqlite3", datasource)
	}

	orm.DefaultTimeLoc = time.FixedZone("Asia/Tokyo", 9*60*60)

	force, verbose := false, true
	err := orm.RunSyncdb(dbname, force, verbose)
	if err != nil {
		panic(err)
	}

	// orm.RunCommand()
}
