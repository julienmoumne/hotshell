item({desc: 'nested'}, function () {

    item({key: 'd', desc: 'docker', cmd: 'hs -f ../docker/docker.hs.js'})
    item({key: 'c', desc: 'docker-compose', cmd: 'hs -f ../docker/docker-compose.hs.js'})
    item({key: 'm', desc: 'docker-machine', cmd: 'hs -f ../docker/docker-machine.hs.js'})
})