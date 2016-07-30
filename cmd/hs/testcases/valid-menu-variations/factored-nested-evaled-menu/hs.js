item({desc: 'Factored Menu'}, function () {

    // item defined using eval
    eval("item({key: 'k', desc: 'ls *.js', cmd: 'ls *.js'})")

    factoredMenu({key: 'l', desc: 'less', cmd: 'less +F', els: ['/etc/hosts', '/etc/fstab']})

    item({key: 's', desc: 'factored submenu'}, function () {
        item({desc: 'submenu desc\n'})
        factoredMenu({key: 'e', desc: 'echo', cmd: 'echo', els: ['hello', 'world']})
    })
})

function factoredMenu(config) {
    item(_(config).omit('cmd'), function () {
        _(config.els).each(function (el, ix) {
            item({key: ix, cmd: config.cmd + ' ' + el})
        })
    })
}