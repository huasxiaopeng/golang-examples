package main

import "unicode"

/**
              入参          返回值类型
 */
func nextInt(b[]byte,i int)(int ,int){
	for ; i <len(b)&&!unicode.IsDigit(rune(b[i])) ; i++ {

	}
	x:=0
	for ; i <len(b)&&unicode.IsDigit(rune(b[i])) ; i++ {
		x=x*10+int(b[i])-'0'
	}
	return x,i
}

type Reader struct {


}

func (r Reader) Read(buf []byte) (int, error) {
	return 1,nil
}

/*
                                      指定名称的返回值
 */
func ReadFull(r Reader, buf []byte) (n int, err error) {
	for len(buf) > 0 && err == nil {
		var nr int
		nr, err = r.Read(buf)
		n += nr
		buf = buf[nr:]
	}
	return
}

func main() {

}
