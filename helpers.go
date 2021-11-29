package vegam

import "os"

func (v *Vegam) CreateDirIfNotExist(path string) error {

	const mode = 0755

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, mode)
		if err != nil {
			return err
		}
	}

	return nil
}
