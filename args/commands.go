package args

type Command struct {
	Name        CommandName
	Description string
	Flags       []Flag
}

type Flag struct {
	Long        string
	Short       string
	Description string
}

type CommandName string

var (
	GlobalCommand  CommandName = ""
	ClusterCommand CommandName = "cluster"
	DeployCommand  CommandName = "deploy"
	StatusCommand  CommandName = "status"
	RemoveCommand  CommandName = "remove"
)

var Commands = []Command{
	{
		Name:        GlobalCommand,
		Description: "",
		Flags: []Flag{
			{
				Long:        "--help",
				Short:       "-h",
				Description: "Print any command",
			},
			{
				Long:        "--version",
				Short:       "-v",
				Description: "Print the version of the program",
			},
		},
	},
	{
		Name:        ClusterCommand,
		Description: "Deploy a k3s Kubernetes cluster",
		Flags: []Flag{
			{
				Long:        "--help",
				Short:       "-h",
				Description: "Print details about this command",
			},
		},
	},
	{
		Name:        DeployCommand,
		Description: "Deploy two pods: MySQL and WordPress",
		Flags: []Flag{
			{
				Long:        "--help",
				Short:       "-h",
				Description: "Print details about this command",
			},
		},
	},
	{
		Name:        StatusCommand,
		Description: "Print the status table of pods in the default namespace",
		Flags: []Flag{
			{
				Long:        "--help",
				Short:       "-h",
				Description: "Print details about this command",
			},
		},
	},
	{
		Name:        RemoveCommand,
		Description: "Remove the k3s cluster",
		Flags: []Flag{
			{
				Long:        "--help",
				Short:       "-h",
				Description: "Print details about this command",
			},
		},
	},
}
