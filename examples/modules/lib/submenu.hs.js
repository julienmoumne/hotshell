var item = require('hotshell').item
var randNumberGenerator = require('./commons.hs.js').randNumberGenerator

module.exports = function () {
    item({key: 'e', desc: 'echo number', cmd: 'echo ' + randNumberGenerator()})
}

// these 2 lines allow the module to be used both as a reusable component and as a standalone menu
// this is the same idea as https://docs.python.org/3/library/__main__.html
if (require.main === module)
    item({desc: 'menu'}, module.exports)