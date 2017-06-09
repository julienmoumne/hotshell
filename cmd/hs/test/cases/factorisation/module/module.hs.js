var item = require('hotshell').item
var _ = require('underscore')

module.exports = function (config) {
     item(_(config).omit('cmd'), function () {
         _(config.els).each(function (el, ix) {
             item({key: ix, cmd: config.cmd + ' ' + el})
         })
     })
 }