var item = require('hotshell').item
var exec = require('hotshell').exec
var _ = require('underscore')

item({desc: 'topten'}, function () {

    var topTen =
        'echo "history" | bash -i 2>/dev/null | sed "s/^ *[0-9]* *//" | ' +
        'sort | uniq -c | sort -nr | head | sed "s/^ *[0-9]* *//"'

    _(exec(topTen).split('\n')).each(function (el, ix) {
        item({key: ix, cmd: el})
    })
})