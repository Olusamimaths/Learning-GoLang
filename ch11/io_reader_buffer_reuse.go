package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}

	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		// note that error checks are performed after trying to process the non-error value
		// BECAUSE: there might have been bytes returned before the error occurs
		if err == io.EOF {
			return out, nil
		}

		if err != nil {
			return nil, err
		}
	}
}

// using in a decorator pattern
func buildGZipReader(filename string) (*gzip.Reader, func(), error) {
	r, err := os.Open(filename) // os.Open meets the io.Reader interface
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r) // gzip.NewReader also meets the io.Reader interface
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func countInGZip(filename string) (error) {
	gr, closer, err := buildGZipReader(filename)
	if err != nil {
		return err
	}
	defer closer()
	
	counts, err := countLetters(gr)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(counts)
	return nil
}

func main() {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)

	counts, err := countLetters(sr)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(counts)
}