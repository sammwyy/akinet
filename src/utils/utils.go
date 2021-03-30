package utils

import (
	"fmt"
	// "sort"
	"log"
    "os/exec"
)

func Shell (command string, arg string) string {
    out, err := exec.Command(command, arg).Output()

    if err != nil {
        log.Fatal(err)
    }

    return string(out)
}

func FormatAmount (amount int, safe int, medium int, dangerus int) string {
	if amount <= safe {
		return fmt.Sprintf("{green}%d", amount)
	} else if amount <= medium {
		return fmt.Sprintf("{yellow}%d", amount)
	} else if amount <= dangerus {
		return fmt.Sprintf("{l_red}%d", amount)
	} else {
		return fmt.Sprintf("{red}%d", amount)
	}
}

func Filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}

	return
}


/*

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func SortMap (m map[string] int) map[string] int {
	output := make(map[string] int)
	p := make(PairList, len(m))

	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	
	for _, k := range p {
		output[k.Key] = k.Value
	}

	return output
}*/