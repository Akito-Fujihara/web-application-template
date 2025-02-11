package main

import (
	"github.com/Akito-Fujihara/web-application-template/app/config/env"
	"github.com/Akito-Fujihara/web-application-template/app/infra/mysql"
	"gorm.io/gen"
)

func main() {
	mysqlConfig, err := env.MysqlConfig()
	if err != nil {
		panic(err)
	}
	DB, cleanup, err := mysql.NewMysqlConn(mysqlConfig)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	g := gen.NewGenerator(gen.Config{
		OutPath: "app/infra/mysql/ormgen",
		ModelPkgPath: "app/infra/mysql/dbschema",
		
		FieldNullable: 	 true,
		FieldCoverable:  true, // ここは要検討
		FieldSignable:   false,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	
		Mode: gen.WithDefaultQuery,
	})

	g.UseDB(DB)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}