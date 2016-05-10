# Examples of menus

  - [Hotshell's hotshell](#hotshells-hotshell)
  - [Hotshell's default menu](#hotshells-default-menu)
  - [Docker, Docker Compose & Docker Machine](#docker-docker-compose--docker-machine)
  - [Vagrant](#vagrant)
  - [Topten](#topten)
  - [Network](#network)
  - [Includes](#includes)

## Hotshell's hotshell

Used to develop, build and test Hotshell.

  - [demo](http://julienmoumne.github.io/hotshell/demos/hs.js.html)
  - [source](../hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/hs.js
```

## Hotshell's default menu

This menu is loaded when no definition file is found.

  - [demo](http://julienmoumne.github.io/hotshell/demos/default.hs.js.html)
  - [source](./default/default.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/examples/default/default.hs.js
```

## Docker, Docker Compose & Docker Machine

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
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/examples/docker/docker.hs.js

# docker-machine
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/examples/docker/docker-machine.hs.js

# docker-compose
cd $DOCKER_COMPOSE_PROJECT
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/examples/docker/docker-compose.hs.js
```

## Vagrant

  - [demo](http://julienmoumne.github.io/hotshell/demos/vagrant.hs.js.html)
  - [source](./vagrant/vagrant.hs.js)
  - try it :
```bash
cd $VAGRANT_PROJECT
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/examples/vagrant/vagrant.hs.js
```

## Topten

Dynamically creates a menu with the 10 most used commands.

  - [demo](http://julienmoumne.github.io/hotshell/demos/topten.hs.js.html)
  - [source](./topten/topten.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/v0.1.0/examples/topten/topten.hs.js
```

## Network

A menu to manage 'eth0' and 'wlan0'. Showcases the ability to factor out groups of commands.

  - [demo](http://julienmoumne.github.io/hotshell/demos/network.hs.js.html)
  - [source](./network/network.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/examples/network/network.hs.js
```

## Includes

Showcases the ability to include menus defined in separate files.

  - [source](./includes/includes.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hotshell/master/examples/includes/includes.hs.js
```