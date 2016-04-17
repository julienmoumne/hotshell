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
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/hs.js
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
alias hsd="hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/examples/docker/docker.hs.js"
# docker compose
alias hsdc="hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/examples/docker/docker-compose.hs.js"
# docker machine
alias hsdm="hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/examples/docker/docker-machine.hs.js"
# vagrant
alias hsv="hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/examples/vagrant/vagrant.hs.js"
```

> Generate an interactive HTML demo of your menus, [example](https://julienmoumne.github.com/hotshell/demos/hs.js.html)

```bash
hs --generate-demo -f ~/projects/web/hs.js > hotshell-web-demo.html  
```

## Menu and item definition

> Output the definition object to help debugging :

```javascript
item({desc: 'debug'}, function() {
  item({key: 'r', desc: 'restart apache', cmd: 'sudo service apache2 restart'})
  item({key: 'a', desc: 'access.log', cmd: 'less +F /var/log/apache2/access.log'})
}) 

// 'items' contains the whole definition
console.log(JSON.stringify(items, null, ' '))
```
displays
```javascript
[
 {
  "cmd": null,
  "desc": "debug",
  "items": [
   {
    "cmd": "sudo service apache2 restart",
    "desc": "restart apache",
    "items": [],
    "key": "r"
   },
   {
    "cmd": "less +F /var/log/apache2/access.log",
    "desc": "access.log",
    "items": [],
    "key": "a"
   }
  ],
  "key": null
 }
]
```
  
> Commands can receive inputs from the user with bash builtin [read](http://wiki.bash-hackers.org/commands/builtin/read) 

```javascript
item({key: 'p', desc: 'check local port', cmd:
    'echo -n "[port] "; ' + // prompt for port number
    'read p; ' + // read port number and assign it to variable 'p'
    'cat < /dev/tcp/127.0.0.1/$p' // check if port 'p' is opened locally
})

item({key: 'f', desc: 'find text in files', cmd:
    'echo -n "[location] [pattern] "; ' + // prompt for location and pattern
    'read l p; ' + // read location and pattern into variables 'l' and 'p'
    'grep -rnws $l -e $p' // search for pattern 'p' traversing files rooted at 'l'
})
```

> Enter other interactive applications

```javascript
item({key: 's', cmd: 'ssh remote-server'})
item({key: 'h', cmd: 'sudo vim /etc/hosts'})
```

> Writing simple scripts is a possibility
 
```javascript
item({key: 't', desc: 'test brew formula' + '\n  ', cmd: script(
    'set -eu',
    'echo -n "[git hash] "',
    'read githash',
    'formula=https://raw.githubusercontent.com/julienmoumne/homebrew/$githash/Library/Formula/hs.rb',
    'brew remove hs || true',
    'brew cleanup -s',
    'brew install $formula',
    'brew test $formula',
    'hs'
)})

function script () {
    src = '';
    _(arguments).each(function (el, ix) { src += '   ' + el + '\n' })
    return src
}
```

> When running out of characters for defining hot keys

```javascript
// use capital letters
item({key: 'S', cmd: 'ssh remote-server'})

// or group commands in submenus
item({key: 'g', desc: 'group of related commands'}, function() {

  // the complete alphabet is available
  item({key: 'a', cmd: 'echo a'})
  item({key: 'b', cmd: 'echo b'})
})
```

> Include menus defined in separate files

See [includes example](examples#includes)

> There is a good number of command examples in

 - the default hotshell `hs --default`
 - the [examples directory](./examples)

> The DSL defined by Hotshell uses some JavaScript tricks

see [Fear and Loathing in JavaScript DSLs](http://alexyoung.org/2009/10/22/javascript-dsl/)

## Exec

> 'console.log()' can be used to help debugging

```javascript
console.log(exec('echo $(date)'))
```

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
_(recentlyUpdatedLogs).each(function(el, ix) {
  item({key: ix, desc: 'less ' + el, cmd: 'less +F ' + el})
})
```
![Generated Items - Logs](doc/generated-items-logs.png)
