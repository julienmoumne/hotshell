this.items = []

function item () {

    function initItem (params) {

		// properties need to be set so they can be found using 'with'
    	var itemProperties = {
			key: null
			desc: null
			cmd: null
			item: item
			items: []
		}

    	return _.extend(itemProperties, params)
    }

    function execWithDelegate (func, delegate) {

        try {

            // https://regex101.com/r/bN4sS3/2
            // name the closure so stacktraces are more readable
            func = func.toString().replace(/^function\s*\(/, "function configClosure(")

            eval('with(delegate){(' + func + ')()}');

        } catch (exception) {

            delegate.items = []
            nameWithError = delegate.desc + ' [Exception caught, ' + exception + ']'
            delegate.desc = nameWithError.trim()
        }
    }

    function extractOptions (args) {

        var firstArg = args[0]
        var secondArg = args[1]
        var firstArgIsFunction = _.isFunction(firstArg)

        return {
            params: firstArgIsFunction ? {} : firstArg
            configClosure: firstArgIsFunction ? firstArg : _.isFunction(secondArg) ? secondArg : function() {}
        }
    }

    var options = extractOptions(arguments)

    var newItem = initItem(options.params)

    execWithDelegate(options.configClosure, newItem)

    this.items.push(newItem)
}