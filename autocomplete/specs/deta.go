// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["deta"] = model.Subcommand{
		Name:        []string{"deta"},
		Description: `Deta CLI for managing Deta Micros`,
		Options: []model.Option{{
			Name:        []string{"-h"},
			Description: `Show help for deta`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"login"},
			Description: `Trigger the login process for the Deta CLI`,
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for login`,
			}},
		}, {
			Name:        []string{"version"},
			Description: `Print the Deta version`,
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for version`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"upgrade"},
				Description: `Upgrade Deta CLI version`,
				Options: []model.Option{{
					Name:        []string{"-h"},
					Description: `Show help for upgrade`,
				}, {
					Name:        []string{"-v"},
					Description: `Upgrade CLI to specific version`,
					Args: []model.Arg{{
						Name: "Version number",
					}},
				}},
			}},
		}, {
			Name:        []string{"projects"},
			Description: `List Deta projects`,
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for projects`,
			}},
		}, {
			Name:        []string{"new"},
			Description: `Create a new Deta Micro`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to new directory for the micro`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for new`,
			}, {
				Name:        []string{"-n"},
				Description: `Create a micro with Node (node14.x) runtime`,
			}, {
				Name:        []string{"-p"},
				Description: `Create a micro with Python (python 3.9) runtime`,
			}, {
				Name:        []string{"--name"},
				Description: `Set the name of the new micro`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Name of the new micro`,
				}},
			}, {
				Name:        []string{"--project"},
				Description: `Set the project under which the micro is created`,
				Args: []model.Arg{{
					Name:        "project",
					Description: `Name of the existing project`,
				}},
			}, {
				Name:        []string{"--runtime"},
				Description: `Create a micro with a specified runtime`,
				Args: []model.Arg{{
					Name:        "runtime",
					Description: `The selected runtime`,
					Generator:   nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"deploy"},
			Description: `Deploy new code to a Deta Micro`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "path",
				Description: `Path to project directory`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for deploy`,
			}},
		}, {
			Name:        []string{"details"},
			Description: `Get detailed information about a specific Deta micro`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "path",
				Description: `Path to project directory`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for details`,
			}},
		}, {
			Name:        []string{"watch"},
			Description: `Auto-deploy locally saved changes in real time to your Deta micro`,
			Args: []model.Arg{{
				Templates:   []model.Template{model.TemplateFolders},
				Name:        "path",
				Description: `Path to project directory`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for watch`,
			}},
		}, {
			Name:        []string{"auth"},
			Description: `Change auth settings for a Deta Micro`,
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for auth`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"disable"},
				Description: `Disable HTTP Auth for a Deta Micro`,
				Options: []model.Option{{
					Name:        []string{"-h"},
					Description: `Show help for auth disable`,
				}},
			}, {
				Name:        []string{"enable"},
				Description: `Enable HTTP Auth for a Deta Micro`,
				Options: []model.Option{{
					Name:        []string{"-h"},
					Description: `Show help for auth enable`,
				}},
			}, {
				Name:        []string{"create-api-key"},
				Description: `Create an API key for a Deta Micro`,
				Options: []model.Option{{
					Name:        []string{"-h"},
					Description: `Show help for auth create-api-key`,
				}, {
					Name:        []string{"-d"},
					Description: `Set the api-key description`,
					Args: []model.Arg{{
						Name:        "description",
						Description: `The api-key description`,
					}},
				}, {
					Name:        []string{"-n"},
					Description: `Set the api-key name`,
					Args: []model.Arg{{
						Name:        "name",
						Description: `The api-key name`,
					}},
				}, {
					Name:        []string{"-o"},
					Description: `Set the api-key output file`,
					Args: []model.Arg{{
						Templates:   []model.Template{model.TemplateFilepaths},
						Name:        "outfile",
						Description: `The api-key output file`,
					}},
				}},
			}, {
				Name:        []string{"delete-api-key"},
				Description: `Delete an API key for a Deta Micro`,
				Options: []model.Option{{
					Name:        []string{"-h"},
					Description: `Show help for auth delete-api-key`,
				}, {
					Name:        []string{"-n"},
					Description: `Set the api-key name`,
					Args: []model.Arg{{
						Name:        "name",
						Description: `The api-key name`,
					}},
				}},
			}},
		}, {
			Name:        []string{"pull"},
			Description: `Pull the latest deployed code of a Deta Micro to your local machine`,
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for pull`,
			}, {
				Name:        []string{"-f"},
				Description: `Force the overwrite of existing files`,
			}},
		}, {
			Name:        []string{"clone"},
			Description: `Clone a Deta Micro`,
			Args: []model.Arg{{
				Name:        "path",
				Description: `Path to new directory for the clone`,
				IsOptional:  true,
			}},
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for clone`,
			}, {
				Name:        []string{"-n"},
				Description: `The name of the micro to be cloned`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `Name of the micro`,
				}},
			}, {
				Name:        []string{"-p"},
				Description: `The name of the project with the micro to be cloned`,
				Args: []model.Arg{{
					Name:        "project",
					Description: `Name of the project`,
				}},
			}},
		}, {
			Name:        []string{"update"},
			Description: `Update a Deta Micro's name or environment variables`,
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for pull`,
			}, {
				Name:        []string{"-n"},
				Description: `The new name of the micro`,
				Args: []model.Arg{{
					Name:        "name",
					Description: `New name for the micro`,
				}},
			}, {
				Name:        []string{"-r"},
				Description: `The new runtime of the micro`,
				Args: []model.Arg{{
					Name:        "runtime",
					Description: `New runtime for the micro`,
					Generator:   nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"-e"},
				Description: `The new env file of the micro`,
				Args: []model.Arg{{
					Templates:   []model.Template{model.TemplateFilepaths},
					Name:        "env",
					Description: `Path to env file`,
				}},
			}},
		}, {
			Name:        []string{"visor"},
			Description: `Change the Visor settings for a Deta Micro`,
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for visor`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"open"},
				Description: `Open Micro's visor page in the browser`,
				Options: []model.Option{{
					Name:        []string{"-h"},
					Description: `Show help for visor open`,
				}},
			}, {
				Name:        []string{"enable"},
				Description: `Enable Visor for a Deta Micro`,
				Options: []model.Option{{
					Name:        []string{"-h"},
					Description: `Show help for visor enable`,
				}},
			}, {
				Name:        []string{"disable"},
				Description: `Disable Visor for a Deta Micro`,
				Options: []model.Option{{
					Name:        []string{"-h"},
					Description: `Show help for visor disable`,
				}},
			}},
		}, {
			Name:        []string{"run"},
			Description: `Run a Deta Micro from the CLI`,
			Args: []model.Arg{{
				Name:        "action",
				Description: `The action to be performed on the micro. See docs for full examples and details`,
			}},
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for run`,
			}, {
				Name:        []string{"-l"},
				Description: `Show the micro logs`,
			}},
		}, {
			Name:        []string{"cron"},
			Description: `Change cron settings for a Deta Micro`,
			Options: []model.Option{{
				Name:        []string{"-h"},
				Description: `Show help for cron`,
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"set"},
				Description: `Set Deta Micro to run on a schedule`,
				Args: []model.Arg{{
					Name:        "expression",
					Description: `The cron expression to be set`,
				}},
				Options: []model.Option{{
					Name:        []string{"-h"},
					Description: `Show help for cron set`,
				}},
			}, {
				Name:        []string{"remove"},
				Description: `Remove a schedule from a Deta Micro`,
				Args: []model.Arg{{
					Name:        "expression",
					Description: `The cron expression to be removed`,
				}},
				Options: []model.Option{{
					Name:        []string{"-h"},
					Description: `Show help for cron remove`,
				}},
			}},
		}},
	}
}