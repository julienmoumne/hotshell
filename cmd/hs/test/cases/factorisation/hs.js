var item = require('hotshell').item
var factoredMenu = require('./module/module.hs.js')

item({desc: 'Factored Menu'}, function () {
    factoredMenu({key: 'l', desc: 'less', cmd: 'less +F', els: ['/etc/hosts', '/etc/fstab']})
    item({key: 's', desc: 'factored submenu'}, function () {
        item({desc: 'submenu desc\n'})
        factoredMenu({key: 'e', desc: 'echo', cmd: 'echo', els: ['hello', 'world']})
    })
})

