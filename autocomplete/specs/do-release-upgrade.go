// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["do-release-upgrade"] = model.Subcommand{
		Name:        []string{"do-release-upgrade"},
		Description: `Upgrade Ubuntu to latest release`,
		Options: []model.Option{{
			Name:        []string{"-h", "--help"},
			Description: `Show help message and exit`,
		}, {
			Name:        []string{"-d", "--devel-release"},
			Description: `If using the latest supported release, upgrade to the development release`,
		}, {
			Name:        []string{"-p", "--proposed"},
			Description: `Try upgrading to the latest release using the upgrader from Ubuntu-proposed`,
		}, {
			Name:        []string{"-m"},
			Description: `Run in a special upgrade mode`,
			Args: []model.Arg{{
				Name:        "mode",
				Suggestions: []model.Suggestion{{Name: []string{`desktop`}}, {Name: []string{`server`}}},
			}},
		}, {
			Name:        []string{"--mode"},
			Description: `Run in a special upgrade mode`,
			Args: []model.Arg{{
				Name:        "mode",
				Suggestions: []model.Suggestion{{Name: []string{`desktop`}}, {Name: []string{`server`}}},
			}},
		}, {
			Name:        []string{"-f"},
			Description: `Run the specified frontend`,
			Args: []model.Arg{{
				Name: "frontend",
			}},
		}, {
			Name:        []string{"--frontend"},
			Description: `Run the specified frontend`,
			Args: []model.Arg{{
				Name: "frontend",
			}},
		}},
	}
}