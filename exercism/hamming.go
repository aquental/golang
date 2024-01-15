package hamming

import "errors"

func Distance(a, b string) (int, error) {
	la := len(a)
    lb := len(b)
    dist := 0
    
	if la != lb {
        return dist, errors.New("must be same size")
    }
	if la == 0 {
        return dist, nil
    }
	for i:=0; i < la; i++ {
        if a[i] != b[i] {
            dist++
        }
    }
    return dist, nil
}
