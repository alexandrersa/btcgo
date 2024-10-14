/**
 * App BTCGO
 * Versão Beta
 */

package main

import (
	app "btcgo/cmd/core"
	"btcgo/cmd/utils"
)

func main() {

	version := "v0.7.0"

	utils.ClearConsole()
	utils.Title(version)

	app.NewApp()
}
