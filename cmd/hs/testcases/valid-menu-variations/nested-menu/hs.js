item({desc: 'Nested Menu'}, function () {

    // item defined using eval
	eval("item({key: 'l', desc: 'ls *.js', cmd: 'ls *.js'})")

    // menu defined solely with a closure
	item(function () {
        key = 's'
        desc = 'Submenu'
        item({desc: 'submenu desc\n'})
    	item({key: 'l', desc: 'ls *.std*', cmd: 'ls *.std*'})
    })
})