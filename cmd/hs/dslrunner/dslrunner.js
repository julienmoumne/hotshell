var _ = require('underscore')

var items = []
var current = {items: items}

function item(config, callback) {

    function recurse() {
        var prev = current
        current = config

        try {
            callback()
        } catch (exception) {
            delete config.items
            config.desc += ' [Exception caught, ' + exception + ']'.trim()
        }

        current = prev
    }

    if (_.isUndefined(current.items))
        current.items = []

    current.items.push(config)

    if (_.isFunction(callback))
        recurse()
}

module.exports.item = item
module.exports.items = items