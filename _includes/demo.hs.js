var item = require('hotshell').item

item({desc: 'hotshell demo'}, function () {

    // command items
    item({key: 'r', desc: 'restart apache', cmd: 'sudo service apache2 restart'})
    item({key: 's', desc: 'synchronize time', cmd: 'sudo /usr/sbin/ntpdate pool.ntp.org'})

    // submenu to group log related commands
    item({key: 'l', desc: 'apache logs'}, function () {
        item({key: 'a', desc: 'access.log', cmd: 'less +F /var/log/apache2/access.log'})
        item({key: 'e', desc: 'error.log', cmd: 'less +F /var/log/apache2/error.log'})
    })
})  
