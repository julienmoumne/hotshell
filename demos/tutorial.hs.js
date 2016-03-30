item({desc: 'top level menu'}, function() {

  item({desc: 'this is a description item, you can use it to describe the menu\n'})

  item({key: 'm', desc: 'apache management'}, function() {
    item({key: 's', cmd: 'sudo service apache2 status'})
    item({key: 'r', cmd: 'sudo service apache2 restart'})
    item({key: 'h', cmd: 'sudo service apache2 stop'})
  })

  item({key: 'l', desc: 'apache logs'}, function() {
    item({key: 'a', desc: 'access', cmd: 'less +F /var/log/apache2/access.log'})
    item({key: 'e', desc: 'error', cmd: 'less +F /var/log/apache2/error.log'})
  })
})