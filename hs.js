item({desc: 'hotshell-web'}, function() {

    item({key: 's', desc: 'jekyll serve', cmd: 'jekyll serve'})

    item({key: 'g', desc: 'generate demos'}, function() {

        item({key: 'f', desc: 'demo', cmd: 'hs --generate-demo -f _includes/demo.hs.js > demos/demo.hs.js.html'})
        item({key: 'u', desc: 'tutorial', cmd: 'hs --generate-demo -f demos/tutorial.hs.js > demos/tutorial.hs.js.html'})
        item({key: 'h', desc: 'hotshell', cmd: demo('hs.js', 'hs.js')})
        item({key: 'd', desc: 'default', cmd: demo('examples/default/default.hs.js', 'default.hs.js')})
        item({key: 'r', desc: 'docker', cmd: demo('examples/docker/docker.hs.js', 'docker.hs.js')})
        item({key: 'c', desc: 'docker-compose', cmd: demo('examples/docker/docker-compose.hs.js', 'docker-compose.hs.js')})
        item({key: 'm', desc: 'docker-machine', cmd: demo('examples/docker/docker-machine.hs.js', 'docker-machine.hs.js')})
        item({key: 'v', desc: 'vagrant', cmd: demo('examples/vagrant/vagrant.hs.js', 'vagrant.hs.js')})
        item({key: 't', desc: 'topten', cmd: demo('examples/topten/topten.hs.js', 'topten.hs.js')})

        function demo(src, out) {
            return 'hs --chdir --generate-demo ' +
                '-f $GOPATH/src/github.com/julienmoumne/hotshell/' + src +
                '> demos/'+ out +'.html'
        }
    })

})