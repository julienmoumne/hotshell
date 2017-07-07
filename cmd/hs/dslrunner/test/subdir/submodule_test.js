var item = require('hotshell').item
var exec = require('hotshell').exec

module.exports = function () {
    item({key: 'e', cmd: exec('pwd')})
}

if (require.main === module)
    item({desc: 'menu'}, module.exports)