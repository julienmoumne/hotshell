var item = require('hotshell').item
var randNumberGenerator = require('./commons.hs.js').randNumberGenerator

module.exports = function () {
    item({key: 'e', desc: 'echo first number', cmd: 'echo ' + randNumberGenerator()})
}