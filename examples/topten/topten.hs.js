item({desc: 'topten'}, function() {

    topTen =
        'echo "history" | bash -i 2>/dev/null | sed "s/^ *[0-9]* *//" | ' +
        'sort | uniq -c | sort -nr | head | sed "s/^ *[0-9]* *//"'

    _.each(exec(topTen).split('\n'), function(el, ix) {
        item({key: ix, cmd: el})
    })
})