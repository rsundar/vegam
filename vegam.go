package vegam

const version = "1.0.0"

type Vegam struct {
	AppName string
	Debug   bool
	Version string
}

func (v *Vegam) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handlers", "migrations", "views", "models", "public", "tmp", "logs", "middleware"},
	}

	err := v.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

func (v *Vegam) Init(p initPaths) error {
	//create the folder structure if it does not already exist
	root := p.rootPath

	for _, path := range p.folderNames {
		err := v.CreateDirIfNotExist(root + "/" + path)
		if err != nil {
			return err
		}
	}

	return nil
}
