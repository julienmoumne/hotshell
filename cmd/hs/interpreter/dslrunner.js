var _ = require('underscore')

var items = []
var current = {items: items}

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

module.exports.item = item
module.exports.items = items