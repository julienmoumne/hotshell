item({desc: 'Hotshell'}, function () {

    linux = exec('uname').indexOf('Linux') > -1
    browser = linux ? 'sensible-browser' : 'open'

    item({desc: "This is Hotshell's default menu, displayed when no definition file is provided.\n"});

    item({desc: "  Don't know what Hotshell is about ?\n"});
    item({key: 'a', desc: "Hit 'a' to get to know it"}, function () {

        item({desc: "Hotshell is a command-line application to efficiently recall and share commands.\n"});

        innerDemo =
            "   // command items\n" +
            "   item({key: 'r', desc: 'restart apache', cmd: 'sudo service apache2 restart'})\n" +
            "   item({key: 's', desc: 'synchronize time', cmd: 'sudo /usr/sbin/ntpdate pool.ntp.org'})\n" +
            "   \n" +
            "   // submenu to group log related commands\n" +
            "   item({key: 'l', desc: 'apache logs'}, function () {\n" +
            "     item({key: 'a', desc: 'access.log', cmd: 'less +F /var/log/apache2/access.log'})\n" +
            "     item({key: 'e', desc: 'error.log', cmd: 'less +F /var/log/apache2/error.log'})\n" +
            "   })\n"

        demo = "item({desc: 'demo'}, function () {\n" + innerDemo + " })\n"

        item({desc: "Menus are defined using a JavaScript configuration DSL. Example :\n"});
        item({desc: demo})
        item({desc: "This creates the following menu :\n"});
        eval(innerDemo)

        item({desc: "\n Want to learn more ?\n"})
        item({key: 'm', desc: "Hit 'm' to open Hotshell's page", cmd: browser + ' https://github.com/julienmoumne/hotshell'})

        item({desc: "\n Hit the spacebar to go to the previous menu or ctrl+c to quit Hotshell"})
    })

    item({desc: "\n   Having trouble with Hothell's command line options?\n"});
    item({key: 'h', desc: "Hit 'h' to display the help", cmd: 'hs --help'})

    item({desc: "\n   Want to see examples of commands?\n"});
    item({key: 'e', desc: "Hit 'e'"}, function () {

        item({key: 'w', desc: 'weather & time'}, function () {
            item({key: 'c', desc: 'check the weather', cmd: 'curl wttr.in'})
            item({key: 's', desc: 'synchronize time', cmd: 'sudo /usr/sbin/ntpdate pool.ntp.org'})
            item({key: 't', desc: 'watch the time go by', cmd: 'watch -n 1 date'})
        })

        item({key: 's', desc: 'SSH your most accessed servers'}, function () {
            item({key: 'l', desc: 'localhost', cmd: 'ssh localhost'})
            item({key: 'g', desc: 'Grex Public Access UNIX', cmd: 'ssh newuser@grex.org'})
        })

        item({key: 'f', desc: 'edit your configuration files'}, function () {
            item({key: 'b', desc: 'vim ~/.bash_profile', cmd: 'vim ~/.bash_profile'})
            item({key: 'B', desc: 'emacs ~/.bash_profile', cmd: 'emacs ~/.bash_profile'})
            item({key: 'h', desc: 'vim /etc/hosts', cmd: 'sudo vim /etc/hosts'})
            item({key: 'H', desc: 'emacs /etc/hosts', cmd: 'sudo emacs /etc/hosts'})
        })

        item({key: 'n', desc: 'network & system utils'}, function () {
            item({key: 'f', desc: 'find text in files', cmd: 'echo -n "[location] [pattern] "; read l p; grep -rnws $l -e $p'})
            item({key: 'o', desc: 'check local port', cmd: 'echo -n "[port] "; read p; cat < /dev/tcp/127.0.0.1/$p'})
            item({key: 'u', desc: 'system uptime', cmd: 'uptime'})
            item({
                key: 'c',
                desc: 'most used commands',
                cmd: 'echo "history" | bash -i 2>/dev/null | sed "s/^ *[0-9]* *//" | sort | uniq -c | sort -nr | head'
            })
            item({key: 'p', desc: 'public ip address & geolocation', cmd: 'curl http://ipinfo.io'})
            item({
                key: 'i',
                desc: 'internet connection speed',
                cmd: 'wget -O /dev/null http://releases.ubuntu.com/14.04.4/ubuntu-14.04.4-desktop-amd64.iso'
            })
            item({key: 'g', desc: 'ping linux.org (ctrl+c to stop)', cmd: 'ping linux.org'})
            item({key: 's', desc: 'serve current directory on port 8081', cmd: 'python -m SimpleHTTPServer 8081'})
        })

        item({key: 'l', desc: 'log files'}, function () {

            item({key: 'l', desc: 'last updated system log files', cmd: 'ls -lt /var/log | head'})
            item({key: 'a', desc: 'less apache2/error', cmd: 'less +F /var/log/apache2/error.log'})
            _(exec('ls -dt /var/log/*.* | head -n 5').split('\n')).each(function (el, ix) {
                item({key: ix, desc: 'less ' + el, cmd: 'less +F ' + el})
            })
        })
    })
})