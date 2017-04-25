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
  - [source](../hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/hs.js
```

## Hotshell's default menu

This menu is loaded when no definition file is found.

  - [demo](http://julienmoumne.github.io/hotshell/demos/default.hs.js.html)
  - [source](./default/default.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/examples/default/default.hs.js
```

## Docker, Docker Compose & Docker Machine

Menus with useful Docker commands.

  - demos
    * [docker](http://julienmoumne.github.io/hotshell/demos/docker.hs.js.html)
    * [docker-compose](http://julienmoumne.github.io/hotshell/demos/docker-compose.hs.js.html)
    * [docker-machine](http://julienmoumne.github.io/hotshell/demos/docker-machine.hs.js.html)
  - sources
    * [docker](./docker/docker.hs.js)
    * [docker-compose](./docker/docker-compose.hs.js)
    * [docker-machine](./docker/docker-machine.hs.js)
  - try it :
```bash
# docker
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/examples/docker/docker.hs.js

# docker-machine
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/examples/docker/docker-machine.hs.js

# docker-compose
cd $DOCKER_COMPOSE_PROJECT
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/examples/docker/docker-compose.hs.js
```
  - integrate it in your own menus using [hotshell-docker](https://github.com/julienmoumne/hotshell-docker)

## Vagrant

Menus with useful Vagrant commands.

  - [demo](http://julienmoumne.github.io/hotshell/demos/vagrant.hs.js.html)
  - [source](./vagrant/vagrant.hs.js)
  - try it :
```bash
cd $VAGRANT_PROJECT
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/examples/vagrant/vagrant.hs.js
```

## Topten

Dynamically creates a menu with the 10 most used commands.

  - [demo](http://julienmoumne.github.io/hotshell/demos/topten.hs.js.html)
  - [source](./topten/topten.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/examples/topten/topten.hs.js
```

## Network

A menu to manage 'eth0' and 'wlan0'. Showcases the ability to factor out groups of commands.

  - [demo](http://julienmoumne.github.io/hotshell/demos/network.hs.js.html)
  - [source](./network/network.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/examples/network/network.hs.js
```

## Modules

Showcases the ability to modularize menu definitions.

  - [source](./modules)

## Nested Hotshells

Showcases the ability to call Hotshell within Hotshell.

  - [source](./nested/nested.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/examples/nested/nested.hs.js
```