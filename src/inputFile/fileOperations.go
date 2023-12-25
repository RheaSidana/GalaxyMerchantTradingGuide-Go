package inputFile

func getFileCommands() fileCommands {
	return fileCommands{
		fileOperations: NewFileOperations(),
	}
}

func ReadFrom(filepath string) ([][]string, error) {
	fc := getFileCommands()
	return fc.ReadCommands(filepath)
}
