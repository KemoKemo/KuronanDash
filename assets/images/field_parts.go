package images

func loadFieldparts() error {
	var err error

	TilePrairie, err = loadSingleImage(tilePrairie_png)
	if err != nil {
		return err
	}
	MountainNear, err = loadSingleImage(mountainNear_png)
	if err != nil {
		return err
	}
	MountainFar, err = loadSingleImage(mountainFar_png)
	if err != nil {
		return err
	}

	return nil
}