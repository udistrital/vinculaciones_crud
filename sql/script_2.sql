CREATE SCHEMA vinculaciones_2;

CREATE SEQUENCE vinculaciones_2.tipo_vinculacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE vinculaciones_2.tipo_vinculacion(
	id integer NOT NULL DEFAULT nextval('vinculaciones_2.tipo_vinculacion_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_vinculacion PRIMARY KEY (id),
    CONSTRAINT unique_nombre_tipo_vinculacion UNIQUE (nombre)
);
COMMENT ON TABLE vinculaciones_2.tipo_vinculacion IS 'Tabla que parametriza los vinculaciones_2.';
COMMENT ON COLUMN vinculaciones_2.tipo_vinculacion.id IS 'Identificador unico de la tabla tipo_vinculacion.';
COMMENT ON COLUMN vinculaciones_2.tipo_vinculacion.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN vinculaciones_2.tipo_vinculacion.descripcion IS 'Campo en el que se puede registrar una descripcion de la información de tipo_vinculacion.';
COMMENT ON COLUMN vinculaciones_2.tipo_vinculacion.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN vinculaciones_2.tipo_vinculacion.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN vinculaciones_2.tipo_vinculacion.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN vinculaciones_2.tipo_vinculacion.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN vinculaciones_2.tipo_vinculacion.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';



CREATE SEQUENCE vinculaciones_2.cargos_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE vinculaciones_2.cargos(
	id integer NOT NULL DEFAULT nextval('vinculaciones_2.cargos_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_cargos PRIMARY KEY (id),
    CONSTRAINT unique_nombre_cargos UNIQUE (nombre)
);

COMMENT ON TABLE vinculaciones_2.cargos IS 'Tabla que parametriza los diferentes cargos.';
COMMENT ON COLUMN vinculaciones_2.cargos.id IS 'Identificador unico de la tabla cargos.';
COMMENT ON COLUMN vinculaciones_2.cargos.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN vinculaciones_2.cargos.descripcion IS 'Campo en el que se puede registrar una descripcion de la información de cargos.';
COMMENT ON COLUMN vinculaciones_2.cargos.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN vinculaciones_2.cargos.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN vinculaciones_2.cargos.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN vinculaciones_2.cargos.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN vinculaciones_2.cargos.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


CREATE SEQUENCE vinculaciones_2.vinculacion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE vinculaciones_2.vinculacion(
	id integer NOT NULL DEFAULT nextval('vinculaciones_2.vinculacion_id_seq'::regclass),
	tercero_principal_id integer NOT NULL,
    tercero_relacionado_id integer,
    tipo_vinculacion_id integer NOT NULL,
    fecha_inicio_vinculacion TIMESTAMP NOT NULL,
    fecha_fin_vinculacion TIMESTAMP,
    soporte INTEGER,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_vinculacion PRIMARY KEY (id)    
);

CREATE SEQUENCE vinculaciones_2.historico_cargos_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE vinculaciones_2.historico_cargos(
	id integer NOT NULL DEFAULT nextval('vinculaciones_2.historico_cargos_id_seq'::regclass),
	tercero_id integer,
    cargo_id integer NOT NULL,
    dependencia_id integer NOT NULL,
    fecha_inicio_vinculacion TIMESTAMP NOT NULL,
    fecha_fin_vinculacion TIMESTAMP,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_historico_cargo PRIMARY KEY (id)    
);



ALTER TABLE vinculaciones_2.vinculacion ADD CONSTRAINT fk_vinculacion_tipo_vinculacion FOREIGN KEY (tipo_vinculacion_id)
REFERENCES vinculaciones_2.tipo_vinculacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;

ALTER TABLE vinculaciones_2.historico_cargos ADD CONSTRAINT fk_vinculacion_cargo FOREIGN KEY (cargo_id)
REFERENCES vinculaciones_2.cargos (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;