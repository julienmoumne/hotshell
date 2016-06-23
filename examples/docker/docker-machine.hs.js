machine = exec('echo $DOCKER_MACHINE_NAME')
if (machine == '') throw 'please set $DOCKER_MACHINE_NAME'

item({desc: 'docker-machine'}, function () {

    item({desc: 'active machine is "' + machine + '"\n'})

    item({key: 'd', desc: 'display env', cmd: 'docker-machine env ' + machine})
    item({key: 'e', desc: 'eval env', cmd: 'eval "$(docker-machine env ' + machine + ')"'})
    item({desc: ' '})
    item({key: 's', desc: 'ssh', cmd: 'docker-machine ssh ' + machine})
    item({desc: ' '})
    item({key: 't', desc: 'status', cmd: 'docker-machine status ' + machine})
    item({key: 'r', desc: 'restart', cmd: 'docker-machine restart ' + machine})
    item({key: 'h', desc: 'stop', cmd: 'docker-machine stop ' + machine})
    item({key: 'k', desc: 'kill', cmd: 'docker-machine kill ' + machine})
    item({desc: ' '})
    item({key: 'l', desc: 'list', cmd: 'docker-machine ls'})
    item({desc: ' '})
    item({key: 'c', desc: 'regenerate-certs', cmd: 'docker-machine regenerate-certs ' + machine})
    item({key: 'u', desc: 'upgrade', cmd: 'docker-machine upgrade ' + machine})
})
