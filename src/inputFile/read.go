package inputFile

type fileCommands struct {
	fileOperations FileOperations
}

func (com *fileCommands) ReadCommands(filePath string) ([][]string, error) {
	fo := com.fileOperations
	file, err := fo.open(filePath)
	if err != nil {
		return nil, err
	}

	defer fo.close(file)

	argLists := fo.read(file)
	// fmt.Println(argLists)

	return argLists, nil
}
