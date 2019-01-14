use codeline;

select os.platform, os.architecture, 
terminal.name as terminal, 
commands.name as command, 
commands.description 
from commands
inner join terminal on terminal.id = commands.terminal
inner join os on os.id = terminal.os;
