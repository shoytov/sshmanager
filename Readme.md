# SSH simple console manager

**Currently Version 0.1**

Allows you to save ssh connections (connection strings) to a remote server in the form: user@host. 
Each connection can be consciously named, and there is also a function for adding a description.
After the list of connections is formed, you can work with each element of this list in 2 modes:
- work directly in a command mode (a child process starts with ssh command)
- work in a file management mode (a child process starts with Midnight Commander in sh mode)

Therefore, ssh — OpenSSH remote login client, the Midnight Commander (mc) console file manager must be installed on your system.
In the future, it is planned to support other console file managers that support working with files via ssh (SFTP).

## Edit connections list manually
You can edit the file with the connections yourself without resorting to the SSH Manager interface.
The connection settings file is located in ```~/.config/sshmanager/config.json```

## Install
Make shure that you have been already installed:
- Midnight Commander (mc)
- ssh

Download the binary executable file for your platform from [Latest release](https://github.com/shoytov/sshmanager/releases/).   
Rename the binary file as you like. For example ```sshm```.
Move the binary file to the directory that is contained in the PATH variable.   
```shell
echo $PATH
```
Then run the program ```sshm```.

Alse you can build the binary themselves.
```
git clone https://github.com/dimidrol/sshmanager.git
cd sshmanager
go build .
```

## Use custom id rsa file
If you want to use a custom key file, add ssh connection to config ```~/.ssh/config```. For example:
```
Host somehost.com
    HostName somwhost.com
    IdentityFile ~/.ssh/id_rsa_somehost
```

Enjoy :)

## DEMO
![Demo gif](docs/demo.gif)
