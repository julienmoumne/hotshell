# Tips

## Command-line usage

  - Implicitly load the current directory's hotshell, `./hs.js`, or if not found, the system-wide hotshell `~/.hs/hs.js`
```bash
hs
```
  - Specify the path to the definition file
```bash
hs -f ~/projects/web/hs.js
# or
hs -f ~/projects/web
```
  - Load a menu remotely
```bash
hs -f https://raw.githubusercontent.com/julienmoumne/hs/v0.1.0/hs.js
```
  - Set the working directory to the location of the menu definition
```bash
hs --chdir -f ~/projects/web/hs.js
```
  - Use aliases
```bash
alias hsdocker="hs -f https://raw.githubusercontent.com/julienmoumne/hs/v0.1.0/examples/docker/docker.hs.js"
```
  - Generate an interactive HTML demo of your menus, [example](https://julienmoumne.github.com/hs/demos/hs.js.html)
```bash
hs --generate-demo -f ~/projects/web/hs.js > hotshell-web-demo.html  
```

## Menu definition
  
  - Prompt for inputs to be provided as command arguments
```bash
// prompts for a location and a pattern and triggers a grep search
item({
  desc: 'find text in files',
  cmd: 'echo -n "[location] [pattern] " ' + // prompt for inputs
         '&& read l p ' + // read inputs
         '&& grep -rnws $l -e $p' // provide inputs as command arguments
})
```
  - Access other interactive applications
```javascript
item({cmd: 'ssh remote-server'})
item({cmd: 'sudo vim /etc/hosts'})
```
  - When running out of characters for defining hot keys, use capital letters
```javascript
item({key: 'S', cmd: 'ssh remote-server'})
```
  - Include menus defined in separate files, see [composite example](examples#composite)  
  - There is a good number of command examples in the default hotshell
```bash
hs --default
```
- The DSL defined by Hotshell uses some JavaScript tricks, learn more about it : http://alexyoung.org/2009/10/22/javascript-dsl/