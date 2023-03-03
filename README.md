**Goilerplate** is a simple to use boilerplate generator for Go projects. It is a command line tool that allows you to create a new project from a template.

## Installation
```bash
git clone https://github.com/SolarWolf-Code/GoilerPlate && mkdir -p ~/bin && cp -r GoilerPlate ~/bin/ && chmod +x ~/bin/GoilerPlate/gp && export PATH=$PATH:~/bin/GoilerPlate && echo "export PATH=$PATH:~/bin/GoilerPlate" >> ~/.bashrc && source ~/.bashrc && rm -rf GoilerPlate
```

## Usage
```bash
goilerplate -n <project name> -p <project path> 
```


### Example
```bash
goilerplate -n myproject -p ~/Desktop
```
The result of the command will be a new project with the following structure:
```
~/Desktop/myproject
├── main.go
├── go.mod
```
## Template
To edit your boilerplate that is written to your projects, simply navigate to the file ~/bin/GoilerPlate/template.go and change the contents to your liking.