package initializers

func InitApp() error {
	if err := LoadEnvVariables(); err != nil {
		return err
	}
	if err := ConnectToDB(); err != nil {
		return err
	}
	return nil
}