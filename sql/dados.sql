insert into perfis(descricao) values('administrador');
insert into perfis(descricao)  values('moderador');
insert into perfis(descricao)  values('membro');

-- senha 123456
insert into usuarios(nome,nick,email,senha)
values('Usuario 1','usuario_1','usuario1@gmail.com','$2a$10$v0.f77VUVqVJ/X6nIiYhZeR0YszFEYUXxcGCpQrKnd.e//Z/OeNBC');
insert into usuarios(nome,nick,email,senha)
values('Usuario 2','usuario_2','usuario2@gmail.com','$2a$10$v0.f77VUVqVJ/X6nIiYhZeR0YszFEYUXxcGCpQrKnd.e//Z/OeNBC');
insert into usuarios(nome,nick,email,senha)
values('Usuario 3','usuario_3','usuario3@gmail.com','$2a$10$v0.f77VUVqVJ/X6nIiYhZeR0YszFEYUXxcGCpQrKnd.e//Z/OeNBC');

insert into seguidores(usuario_id,seguidor_id)values(1,2);
insert into seguidores(usuario_id,seguidor_id)values(3,1);
insert into seguidores(usuario_id,seguidor_id)values(1,3);

insert into publicacoes(titulo,conteudo,autor_id)
values
("Publicação do Usuário 1","Essa é a publicação do usuario 1! Oba!",1),
("Publicação do Usuário 2","Essa é a publicação do usuario 2! Oba!",2),
("Publicação do Usuário 3","Essa é a publicação do usuario 3! Oba!",3);

insert into devbook.perfissusuarios(usuario_id,perfil_id)values(1,1);
insert into devbook.perfissusuarios(usuario_id,perfil_id)values(2,2);