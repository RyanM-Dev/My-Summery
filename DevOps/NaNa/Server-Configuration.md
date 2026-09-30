you don't work with root user, you should have user for each service and even for server administration we shouldn't use root user.

the user who config and need sudo access add it to sudo group
when using sudo commands in the user added to sudo group, it ask you the user password not the root pass!

you need to add ssh public key in the user .ssh/authorized_keys to be able to ssh to server using ssh key
