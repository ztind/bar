package bar

import (
	"fmt"
	"strings"
)

type progressBar struct {
	current int //当前下载上传字节数
	total int  //文件总大小
	percent float64 //下载上传百分比
	unitFormat string //打印格式单元
}
func NewProgressBar(total int)*progressBar{
	bar := &progressBar{total:total}
	bar.unitFormat = "▊"
	return bar
}
func NewProgressBarFormat(total int,format string)*progressBar{
	bar := &progressBar{total:total}
	bar.unitFormat = format
	return bar
}
func (this *progressBar)Print(current int){
	this.current = current
	this.percent = float64(this.current) / float64(this.total) * 100
	format := strings.Repeat(this.unitFormat, int(this.percent)/2)
	fmt.Printf("\r[%-50s] %.1f%s  %d/%d",format,this.percent,"%",this.current,this.total) //最大长度50
}

