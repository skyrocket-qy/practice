package main

import (
	"time"

	"test/query"

	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(
		sqlite.Open(
			"./sqlite.db",
		),
	)
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	g.ApplyBasic(
		// Generate struct `User` based on table `users`
		g.GenerateModel("Accounts"),
		// g.GenerateModel("Roles"),
		// g.GenerateModel("Stores"),

		// Generate struct `Employee` based on table `users`
		// g.GenerateModelAs("users", "Employee"),

		// Generate struct `User` based on table `users` and generating options
		// g.GenerateModel(
		// 	"users",
		// 	gen.FieldIgnore("address"),
		// 	gen.FieldType("id", "int64"),
		// ),

		// Generate struct `Customer` based on table `customer` and generating options
		// customer table may have a tags column, it can be JSON type, gorm/gen tool can generate for your JSON data type
		// g.GenerateModel("customer", gen.FieldType("tags", "datatypes.JSON")),
	)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}

func test() {
	query.Account.Where(query.Account.CreatedAt.Eq(time.Now())).Find()
}
