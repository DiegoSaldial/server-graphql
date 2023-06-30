create table `usuarios`(
    `id` integer not null auto_increment primary key,
    `nombres` varchar(30) not null,
    `apellidos` varchar(30) not null,
    `username` varchar(25) not null unique,
    `password` varchar(100) not null,
    `foto_url` varchar(200),
    `telefono` varchar(20),
    `correo` varchar(100),
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    `estado` tinyint(1) not null default 1
);

create table `roles`(
    `id` integer unsigned not null auto_increment primary key,
    `nombre` varchar(30) not null unique
    -- `rol_bit` tinyint(1) unsigned not null unique comment "debe ser una potencia de 2 si o si"
);

create table `rol_permiso`(
    `metodo` varchar(30) not null,
    `rol_bits` integer unsigned not null,
    primary key(`metodo`,`rol_bits`)
);

create table `rol_usuario`(
    `usuario_id` integer not null,
    `rol_id` integer unsigned not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00'),
    foreign key(`usuario_id`) references `usuarios`(`id`),
    foreign key(`rol_id`) references `roles`(`id`),
    primary key(`usuario_id`, `rol_id`)
);

create table `tokens`(
    `username` varchar(25) not null primary key,
    `token` varchar(500) not null,
    `registrado` datetime not null default CONVERT_TZ(NOW(), @@session.time_zone, '-4:00') ON UPDATE CURRENT_TIMESTAMP
);


insert into usuarios(nombres,apellidos,username,password) values('Usuario','Administrador','admin','admin');
insert into usuarios(nombres,apellidos,username,password) values('Usuario','Taxista','taxista','taxista');
insert into usuarios(nombres,apellidos,username,password) values('Usuario','Pasajero','pasajero','pasajero');

insert into roles(nombre) values('administrador');
insert into roles(nombre) values('conductor');
insert into roles(nombre) values('pasajero');

insert into rol_usuario(usuario_id,rol_id) values(1,1);
insert into rol_usuario(usuario_id,rol_id) values(2,2);
insert into rol_usuario(usuario_id,rol_id) values(3,3);
insert into rol_usuario(usuario_id,rol_id) values(1,3);

-- administrador = 1
-- conductor = 2
-- pasajero = 4
-- otros = 8
-- se asigna la suma de valores de roles permitidos
-- valorAsignar = 1 << (rol_id-1) 
-- se verifica: rol_bits & 4 == 4 para el rol 4
insert into rol_permiso(metodo,rol_bits) values
("usuarioByUsername",   7),
("roles",               7),  
("createUsuario",       1), 
("updateUsuario",       1), 
("createRol",           1),
("deleteRol",           1),
("modificarRol",           1),
("crearNuevoPermiso",           1),
("actualizarPermiso",           1),
("eliminarPermiso",           1);


-- select * from rol_permiso;

-- delete from roles where nombre = "xadministrador";

-- select * from roles;


-- UPDATE rol_permiso
-- SET rol_bits = 
--     CASE 
--         WHEN metodo = 'roles' THEN 1
--         WHEN metodo = 'metodo2' THEN 'nuevo_rol_bits2'
--         WHEN metodo = 'metodo3' THEN 'nuevo_rol_bits3'
--     END
-- WHERE metodo IN ('roles', 'metodo2', 'metodo3');