item({desc: 'docker'}, function () {  
    
    item({key: 'p', desc: 'ps'}, function () {
        item({key: 'a', desc: 'ps all', cmd: 'docker ps -a'})
        item({key: 'r', desc: 'ps running', cmd: 'docker ps '})
    })
    
    item({key: 'c', desc: 'containers'}, function () {
        
        allContainers = adjustList(exec('docker ps -a --format={{.ID}}'))
        runningContainers = adjustList(exec('docker ps --format={{.ID}}'))
        containerLabels = _.object(_.map(allContainers, function(el) {
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
                        _args = _.isUndefined(delegate.args) ? '' : ' ' + args
                        post = _.isUndefined(delegate.inspect) ? '' : ' && docker inspect --format="{{.Name}} {{.State.Status}}"' + el
                        return 'docker ' + desc + el + _args + post
                    }
                    
                    if (_.isUndefined(delegate.individual)) {
                        item({key: 'a', desc: 'all', cmd: createCmd('$(docker ps -a -q)')})
                        item({key: 'r', desc: 'all running', cmd: createCmd('$(docker ps -q)')})
                    }

                    containers = _.isUndefined(delegate.running) ? allContainers : runningContainers
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
