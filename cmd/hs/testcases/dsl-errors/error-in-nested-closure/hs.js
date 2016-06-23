item({desc: 'Runtime Error in nested closure'}, function () {

    item({desc: 'nested closure'}, function () {

        throw new Error('Runtime Error')
    })
})