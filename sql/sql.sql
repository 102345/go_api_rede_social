CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publicacoes;
DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS perfis;
DROP TABLE IF EXISTS perfisusuarios;

CREATE TABLE perfis(
    perfil_id int auto_increment primary key,
    descricao varchar(50) not null
)ENGINE=INNODB;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
)ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    primary key(usuario_id,seguidor_id)
)ENGINE=INNODB;

CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not  null,
    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    curtidas int default 0,
    criadaEm timestamp default current_timestamp()
)ENGINE=INNODB;


CREATE TABLE perfissusuarios(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,
    perfil_id int not null,
    FOREIGN KEY (perfil_id)
    REFERENCES perfis(perfil_id)
    ON DELETE CASCADE,
    primary key(usuario_id,perfil_id)
)ENGINE=INNODB;