item({desc: 'hotshell-dev'}, function () {

    linux = exec('uname').indexOf('Linux') > -1
    browser = linux ? 'sensible-browser' : 'open'
    hsCmdDir = './cmd/hs'
    generate = './scripts/generate.sh'
    installAndRun = './scripts/install.sh; $GOPATH/bin/hs'
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
        item({key: 'i', desc: 'interactively run end to end tests'}, function () {

            testDir = hsCmdDir + '/test/cases/'
            _(exec('ls -d ' + testDir + '*/').split('\n')).each(function(subdir, ix) {
                item({key: ix, desc: subdir}, function () {

                    _(exec('ls ' + subdir).split('\n')).each(function(testName, ix) {
                        item({key: ix, desc: testName, cmd: installAndRun + ' --chdir -f ' + subdir + testName})
                    })
                })
            })
        })

        item({key: 'f', desc: 'failed end to end tests'}, function () {

            item({desc: '(reload menu to update failed tests list)\n'})

            faileTestsDir = hsCmdDir + '/tmp/failed-cases'

            item({key: 'r', desc: 'run tests', cmd: runTests})
            item({key: 'o', desc: 'open failed tests directory', cmd: browser + ' ' + faileTestsDir})

            if (exec('if [ -d "' + faileTestsDir + '" ]; then echo true; fi') == '') return

            _(exec('find ' + faileTestsDir + ' -name *.html').split('\n')).each(function (el, ix) {
                item({key: ix, desc: 'failed test ' + ix, cmd: browser + ' ' + el})
            })
        })
    })

    item({key: 'e', desc: 'examples'}, function () {

        _(exec('ls examples/**/*.js').split('\n')).each(function (el, ix) {
            exampleName = el.substr(el.lastIndexOf('/') + 1)
            item({key: ix, desc: exampleName, cmd: installAndRun + ' --chdir -f ' + el})
        })
    })

    item({key: 'g', desc: 'generate doc', cmd: './scripts/generate-doc.sh'})
    item({key: 'i', desc: 'install dev dependencies', cmd: './scripts/install-dev-deps.sh'})
    
    item({key: 'v', desc: 'vendoring'}, function() {
        item({key: 'a', desc: 'add dependency', cmd: './scripts/add-dependency.sh'})
        item({key: 'd', desc: 'delete dependency', cmd: './scripts/rm-dependency.sh'})
        item({key: 'r', desc: 'reset vendor', cmd: 'find ./vendor/* -not -name \'vendor.json\' -print0 | xargs -0 rm -Rf –-- && govendor sync'})
    })
})
