item({desc: 'hotshell website'}, function () {

    linux = exec('uname').indexOf('Linux') > -1
    browser = linux ? 'sensible-browser' : 'open'
    hotshellDir = '$GOPATH/src/github.com/julienmoumne/hotshell'

    item({key: 's', cmd: 'bundle exec jekyll serve --trace'})
    item({key: 'u', desc: 'update GitHub Pages gem', cmd: 'bundle update'})

    item({key: 'd', desc: 'demos'}, function () {

        item({key: 'g', desc: 'generate\n  ', cmd: script(
                'set -e',
                generateDemo('_includes/demo.hs.js', 'demo.hs.js'),
                generateDemo('demos/tutorial.hs.js', 'tutorial.hs.js'),
                generateDemo(fromHotshell('hs.js'), 'hs.js'),
                generateDemo(fromHotshell('examples/default/default.hs.js'), 'default.hs.js'),
                generateDemo(fromHotshell('examples/docker/docker.hs.js'), 'docker.hs.js'),
                generateDemo(fromHotshell('examples/docker/docker-compose.hs.js'), 'docker-compose.hs.js'),
                'DOCKER_MACHINE_NAME=dev ' + generateDemo(fromHotshell('examples/docker/docker-machine.hs.js'), 'docker-machine.hs.js'),
                generateDemo(fromHotshell('examples/vagrant/vagrant.hs.js'), 'vagrant.hs.js'),
                generateDemo(fromHotshell('examples/topten/topten.hs.js'), 'topten.hs.js'),
                generateDemo(fromHotshell('examples/network/network.hs.js'), 'network.hs.js')
        )})

        item({key: 'o', desc: 'open demos', cmd: 'find demos  -name "*.html" -exec ' + browser + ' {} \\;'})
    })

    item({key: 'h', desc: 'hotshell', cmd: 'hs --chdir -f ' + hotshellDir + '/hs.js'})
})

function generateDemo(src, out) {
    return 'hs --chdir --generate-demo -f ' + src + ' > demos/' + out + '.html'
}

function fromHotshell(file) {
    return hotshellDir + '/' + file
}

function script() {
    src = '';
    _.each(arguments, function (el) {
        src += '   ' + el + '\n'
    })
    return src
}