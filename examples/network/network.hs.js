item({desc: 'network'}, function () {

    item({key: 'e', desc: 'eth0'}, function () {
        interface('eth0')
    })

    item({key: 'w', desc: 'wlan0'}, function () {
        interface('wlan0')
        // specific commands can be added when needed
        item({key: 'l', desc: 'list access points', cmd: 'iwlist scan'})
    })
})

function interface(id) {
    ifconfig = 'ifconfig ' + id
    sudo = 'sudo ' + ifconfig
    item({key: 'i', desc: 'info', cmd: ifconfig})
    item({key: 'u', desc: 'up', cmd: sudo + ' up'})
    item({key: 'd', desc: 'down', cmd: sudo + ' down'})
}