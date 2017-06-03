item({desc: 'hotshell-dev'}, function () {

    linux = exec('uname').indexOf('Linux') > -1
    browser = linux ? 'sensible-browser' : 'open'
    hsCmdDir = './cmd/hs'
    generate = './scripts/generate.sh'
    runTests = './scripts/test.sh'

    item({key: 'c', desc: 'clean install', cmd: './scripts/clean-install.sh'})
    item({key: 'f', desc: 'format', cmd: './scripts/format.sh'})

    item({key: 'p', desc: 'packaging'}, function () {

        man = 'debian/usr/share/man/man1/hs.1.gz'
        item({key: 'm', desc: 'test man', cmd: './scripts/generate-man.sh; gunzip -c ' + man + ' | groff -Tascii -man -'})
        item({key: 'p', desc: 'package', cmd: "./scripts/package.sh"})
    })

    item({key: 't', desc: 'tests'}, function () {

        item({key: 't', desc: 'test', cmd: runTests})

        item({key: 'f', desc: 'failed end to end tests'}, function () {

            item({desc: '(reload menu to update failed tests list)\n'})

            failedTestsDir = hsCmdDir + '/test/tmp/failed-cases'

            item({key: 'r', desc: 'run tests', cmd: runTests})
            item({key: 'o', desc: 'open failed tests directory', cmd: browser + ' ' + failedTestsDir})

            if (exec('if [ -d "' + failedTestsDir + '" ]; then echo true; fi') == '') return

            _(exec('find ' + failedTestsDir + ' -name *.html').split('\n')).each(function (el, ix) {
                item({key: ix, desc: 'failed test ' + ix, cmd: browser + ' ' + el})
            })
        })
    })

    item({key: 'e', desc: 'examples'}, function () {

        _(exec('ls examples/**/*.js').split('\n')).each(function (el, ix) {
            hsFile = basename(el)
            item({key: ix, desc: hsFile, cmd: 'cd ' + dirname(el) + '; $GOPATH/bin/hs -f ' + hsFile})
        })
    })

    item({key: 'd', desc: 'doc'}, function() {
        item({key: 'c', desc: 'generate changelog', cmd: './scripts/generate-changelog.sh'})
        item({key: 'm', desc: 'generate md', cmd: './scripts/generate-md.sh'})
    })
    item({key: 'i', desc: 'install dev dependencies', cmd: './scripts/install-dev-deps.sh'})
    
    item({key: 'v', desc: 'vendoring'}, function() {
        item({key: 'a', desc: 'add dependency', cmd: './scripts/add-dependency.sh'})
        item({key: 'd', desc: 'delete dependency', cmd: './scripts/rm-dependency.sh'})
        item({key: 'r', desc: 'reset vendor', cmd: 'find ./vendor/* -not -name \'vendor.json\' -print0 | xargs -0 rm -Rf –-- && govendor sync'})
    })

    function basename(path) {
         return path.replace(dirname(path), "")
    }

    function dirname(path) {
         return path.match(".*/")
    }
})
