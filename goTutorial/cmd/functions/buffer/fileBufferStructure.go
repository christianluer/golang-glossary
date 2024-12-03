package buffer

import "io"

type Person struct {
	First string
}

func (p Person) WriteOut(w io.Writer) {
	w.Write([]byte(p.First))
}

// file, err := os.Create("output.txt")
// if err != nil {
// 	log.Fatalf("error %s", err)
// }
// defer file.Close()
// var b bytes.Buffer
// p := buffer.Person{
// 	First: "tomas",
// }
// p.WriteOut(file)
// p.WriteOut(&b)
