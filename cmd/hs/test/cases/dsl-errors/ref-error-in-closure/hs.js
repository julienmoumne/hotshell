var item = require('hotshell').item

item({desc: 'Reference Error in closure'}, function () {

    callToUndefinedMethod()
})