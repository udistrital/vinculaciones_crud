package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearLlavesForaneas_20200107_122641 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearLlavesForaneas_20200107_122641{}
	m.Created = "20200107_122641"

	migration.Register("CrearLlavesForaneas_20200107_122641", m)
}

// Run the migrations
func (m *CrearLlavesForaneas_20200107_122641) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE vinculaciones.vinculacion ADD CONSTRAINT fk_vinculacion_tipo_vinculacion FOREIGN KEY (tipo_vinculacion_id) REFERENCES vinculaciones.tipo_vinculacion (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")
	m.SQL("ALTER TABLE vinculaciones.vinculacion ADD CONSTRAINT fk_vinculacion_cargo FOREIGN KEY (cargo_id) REFERENCES vinculaciones.cargos (id) MATCH FULL ON DELETE RESTRICT ON UPDATE CASCADE;")
}

// Reverse the migrations
func (m *CrearLlavesForaneas_20200107_122641) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE vinculaciones.vinculacion DROP CONSTRAINT IF EXISTS fk_vinculacion_tipo_vinculacion CASCADE;")
	m.SQL("ALTER TABLE vinculaciones.vinculacion DROP CONSTRAINT IF EXISTS fk_vinculacion_cargo CASCADE;")
}
