package hotshell

type Options struct {
	Default      bool   `long:"default" description:"Load the default menu to get to know Hotshell"`
	File         string `short:"f" long:"file" description:"Specify an alternate definition file (defaults: ./hs.js, ~/.hs/hs.js)" value-name:"FILE|DIR"`
	Chdir        bool   `long:"chdir" description:"Set the working directory to the location of the menu definition"`
	GenerateDemo bool   `long:"generate-demo" description:"Generate an interactive HTML demo to stdout"`
	GenerateDoc  bool   `long:"generate-doc" description:"Generate a markdown documentation of the menu to stdout"`
	Version      bool   `short:"v" long:"version" description:"Print version information and quit"`
}
