var items = []
var options = {}
var current = {items: items}

function configure( opts ) {
    if (opts && typeof opts == 'object') {
        options = opts
    }
}

function item(config, callback) {

    function initSubItems() {
        config.items = []
    }

    function recurse() {
        var prev = current
        current = config

        try {
            callback()
        } catch (exception) {
            initSubItems()
            config.desc += ' [Exception caught, ' + exception + ']'.trim()
        }

        current = prev
    }

    initSubItems()

    current.items.push(config)

    if (_.isFunction(callback))
        recurse()
}