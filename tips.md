# Tips

  - [Command-line usage](#command-line-usage)
  - [Menu and item definition](#menu-and-item-definition)
  - [Exec](#exec)

## Command-line usage

> Implicitly load the current directory's hotshell, `./hs.js`, or if not found, the system-wide hotshell `~/.hs/hs.js`

```bash
hs
```

> Specify the path to the definition file

```bash
hs -f ~/projects/web/hs.js
# or
hs -f ~/projects/web
```

> Load a menu remotely

```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hs/v0.1.0/hs.js
```

> Set the working directory to the location of the menu definition

```bash
hs --chdir -f ~/projects/web/hs.js
```

> Use aliases

```bash
# system-wide hotshell
alias hsh="hs -f ~/.hs/hs.js"

# generic hotshells, docker :
alias hsd="hs -f https://raw.githubusercontent.com/julienmoumne/hs/v0.1.0/examples/docker/docker.hs.js"
# docker compose
alias hsdc="hs -f https://raw.githubusercontent.com/julienmoumne/hs/v0.1.0/examples/docker/docker-compose.hs.js"
# docker machine
alias hsdm="hs -f https://raw.githubusercontent.com/julienmoumne/hs/v0.1.0/examples/docker/docker-machine.hs.js"
# vagrant
alias hsv="hs -f https://raw.githubusercontent.com/julienmoumne/hs/v0.1.0/examples/vagrant/vagrant.hs.js"
```

> Generate an interactive HTML demo of your menus, [example](https://julienmoumne.github.com/hs/demos/hs.js.html)

```bash
hs --generate-demo -f ~/projects/web/hs.js > hotshell-web-demo.html  
```

## Menu and item definition
  
> Prompt for inputs to be provided as command arguments

```javascript
// prompts for a location and a pattern and triggers a grep search
item({
  desc: 'find text in files',
  cmd: 'echo -n "[location] [pattern] " ' + // prompt for inputs
         '&& read l p ' + // read inputs
         '&& grep -rnws $l -e $p' // provide inputs as command arguments
})
```

> Access other interactive applications

```javascript
item({cmd: 'ssh remote-server'})
item({cmd: 'sudo vim /etc/hosts'})
```

> When running out of characters for defining hot keys, use capital letters

```javascript
item({key: 'S', cmd: 'ssh remote-server'})
```

> Include menus defined in separate files

See [includes example](examples#includes)

> There is a good number of command examples

In the default hotshell `hs --default` and in the [examples directory](./examples).

> The DSL defined by Hotshell uses some JavaScript tricks, learn more about it : http://alexyoung.org/2009/10/22/javascript-dsl/

## Exec

  > Retrieve environment variables
  
```javascript
httpPort = exec('echo $HTTP_PORT'); if (httpPort == '') throw 'please set $HTTP_PORT'
item({key: 's', desc: 'start http server', cmd: 'python -m SimpleHTTPServer ' + httpPort})
```

  > Conditionally set-up items
  
```javascript
linux = exec('uname').indexOf('Linux') > -1
item({key: 'u', desc: 'update', cmd: linux ? 'sudo apt-get update' : 'brew update'})
```

  > Dynamically create menus
  
```javascript
recentlyUpdatedLogs = exec('ls -dt /var/log/*.* | head -n 3').split('\n')
_.each(recentlyUpdatedLogs, function(el, ix) {
  item({key: ix, desc: 'less ' + el, cmd: 'less +F ' + el})
})
```
![Generated Items - Logs](doc/generated-items-logs.png)
