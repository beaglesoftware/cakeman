package cmd

func initCommands() {
	RegisterCommand("module", HandleModule)
	RegisterCommand("cake", HandleCake)
	RegisterCommand("build", HandleBuild)
}
