package main

type Files struct {
	fd int
	name string
	dirinfo int
	nepipe int

}
//对象构造函数初始化
func  newFiles(fd int ,name string)*Files  {
	if fd < 0 {
		return nil
	}
	f := new(Files)
	f.fd = fd
	f.name = name
	f.dirinfo = 0
	f.nepipe = 0
	return f
}
/**
   上面简化
 */
func NewFiles(fd int ,name string)*Files  {
	if fd<0{
		return  nil
	}
	f:=Files{fd,name,0,0}
	return &f
}
/**
  继续简化
 */
func newFiless(fd int,name string) *Files {
	if fd<0{
		return  nil
	}
	return &Files{fd,name,0,0}
}
/**
  接着简化
 */
func newFilesss(fd int,name string)*Files  {
	if fd<0{
		return nil
	}
	return &Files{fd: fd,name: name}
}
func main() {
	a:=new(Files)
	a.fd=0
	a.name="zs"
	a.dirinfo=0
	a.nepipe=0
}
