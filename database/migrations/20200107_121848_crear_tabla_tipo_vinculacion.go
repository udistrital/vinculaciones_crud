package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaTipoVinculacion_20200107_121848 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaTipoVinculacion_20200107_121848{}
	m.Created = "20200107_121848"

	migration.Register("CrearTablaTipoVinculacion_20200107_121848", m)
}

// Run the migrations
func (m *CrearTablaTipoVinculacion_20200107_121848) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE vinculaciones.tipo_vinculacion( id integer NOT NULL DEFAULT nextval('vinculaciones.tipo_vinculacion_id_seq'::regclass), nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP, fecha_modificacion TIMESTAMP, CONSTRAINT pk_tipo_vinculacion PRIMARY KEY (id), CONSTRAINT unique_nombre_tipo_vinculacion UNIQUE (nombre) );")
	m.SQL("ALTER TABLE vinculaciones.tipo_vinculacion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE vinculaciones.tipo_vinculacion IS 'Tabla que parametriza los tipos de vinculaciones, siendo una tipo de vinculación la relación que posee un tercero con la Universidad.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.tipo_vinculacion.id IS 'Identificador unico de la tabla tipo_vinculacion.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.tipo_vinculacion.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.tipo_vinculacion.descripcion IS 'Campo en el que se puede registrar una descripcion de la información de tipo_vinculacion.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.tipo_vinculacion.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.tipo_vinculacion.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.tipo_vinculacion.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.tipo_vinculacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.tipo_vinculacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaTipoVinculacion_20200107_121848) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS vinculaciones.tipo_vinculacion")
}
