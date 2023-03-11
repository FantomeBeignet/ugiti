package utils

import (
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var logger log.Logger
var once sync.Once

func initLogger() log.Logger {
	log.FatalLevelStyle = lipgloss.NewStyle().
		SetString("ERROR!!").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("4")).
		Foreground(lipgloss.Color("0"))
    logger := log.New()
    return logger
}

func GetLogger() log.Logger {
    once.Do(func() {
        logger = initLogger()
    })
    return logger
}