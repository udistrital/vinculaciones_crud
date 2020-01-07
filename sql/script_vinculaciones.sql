--
CREATE SCHEMA vinculaciones;

CREATE SEQUENCE vinculaciones.tipo_vinculacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE vinculaciones.tipo_vinculacion(
	id integer NOT NULL DEFAULT nextval('vinculaciones.tipo_vinculacion_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_tipo_vinculacion PRIMARY KEY (id),
    CONSTRAINT unique_nombre_tipo_vinculacion UNIQUE (nombre)
);

COMMENT ON TABLE vinculaciones.tipo_vinculacion IS 'Tabla que parametriza los tipos de vinculaciones, siendo una tipo de vinculación la relación que posee un tercero con la Universidad.';
COMMENT ON COLUMN vinculaciones.tipo_vinculacion.id IS 'Identificador unico de la tabla tipo_vinculacion.';
COMMENT ON COLUMN vinculaciones.tipo_vinculacion.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN vinculaciones.tipo_vinculacion.descripcion IS 'Campo en el que se puede registrar una descripcion de la información de tipo_vinculacion.';
COMMENT ON COLUMN vinculaciones.tipo_vinculacion.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN vinculaciones.tipo_vinculacion.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN vinculaciones.tipo_vinculacion.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN vinculaciones.tipo_vinculacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN vinculaciones.tipo_vinculacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


CREATE SEQUENCE vinculaciones.cargos_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE vinculaciones.cargos(
	id integer NOT NULL DEFAULT nextval('vinculaciones.cargos_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_cargos PRIMARY KEY (id),
    CONSTRAINT unique_nombre_cargos UNIQUE (nombre)
);

COMMENT ON TABLE vinculaciones.cargos IS 'Tabla que parametriza los diferentes cargos que puede tener un tercero dentro de la Universidad.';
COMMENT ON COLUMN vinculaciones.cargos.id IS 'Identificador unico de la tabla cargos.';
COMMENT ON COLUMN vinculaciones.cargos.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN vinculaciones.cargos.descripcion IS 'Campo en el que se puede registrar una descripcion de la información de cargos.';
COMMENT ON COLUMN vinculaciones.cargos.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN vinculaciones.cargos.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN vinculaciones.cargos.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN vinculaciones.cargos.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN vinculaciones.cargos.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


CREATE SEQUENCE vinculaciones.vinculacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE vinculaciones.vinculacion(
	id integer NOT NULL DEFAULT nextval('vinculaciones.vinculacion_id_seq'::regclass),
	tercero_principal_id integer NOT NULL,
    tercero_relacionado_id integer,
    tipo_vinculacion_id integer NOT NULL,
    cargo_id integer,
	dependencia_id integer,
	soporte integer,
	periodo_id integer,	
    fecha_inicio_vinculacion TIMESTAMP,
    fecha_fin_vinculacion TIMESTAMP,
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_vinculacion PRIMARY KEY (id)    
);

COMMENT ON TABLE vinculaciones.vinculacion IS 'Tabla que muestra el histórico de vinculaciones que tiene un tercero, ya sea con otro tercero o con un cargo dentro de la universidad.';
COMMENT ON COLUMN vinculaciones.vinculacion.id IS 'Identificador unico de la tabla vinculacion';
COMMENT ON COLUMN vinculaciones.vinculacion.tercero_principal_id IS 'Campo obligatorio que referencia al esquema terceros. Indica qué tercero tiene una vinculación.';
COMMENT ON COLUMN vinculaciones.vinculacion.tercero_relacionado_id IS 'Campo no obligatorio que referencia al esquema terceros. Es el tercero relacionado al tercero principal.';
COMMENT ON COLUMN vinculaciones.vinculacion.tipo_vinculacion_id  IS 'Campo obligatorio que referencia a la tabla tipo_vinculación. Indica el tipo de vinculación que posee el tercero principal.';
COMMENT ON COLUMN vinculaciones.vinculacion.cargo_id IS 'Campo no obligatorio que referencia a la tabla cargo. Indica, de aplicarse, qué cargo específico desempeña el tercero principal.';
COMMENT ON COLUMN vinculaciones.vinculacion.dependencia_id IS 'Campo no obligatorio que referencia a esquema oikos. Indica la dependencia en la que el tercero principal desempeña la vinculación.';
COMMENT ON COLUMN vinculaciones.vinculacion.soporte IS 'Campo no obligatorio que referencia a esquema documentos. Relaciona un documento que sirva de soporte para la vinculación indicada en el registro.';
COMMENT ON COLUMN vinculaciones.vinculacion.periodo_id IS 'Campo no obligatorio que referencia a esquema parametros_estandar (tabla unidad). Indica, de ser válido para la vinculación, el periodo al que corresponde.';
COMMENT ON COLUMN vinculaciones.vinculacion.fecha_inicio_vinculacion IS 'Campo no obligatorio que indica la fecha exacta del inicio de la vinculación del tercero. ';
COMMENT ON COLUMN vinculaciones.vinculacion.fecha_fin_vinculacion IS 'Campo no obligatorio que indica la fecha exacta en la que termina la vinculación del tercero. ';
COMMENT ON COLUMN vinculaciones.vinculacion.activo IS 'Campo obligatorio que indica el estado de la vinculación.';
COMMENT ON COLUMN vinculaciones.vinculacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN vinculaciones.vinculacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

ALTER TABLE vinculaciones.vinculacion ADD CONSTRAINT fk_vinculacion_tipo_vinculacion FOREIGN KEY (tipo_vinculacion_id)
REFERENCES vinculaciones.tipo_vinculacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE vinculaciones.vinculacion ADD CONSTRAINT fk_vinculacion_cargo FOREIGN KEY (cargo_id)
REFERENCES vinculaciones.cargos (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;