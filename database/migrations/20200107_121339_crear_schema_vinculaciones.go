package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchemaVinculaciones_20200107_121339 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchemaVinculaciones_20200107_121339{}
	m.Created = "20200107_121339"

	migration.Register("CrearSchemaVinculaciones_20200107_121339", m)
}

// Run the migrations
func (m *CrearSchemaVinculaciones_20200107_121339) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE SCHEMA IF NOT EXISTS vinculaciones;")
	m.SQL("ALTER SCHEMA vinculaciones OWNER TO desarrollooas;")
	m.SQL("SET search_path TO pg_catalog,public,vinculaciones;")
}

// Reverse the migrations
func (m *CrearSchemaVinculaciones_20200107_121339) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP SCHEMA IF EXISTS vinculaciones");
}
