package apc_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/jeromedoucet/apc"
)

func TestFactoryOptimums(t *testing.T) {
	// given
	var res []byte
	var err error
	expectedRes := `
Nr solutions=4
10 1
2 4 5
1 4 6
5 6
Waste=0
	`
	output := "./output_1_test"
	conf, err1 := apc.NewConf("./input_1_test")

	fmt.Println(err1)

	factory := apc.NewFactory(conf)
	defer os.Remove(output)

	// when
	factory.Optimums(output)

	// then
	res, err = ioutil.ReadFile(output)

	if err != nil {
		t.Fatal("Expect no error, got ", err)
	}

	if string(res) != expectedRes {
		t.Fatalf("Expect %s, got %s", expectedRes, string(res))
	}

}
