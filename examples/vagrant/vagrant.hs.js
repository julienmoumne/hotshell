item({desc: 'vagrant'}, function() {

    vagrantVms = exec('vagrant status | sed -n "s;\\([^ ]*\\).*(.*;\\1;p"').split('\n')

    vmActions('u', 'up', 'vagrant up')
    vmActions('r', 'reload', 'vagrant reload')
    vmActions('d', 'destroy', 'vagrant destroy')
    vmActions('p', 'provision', 'vagrant provision')
    vmActions('h', 'halt', 'vagrant halt')
    vmActions('s', 'ssh', 'vagrant ssh')

    function vmActions(key, desc, action) {
        item({key: key, desc: desc, action: action}, function() {
            item({key: 'a', desc: 'all', cmd: action})
            _.each(vagrantVms, function(el, ix){
                item({key: ix, desc: el, cmd: action + ' ' + el})
            })
        })
    }
})