package bar

import (
	"testing"
	"time"
)

func TestPrint(t *testing.T){
	bar := NewProgressBar(150)
	for i:=0;i<=150;i++{
		time.Sleep(time.Millisecond*100)
		bar.Print(i)
	}
}
