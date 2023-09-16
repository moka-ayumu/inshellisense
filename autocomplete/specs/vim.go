// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["vim"] = model.Subcommand{
		Name:        []string{"vim"},
		Description: `Vi IMproved, a programmer's text editor`,
		Args: []model.Arg{{
			Templates: []model.Template{model.TemplateFilepaths},
		}},
		Options: []model.Option{{
			Name:        []string{"-v"},
			Description: `Vi mode (like 'vi')`,
		}, {
			Name:        []string{"-e"},
			Description: `Ex mode (like 'ex')`,
		}, {
			Name:        []string{"-E"},
			Description: `Improved Ex mode`,
		}, {
			Name:        []string{"-s"},
			Description: `Enable silent mode (when in ex mode), or Read Normal mode commands from file`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "scriptin",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"-d"},
			Description: `Diff mode (like 'vimdiff')`,
		}, {
			Name:        []string{"-y"},
			Description: `Easy mode (like 'evim', modeless)`,
		}, {
			Name:        []string{"-R"},
			Description: `Readonly mode (like 'view')`,
		}, {
			Name:        []string{"-Z"},
			Description: `Restricted mode (like 'rvim')`,
		}, {
			Name:        []string{"-m"},
			Description: `Modifications (writing files) not allowed`,
		}, {
			Name:        []string{"-M"},
			Description: `Modifications in text not allowed`,
		}, {
			Name:        []string{"-b"},
			Description: `Binary mode`,
		}, {
			Name:        []string{"-l"},
			Description: `Lisp mode`,
		}, {
			Name:        []string{"-C"},
			Description: `Compatible with Vi: 'compatible'`,
		}, {
			Name:        []string{"-N"},
			Description: `Not fully Vi compatible: 'nocompatible'`,
		}, {
			Name:        []string{"-V"},
			Description: `Be verbose [level N] [log messages to fname]`,
			Args: []model.Arg{{
				Name: "N",
			}, {
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "fname",
			}},
		}, {
			Name:        []string{"-D"},
			Description: `Debugging mode`,
		}, {
			Name:        []string{"-n"},
			Description: `No swap file, use memory only`,
		}, {
			Name:        []string{"-r"},
			Description: `Recover crashed session if filename is specified, otherwise list swap files and exit`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "filename",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"-L"},
			Description: `Same as -r`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "filename",
			}},
		}, {
			Name:        []string{"-T"},
			Description: `Set terminal type to <terminal>`,
			Args: []model.Arg{{
				Name: "terminal",
			}},
		}, {
			Name:        []string{"--not-a-term"},
			Description: `Skip warning for input/output not being a terminal`,
		}, {
			Name:        []string{"--ttyfail"},
			Description: `Exit if input or output is not a terminal`,
		}, {
			Name:        []string{"-u"},
			Description: `Use <vimrc> instead of any .vimrc`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "vimrc",
			}},
		}, {
			Name:        []string{"--noplugin"},
			Description: `Don't load plugin scripts`,
		}, {
			Name:        []string{"-p"},
			Description: `Open N tab pages (default: one for each file)`,
			Args: []model.Arg{{
				Name:       "N",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"-o"},
			Description: `Open N windows (default: one for each file)`,
			Args: []model.Arg{{
				Name:       "N",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"-O"},
			Description: `Like -o but split vertically`,
			Args: []model.Arg{{
				Name:       "N",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"+"},
			Description: `Start at end of file, if line number is specified, start at that line`,
			Args: []model.Arg{{
				Name:       "lnum",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--cmd"},
			Description: `Execute <command> before loading any vimrc file`,
			Args: []model.Arg{{
				Name:      "command",
				IsCommand: true,
			}},
		}, {
			Name:        []string{"-c"},
			Description: `Execute <command> after loading the first file`,
			Args: []model.Arg{{
				Name: "command",
			}},
		}, {
			Name:        []string{"-S"},
			Description: `Source file <session> after loading the first file`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "session",
			}},
		}, {
			Name:        []string{"-w"},
			Description: `Append all typed commands to file <scriptout>`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "scriptout",
			}},
		}, {
			Name:        []string{"-W"},
			Description: `Write all typed commands to file <scriptout>`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "scriptout",
			}},
		}, {
			Name:        []string{"-x"},
			Description: `Edit encrypted files`,
		}, {
			Name:        []string{"--startuptime"},
			Description: `Write startup timing messages to <file>`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "file",
			}},
		}, {
			Name:        []string{"-i"},
			Description: `Use <viminfo> instead of .viminfo`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "viminfo",
			}},
		}, {
			Name:        []string{"--clean"},
			Description: `'nocompatible', Vim defaults, no plugins, no viminfo`,
		}, {
			Name:        []string{"-h", "--help"},
			Description: `Print Help message and exit`,
		}, {
			Name:        []string{"--version"},
			Description: `Print version information and exit`,
		}},
	}
}