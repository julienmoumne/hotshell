var item = require('hotshell').item
var exec = require('hotshell').exec
var _ = require('underscore')

item({desc: 'docker'}, function () {

    item({key: 'p', desc: 'ps'}, function () {
        item({key: 'a', desc: 'ps all', cmd: 'docker ps -a'})
        item({key: 'r', desc: 'ps running', cmd: 'docker ps '})
    })

    item({key: 'c', desc: 'containers'}, function () {

        var allContainers = adjustList(exec('docker ps -a --format={{.ID}}'))
        var runningContainers = adjustList(exec('docker ps --format={{.ID}}'))
        var containerLabels = _.object(_.map(allContainers, function (el) {
            return [el, exec('docker inspect --format="{{.Name}} ({{.Config.Image}})" ' + el)]
        }));

        forAllContainers({key: 'i', desc: 'inspect'})
        forAllContainers({key: 's', desc: 'start', inspect: true})
        forAllContainers({key: 'e', desc: 'exec -it', running: true, individual: true, args: 'bash'})
        forAllContainers({key: 'h', desc: 'stop', inspect: true, running: true})
        forAllContainers({key: 'k', desc: 'kill', inspect: true, running: true})
        forAllContainers({key: 'r', desc: 'restart', inspect: true})
        forAllContainers({key: 'l', desc: 'logs -f', individual: true, running: true})
        forAllContainers({key: 't', desc: 'top', individual: true})
        forAllContainers({key: 'd', desc: 'rm -f'})

        function adjustList(list) {
            return list == '' ? [] : list.split('\n')
        } 
            
        function forAllContainers (config) {        
            item(config, function () {                          
                    function createCmd(el) {
                        el = ' ' + el
                        var _args = _(config.args).isUndefined() ? '' : ' ' + config.args
                        var post = _(config.inspect).isUndefined() ? '' : ' && docker inspect --format="{{.Name}} {{.State.Status}}"' + el
                        return 'docker ' + config.desc + el + _args + post
                    }
                    
                    if (_(config.individual).isUndefined()) {
                        item({key: 'a', desc: 'all', cmd: createCmd('$(docker ps -a -q)')})
                        item({key: 'r', desc: 'all running', cmd: createCmd('$(docker ps -q)')})
                    }

                    var containers = _(config.running).isUndefined() ? allContainers : runningContainers
                    _(containers).each(function(el, ix){
                        item({
                            key: ix,
                            desc: containerLabels[el],
                            cmd: createCmd(el)
                        })
                    })
                }
            )
        }
    })

    item({key: 'i', desc: 'images'}, function () {
        item({key: 'l', desc: 'list', cmd: 'docker images'})
        item({key: 'd', desc: 'remove all', cmd: 'docker rmi $(docker images -q)'})
    })

    item({key: 'v', desc: 'volumes'}, function () {
        item({key: 'l', desc: 'list', cmd: 'docker volume ls'})
        item({key: 'd', desc: 'remove all', cmd: 'docker volume rm $(docker volume ls -q)'})
    })

    item({key: 'n', desc: 'networks', cmd: 'docker network ls'})
    item({key: 's', desc: 'stats', cmd: 'docker stats'})
})
