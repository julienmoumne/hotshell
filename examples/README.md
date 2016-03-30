# Examples of menus

  - [Hotshell's hotshell](#hotshells-hotshell)
  - [Hotshell's default menu](#hotshells-default-menu)
  - [Docker, Docker Compose & Docker Machine](#docker-docker-compose--docker-machine)
  - [Vagrant](#vagrant)
  - [Topten](#topten)
  - [Composite](#composite)

## Hotshell's hotshell

Used to develop, build and test Hotshell.

  - [demo](https://julienmoumne.github.com/hs/demos/hs.js.html)
  - [source](../hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hs/0.1.0/hs.js
```

## Hotshell's default menu

This menu is loaded when no definition file is found.

  - [demo](https://julienmoumne.github.com/hs/demos/default.hs.js.html)
  - [source](./default/default.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hs/0.1.0/examples/default/default.hs.js
```

## Docker, Docker Compose & Docker Machine

  - demos
    * [docker](https://julienmoumne.github.com/hs/demos/docker.hs.js.html)
    * [docker-compose](https://julienmoumne.github.com/hs/demos/docker-compose.hs.js.html)
    * [docker-machine](https://julienmoumne.github.com/hs/demos/docker-machine.hs.js.html)
  - sources
    * [docker](./docker/docker.hs.js)
    * [docker-compose](./docker/docker-compose.hs.js)
    * [docker-machine](./docker/docker-machine.hs.js)
  - try it :
```bash
# docker
hs -f https://raw.githubusercontent.com/julienmoumne/hs/0.1.0/examples/docker/docker.hs.js

# docker-machine
hs -f https://raw.githubusercontent.com/julienmoumne/hs/0.1.0/examples/docker/docker-machine.hs.js

# docker-compose
cd $DOCKER_COMPOSE_PROJECT
hs -f https://raw.githubusercontent.com/julienmoumne/hs/0.1.0/examples/docker/docker-compose.hs.js
```

## Vagrant

  - [demo](https://julienmoumne.github.com/hs/demos/vagrant.hs.js.html)
  - [source](./vagrant/vagrant.hs.js)
  - try it :
```bash
cd $VAGRANT_PROJECT
hs -f https://raw.githubusercontent.com/julienmoumne/hs/0.1.0/examples/vagrant/vagrant.hs.js
```

## Topten

Dynamically creates a menu with the 10 most used commands.

  - [demo](https://julienmoumne.github.com/hs/demos/topten.hs.js.html)
  - [source](./topten/topten.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hs/0.1.0/examples/topten/topten.hs.js
```

## Composite

Showcases the ability to include menus defined in separate files.

  - [source](./composite/composite.hs.js)
  - try it :
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hs/0.1.0/examples/composite/composite.hs.js
```