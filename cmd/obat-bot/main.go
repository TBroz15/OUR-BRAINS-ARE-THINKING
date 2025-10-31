// OUR BRAINS ARE THINKING ‼️‼️
package main

import (
	"github.com/TBroz15/OUR-BRAINS-ARE-THINKING/internals/helpers"
	"github.com/charmbracelet/log"
)

func main() {
	isApproved := helpers.HasTheWords("i like how this short help think with our brains lmao")

	log.Info(isApproved)
}
