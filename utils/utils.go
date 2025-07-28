package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(s any) {
	r, _ := json.MarshalIndent(s, "", "\t")
	fmt.Print(string(r))
}
