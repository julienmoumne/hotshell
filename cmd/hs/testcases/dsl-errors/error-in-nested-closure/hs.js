item({desc: 'Runtime Error in nested closure'}, function () {

    item({desc: 'nested closure'}, function () {

        item({key: 's', desc: 'skipped'})
        throw new Error('Runtime Error')
    })

    item({key: 'n', desc: 'not skipped', cmd: 'echo not skipped'})
})