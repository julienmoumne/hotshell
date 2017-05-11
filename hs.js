item({desc: 'hotshell-dev'}, function () {

    linux = exec('uname').indexOf('Linux') > -1
    browser = linux ? 'sensible-browser' : 'open'
    generate = 'go generate ./...'
    install = generate + ' && go install ./...'
    hsCmdDir = './cmd/hs'
    buildAndRun = generate + ' && go build ' + hsCmdDir + ' && ./hs'
    allButVendor = '$(go list ./... | grep -v /vendor/)'
    runTests = generate + ' && go test ' + allButVendor + ' -timeout 10s'

    item({key: 'i', desc: 'install', cmd: install})
    item({key: 'c', desc: 'clean', cmd: 'go clean -i ./... && find . -type f -name \'bindata.go\' -exec rm {} +'})
    item({key: 'v', desc: 'vet', cmd: 'go vet ' + allButVendor})
    item({key: 'f', desc: 'fmt', cmd: 'go fmt ' + allButVendor})

    item({key: 'p', desc: 'packaging'}, function () {

        man = 'debian/usr/share/man/man1/hs.1.gz'
        item({key: 'g', desc: 'generate man', cmd: install + ' && hs-man | gzip > ' + man})
        item({key: 's', desc: 'show man', cmd: 'gunzip -c ' + man + ' | groff -Tascii -man -'})
        item({key: 'p', desc: 'package', cmd: "goxc -pv $(cat VERSION) -wd " + hsCmdDir})
        item({key: 'c', desc: 'generate changelog', cmd: 'github_changelog_generator'})
    })

    item({key: 't', desc: 'tests'}, function () {

        item({key: 't', desc: 'test', cmd: runTests})
        item({key: 'i', desc: 'interactively run end to end tests'}, function () {

            testDir = hsCmdDir + '/test/cases/'
            _(exec('ls -d ' + testDir + '*/').split('\n')).each(function(subdir, ix) {
                item({key: ix, desc: subdir}, function () {

                    _(exec('ls ' + subdir).split('\n')).each(function(testName, ix) {
                        item({key: ix, desc: testName, cmd: buildAndRun + ' --chdir -f ' + subdir + testName})
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

        var examples = exec('ls examples/**/*.js').split('\n')
        _(examples).each(function (el, ix) {
            item({key: ix, desc: el, cmd: buildAndRun + ' --chdir -f ' + el})
        })
        var mdGeneration = ''
        _(examples).each(function (el, ix) {
            mdGeneration += 'hs --generate-doc --chdir -f ' + el + ' > ' + el + '.md;'
        })
        item({key: 'g', desc: 'generate markdowns', cmd: mdGeneration})
    })

    item({key: 'd', desc: 'install dev dependencies', cmd: './scripts/install-dev-deps.sh'})
})