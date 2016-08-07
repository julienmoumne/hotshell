var submenu1 = require('./lib/submenu1.hs.js')
var submenu2 = require('./lib/submenu2.hs.js')

item({desc: 'modules'}, function () {
    item({key: 'f', desc: 'First submenu'}, submenu1)
    item({key: 's', desc: 'Second submenu'}, submenu2)
})