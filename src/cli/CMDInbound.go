package cli

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
	"github.com/spf13/cobra"

	utils "../utils"
)

var port int
var sources bool

func init () {
	CMDInbound.Flags().IntVarP(&port, "port", "p", -1, "Display only connections on a specified port")
	CMDInbound.Flags().BoolVarP(&sources, "sources", "s", false, "Show amount of connection for each address.")
}

func getActiveConnections () []string {
	var output string = utils.Shell("netstat", "-ant")
	var lines = strings.Split(strings.TrimSuffix(output, "\n"), "\n")
	var testFilter = func (s string) bool { 
		if port >= 1 {
			return strings.Contains(s, ":" + strconv.Itoa(port))
		} else { 
			return true
		}
	}

	return utils.Filter(lines, testFilter)
}

func printStatus (connections []string) {
	var statuses = make(map[string]int)

	for _, line := range connections {
		space := regexp.MustCompile(`\s+`)
		var raw = space.ReplaceAllString(line, " ")
		var connection = strings.Fields(raw) 

		var status = connection[len(connection) - 1]

		if strings.ToUpper(status) == status {
			statuses[status]++
		}
	}

	for status, amount := range statuses {
		var line = fmt.Sprintf("\t> %s: %s", status, utils.FormatAmount(amount, 5, 50, 100))
		Println(line)
	}

	var line = fmt.Sprintf("\n\t> TOTAL: %s", utils.FormatAmount(len(connections), 5, 50, 100))
	Println(line)
}

func printSources (connections []string) {
	var addresses = make(map[string]int)
	for _, line := range connections {
		space := regexp.MustCompile(`\s+`)
		var raw = space.ReplaceAllString(line, " ")
		var connection = strings.Fields(raw)

		var source = connection[len(connection) - 2]
		var address = strings.Split(source, ":")[0]

		if address != "127.0.0.1" {
			addresses[address]++;
		}
	}

	for address, amount := range addresses {
		var line = fmt.Sprintf("\t> %s: {cyan}%d", address, amount)
		Println(line)
	}
}

var CMDInbound = &cobra.Command{
	Use:   "inbound",
	Short: "Gets the currently incoming connections.",
	Long: `Displays a list of incoming connections ordered by their network status.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var connections = getActiveConnections()
		
		if sources {
			printSources(connections)
		} else {
			printStatus(connections)
		}
	},
}