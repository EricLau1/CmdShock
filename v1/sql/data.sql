use codeline;

insert into os (platform, architecture) values ('Linux', 'x64');
insert into terminal(name, os) values ('bash', 1);
insert into commands(name, description, terminal) values
('ls', 'lista arquivos e diretórios', 1);

