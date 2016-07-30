item({desc: 'vagrant'}, function () {

    var vagrantVms = exec('vagrant status | sed -n "s;\\([^ ]*\\).*(.*;\\1;p"').split('\n')

    forAllVMs({key: 'u', desc: 'up', cmd: 'vagrant up'})
    forAllVMs({key: 'r', desc: 'reload', cmd: 'vagrant reload'})
    forAllVMs({key: 'd', desc: 'destroy', cmd: 'vagrant destroy'})
    forAllVMs({key: 'p', desc: 'provision', cmd: 'vagrant provision'})
    forAllVMs({key: 'h', desc: 'halt', cmd: 'vagrant halt'})
    forAllVMs({key: 's', desc: 'ssh', cmd: 'vagrant ssh', individual: true})

    item({key: 'b', desc: 'box update', cmd: 'vagrant box update'})

    function forAllVMs(config) {
        item(_(config).omit('cmd'), function () {

            if (_(config.individual).isUndefined())
                item({key: 'a', desc: 'all', cmd: config.cmd})

            _(vagrantVms).each(function (el, ix) {
                item({key: ix, desc: el, cmd: config.cmd + ' ' + el})
            })
        })
    }
})