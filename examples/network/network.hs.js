item({desc: 'network'}, function() {

 item({key: 'e', desc: 'eth0'}, function() {
   interface(delegate, 'eth0')
 })

 item({key: 'w', desc: 'wlan0'}, function() {
   interface(delegate, 'wlan0')
   // specific commands can be added here when needed
   item({key: 'l', desc: 'list access points', cmd: 'iwlist scan'})
 })
})

function interface(delegate, id) {
  with(delegate) { // see http://alexyoung.org/2009/10/22/javascript-dsl
    ifconfig = 'ifconfig ' + id
    sudo = 'sudo ' + ifconfig
    item({key: 'i', desc: 'info', cmd: ifconfig})
    item({key: 'u', desc: 'up', cmd: sudo + ' up'})
    item({key: 'd', desc: 'down', cmd: sudo + ' down'})
  }
}