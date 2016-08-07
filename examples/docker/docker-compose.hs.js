item({desc: 'docker-compose'}, function () {

    var services = exec('docker-compose config --services | sort').split('\n')

    forAllServices({key: 'u', desc: 'up -d', ps: true})
    forAllServices({key: 'r', desc: 'restart', ps: true})
    forAllServices({key: 's', desc: 'stop', ps: true})
    forAllServices({key: 'k', desc: 'kill', ps: true})
    forAllServices({key: 'l', desc: 'logs'})
    forAllServices({key: 'p', desc: 'ps'})
    forAllServices({key: 'b', desc: 'build'})
    forAllServices({key: 'd', desc: 'rm', ps: true})

    item({key: 'c', desc: 'validate & display config', cmd: 'docker-compose config'})

    function forAllServices(config) {
        item(config, function () {
                function createCmd(el) {
                    el = _(el).isUndefined() ? '' : ' ' + el
                    var post = _(config.ps).isUndefined() ? '' : ' && docker-compose ps' + el
                    return 'docker-compose ' + config.desc + el + post
                }

                item({key: 'a', desc: 'all', cmd: createCmd()})
                _(services).each(function (el, ix) {
                    item({key: ix, desc: el, cmd: createCmd(el)})
                })
            }
        )
    }
})
