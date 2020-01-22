package lib

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func checkBadChar(body []byte) (pass bool) {
	for _, b := range body {
		if b < 65 || b > 90 {
			return
		}
	}
	pass = true
	return
}

func checkX(body []byte) (pass bool) {
	m := len(body) / 10
	n := bytes.Count(body, []byte{'X'})
	if n < m {
		pass = true
	}
	return
}

func checkCompositionBias(body []byte) (pass bool) {
	m := make(map[byte]uint16)
	for _, b := range body {
		m[b]++
	}
	var max1, max2 uint16
	for _, v := range m {
		if v > max1 {
			max1, max2 = v, max1
			continue
		}
		if v > max2 {
			max2 = v
		}
	}
	if int(max1 + max2) * 2 < len(body) {
		pass = true
	}
	return
}

// Work it out
func Work(input, output *os.File, lengthLimit int) error {
	outBuf := bufio.NewWriter(output)

	rec := NewRecord()
	var line []byte

	var p, q, r, t int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line = scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		if line[0] == '>' {
			switch {
			case len(rec.Body) < lengthLimit:
				if len(rec.Head) > 0 {
					p++
					log.Printf("Record ignored [less than %d aa]: %s", lengthLimit, string(rec.Head))
					rec.Reset()
				}
			case !checkBadChar(rec.Body):
				t++
				log.Print("Record ignored [bad character found]: ", string(rec.Head))
			case !checkX(rec.Body):
				q++
				log.Print("Record ignored [X more than 10%]: ", string(rec.Head))
				rec.Reset()
			case !checkCompositionBias(rec.Body):
				r++
				log.Print("Record ignored [aa composition bias]: ", string(rec.Head))
				rec.Reset()
			default:
				_, err := rec.WriteTo(outBuf)
				if err != nil {
					return err
				}
			}
			rec.LoadHead(line)
			continue
		}
		rec.LoadBody(line)
	}

	_, err := rec.WriteTo(outBuf)
	if err != nil {
		return err
	}

	if err = outBuf.Flush(); err != nil {
		return err
	}

	println("------------------------------------------------")
	println("Done.")
	println(p, " records ignored due to length < 50.")
	println(t, " records ignored due to bad character (out of range [A-Z]).")
	println(q, " records ignored due to too many X.")
	println(r, " records ignored due to aa composition bias.")
	return nil
}