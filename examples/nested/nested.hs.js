var item = require('hotshell').item

item({desc: 'nested', wd: '../docker'}, function () {

    item({key: 'd', desc: 'docker', cmd: 'hs -f docker.hs.js'})
    item({key: 'c', desc: 'docker-compose', cmd: 'hs -f docker-compose.hs.js'})
    item({key: 'm', desc: 'docker-machine', cmd: 'hs -f docker-machine.hs.js'})
})