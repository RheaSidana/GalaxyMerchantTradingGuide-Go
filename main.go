package main

import (
	"fmt"
	"galaxy_merchant_trading_guide/src/commandLine"
	"galaxy_merchant_trading_guide/src/commands"

	// "galaxy_merchant_trading_guide/src/commands"
	"galaxy_merchant_trading_guide/src/inputFile"
	romannumerals "galaxy_merchant_trading_guide/src/romanNumerals"
)

func main() {
	filePath, err := commandLine.ParseCommandLineArgs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	commandsList, err := inputFile.ReadFrom(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	romannumerals.BuildRomanNumerals()
	commands.CreateCommands()
	commands.Execute(commandsList)
}