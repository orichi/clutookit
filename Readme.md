### 小工具
* 生成密码，3级复杂度，长度定制
* 摩斯密码转换

#### Usage
```
Usage of ./small_tool:
  -f string
    	pwd1: easy password
    	pwd2: normal password
    	pwd3: hard password
    	morse: 摩斯电码 (default "pwd1")
  -len int
    	default is 8, max is decided by pwd complex level (default 8)
  -s string
    	data be morse encode
```
##### 生成密码
```

> small_tool -f=pwd1 -len=8
time:	2018-12-27 10:26:57
data:	yj2nquk9
```

##### 摩斯电码
```
>small_tool.go -f=morse -s="ABCDEF"
time:	2018-12-27 10:27:40
data:	ABCDEF=>	.- -... -.-. -.. . ..-.
time:	2018-12-27 10:27:40
data:	.- -... -.-. -.. . ..-.=>	ABCDEF
```