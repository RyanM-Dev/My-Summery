| (pipe): redirect output of one command to input of another command
cat filename/dirname | less  :to show them by less page by page
less: 
	previuos page key b
	next page key space
grep: globally search for regular expression and print out

history | grep sudo
history | grep "sudo chmod"  : when more than one word
cat filename | grep port

redirect: 
ridirect and overwrite to file   :   >  exp: history | grep sudo > sudo-commands.txt
ridirect and append to file : >>  :    exp: history | grep chmod >> sudo-commands.txt
 >  