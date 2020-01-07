package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaCargos_20200107_122055 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaCargos_20200107_122055{}
	m.Created = "20200107_122055"

	migration.Register("CrearTablaCargos_20200107_122055", m)
}

// Run the migrations
func (m *CrearTablaCargos_20200107_122055) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE vinculaciones.cargos( id integer NOT NULL DEFAULT nextval('vinculaciones.cargos_id_seq'::regclass), nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP, fecha_modificacion TIMESTAMP, CONSTRAINT pk_cargos PRIMARY KEY (id), CONSTRAINT unique_nombre_cargos UNIQUE (nombre) );")
	m.SQL("ALTER TABLE vinculaciones.cargos OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE vinculaciones.cargos IS 'Tabla que parametriza los diferentes cargos que puede tener un tercero dentro de la Universidad.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.cargos.id IS 'Identificador unico de la tabla cargos.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.cargos.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.cargos.descripcion IS 'Campo en el que se puede registrar una descripcion de la información de cargos.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.cargos.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.cargos.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.cargos.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.cargos.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.cargos.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaCargos_20200107_122055) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS vinculaciones.cargos")
}
