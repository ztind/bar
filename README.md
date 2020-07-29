###### ui 
[##########################                        ] 53.3%  80/150

[##################################################] 100.0%  150/150

[>>>>>>>>>>>>>>>>>>>>>>>>>>                        ] 53.3%  80/150

[>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>] 100.0%  150/150

###### use demo 
```
bar := NewProgressBar(150)  
for i:=0;i<=150;i++{
    time.Sleep(time.Millisecond*100)
    bar.Print(i)
}
```
######  use custom print formats
```
bar := NewProgressBarFormat(150,">")  
for i:=0;i<=150;i++{
    time.Sleep(time.Millisecond*100)
    bar.Print(i)
}
```


