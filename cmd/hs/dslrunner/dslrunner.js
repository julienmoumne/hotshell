var _ = require('underscore')
var items = []
module.exports.current = {items: items, wd: './'}
module.exports.item = item
module.exports.items = items

function item(config, callback) {
    var current = module.exports.current

    function executeCallback() {
        if (!_.isFunction(callback))
            return

        module.exports.current = config
        try {
            callback()
        } catch (exception) {
            delete config.items
            config.desc += ' [Exception caught, ' + exception + ']'.trim()
        }
        module.exports.current = current
    }

    function adjustWd() {
        config.wd = current.wd + (_.isUndefined(config.wd) ? '' : config.wd + '/')
    }

    function initChildArray() {
        if (_.isUndefined(current.items))
            current.items = []
    }

    function addItemToParent() {
        current.items.push(config)
    }

    adjustWd()
    initChildArray()
    addItemToParent()
    executeCallback()
}