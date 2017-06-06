var item = require('hotshell').item
var exec = require('hotshell').exec

item({desc: 'Failed Exec'}, function () {

    exec('cat /im/am/confident/this/file/does/not/exist')
})