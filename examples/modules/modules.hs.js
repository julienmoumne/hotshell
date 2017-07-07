var item = require('hotshell').item
var submenu = require('./lib/submenu.hs.js')
var randNumberGenerator = require('./lib/commons.hs.js').randNumberGenerator

item({desc: 'modules'}, function () {
    item({key: 'f', desc: 'submenu'}, submenu)
    item({key: 's', desc: 'echo number', cmd: 'echo ' + randNumberGenerator()})
})