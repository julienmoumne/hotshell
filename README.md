# Hotshell [![Build Status](https://travis-ci.org/julienmoumne/hotshell.svg?branch=master)](https://travis-ci.org/julienmoumne/hotshell)

> Interactive single keystroke menus for the shell

Hotshell is a command-line application to efficiently recall and share commands :

![demo](doc/demo.png)

Check out the [definition](https://github.com/julienmoumne/hotshell/blob/gh-pages/_includes/demo.hs.js)
and the [simulated demo](http://julienmoumne.github.io/hotshell/demos/demo.hs.js.html)
of this menu.

Quick tip : Name your definition file `hs.js` and `hs` will pick it up without having to specify a filename.

# Content

  - [Installation](#installation)
  - [How to define menus](#how-to-define-menus)
  - [Examples](./examples)
  - [Tips](./tips.md)
  - [Project Motives](#project-motives)
  - [Changelog](CHANGELOG.md)

## Installation

Hotshell is currently supported on Linux, OSX and Windows (10+).

### Precompiled packages

 - Linux & Windows ([WSL](https://msdn.microsoft.com/en-us/commandline/wsl/about)) : see [releases](https://github.com/julienmoumne/hotshell/releases)
 - OSX :
```bash
brew install julienmoumne/hotshell/hs
```

### Install from source

```bash
go get github.com/julienmoumne/hotshell/cmd/hs
```

## How to define menus
 
Menus are defined using a JavaScript DSL.

### Building blocks

> Menus are first defined with a top-level *menu item*

```javascript
item({desc: 'top level menu'}, function () {
    
})
```

> *Command items* associate bash commands to hot keys

```javascript
item({desc: 'top level menu'}, function () {
    
  item({key: 's', cmd: 'sudo service apache2 status'})      
  item({key: 'r', cmd: 'sudo service apache2 restart'})      
  
  // a description can be used to clarify the command intent
  item({key: 'a', desc: 'access logs', cmd: 'less +F /var/log/apache2/access.log'})
  item({key: 'e', desc: 'error logs', cmd: 'less +F /var/log/apache2/error.log'})
})
```

> *Submenus and Description items* can be used to add structure

```javascript
item({desc: 'top level menu'}, function () {
    
  item({desc: 'this is a description item, you can use it to describe the menu\n'})
  
  // a submenu to manage apache's daemon
  item({key: 'm', desc: 'apache management'}, function () {
    item({key: 'r', cmd: 'sudo service apache2 restart'})      
    item({key: 'h', cmd: 'sudo service apache2 stop'})
  })
  
  // a submenu to access apache's logs
  item({key: 'l', desc: 'apache logs'}, function () {
    item({key: 'a', desc: 'access', cmd: 'less +F /var/log/apache2/access.log'})
    item({key: 'e', desc: 'error', cmd: 'less +F /var/log/apache2/error.log'})
  })      
})
```

See a [simulated demo](http://julienmoumne.github.io/hotshell/demos/tutorial.hs.js.html) of the resulting menu.
 
### Tips & Examples

Check out [tips](./tips.md) and fully-featured [examples](./examples).

### JavaScript Runtime

The DSL is interpreted using Otto, an embeddable JavaScript interpreter.
See [Otto](https://github.com/robertkrimen/otto) and
[Otto Underscore](https://github.com/robertkrimen/otto/tree/master/underscore) 
for available JavaScript functions.

A custom function, `exec`, is provided to interact with the system when
defining menus.

`string = exec(string)` executes the specified command using `bash -c` and returns
the stdout if the command returned a non-zero exit code.

In case of failure, stderr is displayed in the menu without stopping the interpretation of the DSL.

See some [examples](./tips.md#exec).

For more information on how the DSL is interpreted see
[Building Trees using a JavaScript DSL](http://moumne.com/2016/07/30/building-trees-using-a-javascript-dsl).

## How to configure behavior

Configuration can be supplied in the same JavaScript DSL as menus using ```configure``` function:

```javascript
configure({
	breadcrumbType: 'vertical'
})
```

For now, there's only one parameter:

| Parameter Name       | Type   | Default            | Description |
| -------------------- | ------ | ------------------ | ----------- |
| ```breadcrumbType``` | string | ```"horizontal"``` | Possible breadcrumb type of display, one of ```"horizontal"``` or ```"vertical"``` |

## Project Motives

### Being more productive using the shell

As developers and system admins we sometimes spend too much time finding and typing commands on the shell.

We rely on a variety of methods to recall commands : 
search engines, reverse-search-history, aliases, auto-completion, ..

While they are effective, they require superfluous lookups and keystrokes.

Hotshell can be used to define menus containing often and not so often used commands :

```javascript
// file ~/.hs/hs.js  
item({desc: 'useful system commands'}, function () {
  item({key: 'f', desc: 'find text in files', cmd: 'echo -n "[location] [pattern] "; read l p; grep -rnws $l -e $p'})
  item({key: 'o', desc: 'check local port', cmd: 'echo -n "[port] "; read p; cat < /dev/tcp/127.0.0.1/$p'})
  // other useful commands..
})
```

### Sharing commands

When working in teams we usually go through common steps such as compiling and running tests.

Some of these commands may be available on your IDE, some may not. Team members may not always have the 
same IDEs.

Most of the time though, IDE task configurations are not committed into the VCS and are therefore not shared.

Hotshell proposes both a way to :

  - spread command-line skills in teams
  - produce executable documentation
  
As an example, checkout [Hotshell's hotshell](examples#hotshells-hotshell)