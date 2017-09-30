package proxy

import (
	"os"
	"os/signal"
	"testing"
)

func TestP(t *testing.T) {
	s := make(chan os.Signal)
	signal.Notify(s)
	<-s
}
