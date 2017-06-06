var item = require('hotshell').item
var submenu = require('./module/submenu.hs.js')

item({desc: 'Submenu Module'}, function () {
    item({key: 's', desc: 'submenu'}, submenu)
})