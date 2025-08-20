package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ramoncl001/comet-cli/generator"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "comet",
		Short: "CLI tool for 'comet' projects",
	}

	var createCmd = &cobra.Command{
		Use:   "new [project-name] [module-name]",
		Short: "Creates a new comet project with selected name",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			module := args[1]
			err := generator.CreateProject(projectName, module)
			if err != nil {
				fmt.Printf("Error creating project: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Project created successfully!")
		},
	}

	var componentCmd = &cobra.Command{
		Use:   "add [type] [name] [location]",
		Short: "Creates a <type - (controller, service, middleware)> with <name> in the selected <location>",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			component := args[0]
			name := args[1]

			location := ""
			if len(args) > 2 {
				location = args[2]
			}

			switch component {
			case "controller":
				if err := generator.CreateController(name, location); err != nil {
					fmt.Printf("Error creating component: %v\n", err)
					os.Exit(1)
				}
			case "service":
				if err := generator.CreateController(name, location); err != nil {
					fmt.Printf("Error creating component: %v\n", err)
					os.Exit(1)
				}
			case "middleware":
				if err := generator.CreateMiddleware(name, location); err != nil {
					fmt.Printf("Error creating component: %v\n", err)
					os.Exit(1)
				}
			default:
				fmt.Printf("Invalid component")
				os.Exit(1)
			}

			fmt.Println("Project created successfully!")
		},
	}

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the selected comet project",
		Args:  cobra.ExactArgs(0),
		Run: func(_ *cobra.Command, args []string) {
			cmd := exec.Command("go", "run", ".")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Error running project: %s - %s", err.Error(), string(output))
				os.Exit(1)
			}
		},
	}

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(componentCmd)
	rootCmd.AddCommand(runCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
