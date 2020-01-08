package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaVinculacion_20200107_122424 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaVinculacion_20200107_122424{}
	m.Created = "20200107_122424"

	migration.Register("CrearTablaVinculacion_20200107_122424", m)
}

// Run the migrations
func (m *CrearTablaVinculacion_20200107_122424) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE vinculaciones.vinculacion(id serial NOT NULL, tercero_principal_id integer NOT NULL, tercero_relacionado_id integer, tipo_vinculacion_id integer NOT NULL, cargo_id integer, dependencia_id integer, soporte integer, periodo_id integer, fecha_inicio_vinculacion TIMESTAMP, fecha_fin_vinculacion TIMESTAMP, activo boolean NOT NULL, fecha_creacion TIMESTAMP, fecha_modificacion TIMESTAMP, CONSTRAINT pk_vinculacion PRIMARY KEY (id) );")
	m.SQL("ALTER TABLE vinculaciones.vinculacion OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE vinculaciones.vinculacion IS 'Tabla que muestra el histórico de vinculaciones que tiene un tercero, ya sea con otro tercero o con un cargo dentro de la universidad.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.id IS 'Identificador unico de la tabla vinculacion';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.tercero_principal_id IS 'Campo obligatorio que referencia al esquema terceros. Indica qué tercero tiene una vinculación.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.tercero_relacionado_id IS 'Campo no obligatorio que referencia al esquema terceros. Es el tercero relacionado al tercero principal.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.tipo_vinculacion_id  IS 'Campo obligatorio que referencia a la tabla tipo_vinculación. Indica el tipo de vinculación que posee el tercero principal.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.cargo_id IS 'Campo no obligatorio que referencia a la tabla cargo. Indica, de aplicarse, qué cargo específico desempeña el tercero principal.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.dependencia_id IS 'Campo no obligatorio que referencia a esquema oikos. Indica la dependencia en la que el tercero principal desempeña la vinculación.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.soporte IS 'Campo no obligatorio que referencia a esquema documentos. Relaciona un documento que sirva de soporte para la vinculación indicada en el registro.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.periodo_id IS 'Campo no obligatorio que referencia a esquema parametros_estandar (tabla unidad). Indica, de ser válido para la vinculación, el periodo al que corresponde.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.fecha_inicio_vinculacion IS 'Campo no obligatorio que indica la fecha exacta del inicio de la vinculación del tercero. ';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.fecha_fin_vinculacion IS 'Campo no obligatorio que indica la fecha exacta en la que termina la vinculación del tercero. ';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.activo IS 'Campo obligatorio que indica el estado de la vinculación.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN vinculaciones.vinculacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CrearTablaVinculacion_20200107_122424) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS vinculaciones.vinculacion")
}
