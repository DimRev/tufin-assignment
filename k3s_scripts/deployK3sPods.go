package k3sscripts

func (ctx *Context) DeployK3sPods() K3sError {
	err := ctx.GenerateManifests()
	defer ctx.CleanupTempFiles()
	return err
}
