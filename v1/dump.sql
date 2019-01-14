insert into os (platform, architecture) values ('Linux', 'x64');
insert into terminal (name, os) values ('bash', 1);
insert into commands (name, description, terminal) values 
('ls','lista arquivos e diretórios',1),
('clear','limpa a tela',1),
('Ctrl+L','Limpa a tela',1),
('env','lista as variáveis de ambiente',1),
('printenv','lista as variáveis de ambiente',1),
('echo $RANDOM','retorna um número randômico
',1),
('pwd','retorna o diretório atual',1),
('> arquivo.ext','este comando cria um arquivo',1),
('touch arquivo.ext','cria um arquivo',1),
('ps','lista os bash"s que estão rodando',1),
('bash','roda um shell dentro do shell',1),
('mkdir ','cria um diretório. É necessário informar o nome do diretório após o comando',1),
('ls -a','lista arquivos e diretórios ocultos',1),
('ls -l','lista os arquivos e diretórios e suas permissões',1),
('comando | less','coloca paginação caso as informações do comando ultrapassem o limite da janela do terminal. Melhora a visualização',1),
('alias','lista todos os ALIAS do sistema.',1),
('comando | more','melhora a leitura de um comando que retorna informações que excedam os limites da janela do terminal',1),
('echo "hello world!" >> arquivo.txt','Este comando escreve um texto em um arquivo passado como parâmetro',1),
('typeset -f','lista as funções criadas no bash.',1),
('rm -Rf ','Apaga arquivos e diretórios recursivamente. É necessário informar o nome do diretório após o comando',1),
('rm arquivo.txt','apaga um arquivo do sistema',1),
('rmdir','apaga um diretório do sistema, caso o diretório esteja vazio',1),
('ls -li','lista arquivos e diretórios com os seus indices',1),
('echo -n "Ola, seja bem vindo!"; date','O comando echo com -n executa o comando sem pular a linha do retorno do próximo comando',1),
('cal','retorna o calendário do mês atual',1),
('date','retorna a data',1),
('whoami','retorna o nome do usuário logado',1),
('date +%d','retorna apenas o dia',1),
('date +%d-%m','Retorna apenas o dia e o mês. O hífen para separar os números e opcional',1),
('date +%d-%m-%y','retorna o dia, mês e ano',1),
('date +%M','retorna o minuto atual',1),
('date +%Y','retorna o ano atual sem abreviar.',1),
('date +%S','retorna o tempo em segundos',1),
('date +%h','retorna o nome do mês atual',1),
('ls -la','lista todos arquivos e diretórios independente se for oculto ou não',1),
('echo "Hello world" > arquivo','sobreescreve o conteudo de um arquivo caso exista, com a mensagem',1),
('wc < arquivo.txt','este comando recebe um arquivo e retorna a quantidade de linhas, de palavras e bytes.',1);