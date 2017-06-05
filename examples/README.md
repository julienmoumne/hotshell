# Examples of menus

  - [Hotshell's hotshell](#hotshells-hotshell) - develop, build and test Hotshell
  - [Hotshell's default menu](#hotshells-default-menu) - loaded when no definition file is found
  - [Docker, Docker Compose & Docker Machine](#docker-docker-compose--docker-machine) - useful commands when working with Docker
  - [Vagrant](#vagrant) - useful commands when working with Vagrant
  - [Topten](#topten) - 10 most used commands
  - [Network](#network) - factor out groups of commands
  - [Modules](#modules) - modularize menu definitions
  - [Nested Hotshells](#nested-hotshells) - call Hotshell within Hotshell

## Hotshell's hotshell

Used to develop, build and test Hotshell.

  - [demo](http://julienmoumne.github.io/hotshell/demos/hs.js.html)
  - [generated markdown](../COMMANDS.md)
  - [source](../hs.js)
  - try it :
```bash
H='raw.githubusercontent.com/julienmoumne/hotshell/master/hs.js'; T=`mktemp`; wget $H -qO $T && hs -f $T; rm $T
```

## Hotshell's default menu

This menu is loaded when no definition file is found.

  - [demo](http://julienmoumne.github.io/hotshell/demos/default.hs.js.html)
  - [generated markdown](./default/default.hs.js.md)
  - [source](./default/default.hs.js)
  - try it :
```bash
H='raw.githubusercontent.com/julienmoumne/hotshell/master/examples/default/default.hs.js'; T=`mktemp`; wget $H -qO $T && hs -f $T; rm $T
```

## Docker, Docker Compose & Docker Machine

Menus with useful Docker commands.

  - demos
    * [docker](http://julienmoumne.github.io/hotshell/demos/docker.hs.js.html)
    * [docker-compose](http://julienmoumne.github.io/hotshell/demos/docker-compose.hs.js.html)
    * [docker-machine](http://julienmoumne.github.io/hotshell/demos/docker-machine.hs.js.html)
  - generated markdown
    * [docker](./docker/docker.hs.js.md)
    * [docker-compose](./docker/docker-compose.hs.js.md)
    * [docker-machine](./docker/docker-machine.hs.js.md)
  - sources
    * [docker](./docker/docker.hs.js)
    * [docker-compose](./docker/docker-compose.hs.js)
    * [docker-machine](./docker/docker-machine.hs.js)
  - try it :
```bash
# docker
H='raw.githubusercontent.com/julienmoumne/hotshell/master/examples/docker/docker.hs.js'; T=`mktemp`; wget $H -qO $T && hs -f $T; rm $T

# docker-machine
H='raw.githubusercontent.com/julienmoumne/hotshell/master/examples/docker/docker-machine.hs.js'; T=`mktemp`; wget $H -qO $T && hs -f $T; rm $T

# docker-compose
cd $DOCKER_COMPOSE_PROJECT
H='raw.githubusercontent.com/julienmoumne/hotshell/master/examples/docker/docker-compose.hs.js'; T=`mktemp`; wget $H -qO $T && hs -f $T; rm $T
```
  - integrate it in your own menus using [hotshell-docker](https://github.com/julienmoumne/hotshell-docker)

## Vagrant

Menus with useful Vagrant commands.

  - [demo](http://julienmoumne.github.io/hotshell/demos/vagrant.hs.js.html)
  - [generated markdown](./vagrant/vagrant.hs.js.md)
  - [source](./vagrant/vagrant.hs.js)
  - try it :
```bash
cd $VAGRANT_PROJECT
H='raw.githubusercontent.com/julienmoumne/hotshell/master/examples/vagrant/vagrant.hs.js'; T=`mktemp`; wget $H -qO $T && hs -f $T; rm $T
```

## Topten

Dynamically creates a menu with the 10 most used commands.

  - [demo](http://julienmoumne.github.io/hotshell/demos/topten.hs.js.html)
  - [generated markdown](./topten/topten.hs.js.md)
  - [source](./topten/topten.hs.js)
  - try it :
```bash
H='raw.githubusercontent.com/julienmoumne/hotshell/master/examples/topten/topten.hs.js'; T=`mktemp`; wget $H -qO $T && hs -f $T; rm $T
```

## Network

A menu to manage 'eth0' and 'wlan0'. Showcases the ability to factor out groups of commands.

  - [demo](http://julienmoumne.github.io/hotshell/demos/network.hs.js.html)
  - [generated markdown](./network/network.hs.js.md)
  - [source](./network/network.hs.js)
  - try it :
```bash
H='raw.githubusercontent.com/julienmoumne/hotshell/master/examples/network/network.hs.js'; T=`mktemp`; wget $H -qO $T && hs -f $T; rm $T
```

## Modules

Showcases the ability to modularize menu definitions.

  - [source](./modules)
  - [generated markdown](./modules/modules.hs.js.md)

## Nested Hotshells

Showcases the ability to call Hotshell within Hotshell.

  - [source](./nested/nested.hs.js)
  - [generated markdown](./nested/nested.hs.js.md)
  - try it :
```bash
H='raw.githubusercontent.com/julienmoumne/hotshell/master/examples/nested/nested.hs.js'; T=`mktemp`; wget $H -qO $T && hs -f $T; rm $T
```