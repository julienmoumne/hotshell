package options

type Options struct {
	Default      bool   `long:"default" description:"Load the default menu to get to know Hotshell"`
	File         string `short:"f" long:"file" description:"Specify an alternate definition file (defaults: ./hs.js, ~/.hs/hs.js)" value-name:"FILE|DIR"`
	GenerateDemo bool   `long:"generate-demo" description:"Generate an interactive HTML demo to stdout"`
	GenerateMd   bool   `long:"generate-md" description:"Generate a markdown documentation of the menu to stdout"`
	Version      bool   `short:"v" long:"version" description:"Print version information and quit"`
}
