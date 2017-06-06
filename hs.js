var item = require('hotshell').item
var exec = require('hotshell').exec
var _ = require('underscore')

item({desc: 'hotshell website'}, function () {

    linux = exec('uname').indexOf('Linux') > -1
    browser = linux ? 'sensible-browser' : 'open'
    hotshellDir = '$GOPATH/src/github.com/julienmoumne/hotshell'

    item({key: 's', cmd: 'bundle exec jekyll serve --trace'})
    item({key: 'u', desc: 'update GitHub Pages gem', cmd: 'bundle update'})

    item({key: 'd', desc: 'demos'}, function () {

        item({key: 'g', desc: 'generate\n  ', cmd: script(
                'set -e',
                generateDemo('_includes', 'demo.hs.js'),
                generateDemo('demos', 'tutorial.hs.js'),
                generateDemo(fromHotshell(''), 'hs.js'),
                generateDemo(fromHotshell('examples/default'), 'default.hs.js'),
                generateDemo(fromHotshell('examples/docker'), 'docker.hs.js'),
                generateDemo(fromHotshell('examples/docker'), 'docker-compose.hs.js'),
                generateDemo(fromHotshell('examples/docker'), 'docker-machine.hs.js'),
                generateDemo(fromHotshell('examples/vagrant'), 'vagrant.hs.js'),
                generateDemo(fromHotshell('examples/topten'), 'topten.hs.js'),
                generateDemo(fromHotshell('examples/network'), 'network.hs.js')
        )})

        item({key: 'o', desc: 'open demos', cmd: 'find demos  -name "*.html" -exec ' + browser + ' {} \\;'})
    })

    item({key: 'h', desc: 'hotshell', cmd: 'cd ' + hotshellDir + '; hs -f hs.js'})
})

function generateDemo(dir, name) {
    return '(cd ' + dir + '; DOCKER_MACHINE_NAME=dev hs --generate-demo -f ' + name + ') > demos/' + name + '.html'
}

function fromHotshell(dir) {
    return hotshellDir + '/' + dir
}

function script() {
    src = '';
    _.each(arguments, function (el) {
        src += '   ' + el + '\n'
    })
    return src
}