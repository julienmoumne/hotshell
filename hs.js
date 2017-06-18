var item = require('hotshell').item
var exec = require('hotshell').exec
var _ = require('underscore')

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
        item({key: 'o', desc: 'failed end2end tests', cmd: browser + ' ' + hsCmdDir + '/test/tmp &'})
    })

    item({key: 'e', desc: 'examples'}, function () {

        _(exec('ls examples/**/*.js').split('\n')).each(function (el, ix) {
            hsFile = basename(el)
            item({key: ix, desc: hsFile, wd: dirname(el), cmd: '$GOPATH/bin/hs -f ' + hsFile})
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
