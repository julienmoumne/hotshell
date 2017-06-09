var item = require('hotshell').item

item({desc: 'bash as item'}, function () {

    item({key: 'b', desc: 'bash', cmd: 'bash -l'})
})