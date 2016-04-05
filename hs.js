item({desc: 'hotshell-dev'}, function() {

    linux = exec('uname').indexOf('Linux') > -1
    browser = linux ? 'sensible-browser' : 'open'
    generate = 'go generate ./...'
    install = generate + ' && go install ./...'
    hsCmdDir = './cmd/hs'
    buildAndRun = generate + ' && go build ' + hsCmdDir + ' && ./hs'
    allButVendor =  '$(go list ./... | grep -v /vendor/)'
    runTests = generate + ' && go test ' + allButVendor + ' -timeout 10s'

    item({key: 'i', desc: 'install', cmd: install})
    item({key: 'c', desc: 'clean', cmd: 'go clean -i ./...'})
    item({key: 'v', desc: 'vet', cmd: 'go vet ' + allButVendor})
    item({key: 'f', desc: 'fmt', cmd: 'go fmt ' + allButVendor})

    item({key: 'p', desc: 'packaging'}, function() {

        man = 'debian/usr/share/man/man1/hs.1.gz'
        item({key: 'g', desc: 'generate man', cmd: install + ' && hs-man | gzip > ' + man})
        item({key: 's', desc: 'show man', cmd: 'gunzip -c ' + man + ' | groff -Tascii -man -'})
        item({key: 'p', desc: 'package', cmd: "goxc -pv $(cat VERSION) -wd " + hsCmdDir})
    })

    item({key: 't', desc: 'tests'}, function() {

        item({key: 't', desc: 'test', cmd: runTests})
        item({key: 'i', desc: 'interactively run end to end tests'}, function() {

            testDir = hsCmdDir + '/testcases/'
            _.each(exec('ls -d ' + testDir + '*/').split('\n'), function(subdir, ix) {
                item({key: ix, desc: subdir}, function() {

                    _.each(exec('ls ' + desc).split('\n'), function(testName, ix) {
                        item({key: ix, desc: testName, cmd: buildAndRun + ' -f ' + desc + testName})
                    })
                })
            })
        })

        item({key: 'f', desc: 'failed end to end tests'}, function() {                    

            item({desc: '(reload menu to update failed tests list)\n'})

            faileTestsDir = hsCmdDir + '/tmp/failed-cases'

            item({key: 'r', desc: 'run tests', cmd: runTests})
            item({key: 'o', desc: 'open failed tests directory', cmd: browser + ' ' + faileTestsDir})

            if (exec('if [ -d "'+faileTestsDir+'" ]; then echo true; fi') == '') return
             
            _.each(exec('find ' + faileTestsDir + ' -name *.html').split('\n'), function(el, ix) {
                item({key: ix, desc: 'failed test ' + ix, cmd: browser + ' ' + el})
            })
        })
    })

    item({key: 'e', desc: 'examples'}, function() {

        _.each(exec('ls examples/**/*.js').split('\n'), function(el, ix) {
            item({key: ix, desc: el, cmd: buildAndRun + ' --chdir -f ' + el})
        })
    })
    
    item({key: 'd', desc: 'dev env setup'}, function() {
        item({key: 'b', cmd: 'go get -u github.com/jteeuwen/go-bindata/...'})
        item({key: 'v', cmd: 'go get -u github.com/kardianos/govendor'})
        item({key: 'x', cmd: 'go get -u github.com/laher/goxc'})
    })

    item({key: 'b', desc: 'brew'}, function() {
        item({key: 't', desc: 'test' + '\n  ', cmd: script(
            'set -eu',
            'echo -n "[git hash] "',
            'read githash',
            'formula=https://raw.githubusercontent.com/julienmoumne/homebrew/$githash/Library/Formula/hs.rb',
            'brew remove hs || true',
            'brew cleanup -s',
            'brew install $formula',
            'brew test $formula',
            'hs'
        )})
    })
})

function script () {
    src = '';
    _.each(arguments, function (el, ix) { src += '   ' + el + '\n' })
    return src
}