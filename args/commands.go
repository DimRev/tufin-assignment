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
	HasArg      bool
}

type CommandName string

var (
	GlobalCommand  CommandName = ""
	ClusterCommand CommandName = "cluster"
	DeployCommand  CommandName = "deploy"
	StatusCommand  CommandName = "status"
	RemoveCommand  CommandName = "remove"
)

type ExecutionFunc func(flags map[string]string) error

var Commands = []Command{
	{
		Name:        GlobalCommand,
		Description: "",
		Flags: []Flag{
			{
				Long:        "--help",
				Short:       "-h",
				Description: "Print any command",
				HasArg:      false,
			},
			{
				Long:        "--version",
				Short:       "-v",
				Description: "Print the version of the program",
				HasArg:      false,
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
				HasArg:      false,
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
				HasArg:      false,
			},
			{
				Long:        "--helm",
				Short:       "-e",
				Description: "Deploy the helm chart",
				HasArg:      false,
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
				HasArg:      false,
			},
			{
				Long:        "--namespace",
				Short:       "-n",
				Description: "The namespace to use",
				HasArg:      true,
			},
			{
				Long:        "--pod",
				Short:       "-p",
				Description: "Show the status of the pods in the given(or default) namespace. If this flag is not set, the status of all pods in the namespace will be shown.",
				HasArg:      false,
			},
			{
				Long:        "--service",
				Short:       "-s",
				Description: "Show the status of the services in the given(or default) namespace. If this flag is not set, the status of all services in the namespace will be shown.",
				HasArg:      false,
			},
			{
				Long:        "--volume",
				Short:       "-v",
				Description: "Show the status of the PersistentVolumeClaims in the given(or default) namespace. If this flag is not set, the status of all PVCs in the namespace will be shown.",
				HasArg:      false,
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
				HasArg:      false,
			},
		},
	},
}
