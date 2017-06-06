var item = require('hotshell').item
var exec = require('hotshell').exec
var _ = require('underscore')

var services = exec('docker-compose config --services | sort').split('\n')
var commands = [
    {key: 'u', desc: 'up -d', ps: true},
    {key: 'r', desc: 'restart', ps: true},
    {key: 's', desc: 'stop', ps: true},
    {key: 'k', desc: 'kill', ps: true},
    {key: 'l', desc: 'logs'},
    {key: 'p', desc: 'ps'},
    {key: 'b', desc: 'build'},
    {key: 'd', desc: 'rm', ps: true}
]

item({desc: 'docker-compose'}, function () {

    item({key: 'c', desc: 'commands > services'}, function() {
        _(commands).each(createCommandForEveryServices)
    })
    
    item({key: 's', desc: 'services > commands'}, function() {
        createCommandsForSingleService('all', 'a')
        _(services).each(createCommandsForSingleService)
    })
    
    item({key: 'd', desc: 'display config', cmd: 'docker-compose config'})
    
    function createCmdForService(command, service) {
        var service = service == 'all' ? '' : ' ' + service
        var post = _(command.ps).isUndefined() ? '' : ' && docker-compose ps' + service
        return 'docker-compose ' + command.desc + service + post
    }
    
    function createCommandsForSingleService(service, key) {
        item({key: key, desc: service}, function () {
            _(commands).each(function (command, cmdIndex) {
                item({key: command.key, desc: command.desc, cmd: createCmdForService(command, service)})
            })
        })
    }

    function createCommandForEveryServices(command) {
        item(command, function () {
            item({key: 'a', desc: 'all', cmd: createCmdForService(command, 'all')})
            _(services).each(function (service, serviceIndex) {
                item({key: serviceIndex, desc: service, cmd: createCmdForService(command, service)})
            })
        })
    }
})