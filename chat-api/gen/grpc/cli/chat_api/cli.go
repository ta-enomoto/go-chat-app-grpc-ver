// Code generated by goa v3.4.3, DO NOT EDIT.
//
// chat api gRPC client CLI support package
//
// Command:
// $ goa gen chat-api/design

package cli

import (
	chatapic "chat-api/gen/grpc/chatapi/client"
	"flag"
	"fmt"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `chatapi (getchat|postchat)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` chatapi getchat --message '{
      "id": 3009151431332760301
   }' --key "Aspernatur et nihil ab sint voluptas."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		chatapiFlags = flag.NewFlagSet("chatapi", flag.ContinueOnError)

		chatapiGetchatFlags       = flag.NewFlagSet("getchat", flag.ExitOnError)
		chatapiGetchatMessageFlag = chatapiGetchatFlags.String("message", "", "")
		chatapiGetchatKeyFlag     = chatapiGetchatFlags.String("key", "REQUIRED", "")

		chatapiPostchatFlags       = flag.NewFlagSet("postchat", flag.ExitOnError)
		chatapiPostchatMessageFlag = chatapiPostchatFlags.String("message", "", "")
		chatapiPostchatKeyFlag     = chatapiPostchatFlags.String("key", "REQUIRED", "")
	)
	chatapiFlags.Usage = chatapiUsage
	chatapiGetchatFlags.Usage = chatapiGetchatUsage
	chatapiPostchatFlags.Usage = chatapiPostchatUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "chatapi":
			svcf = chatapiFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "chatapi":
			switch epn {
			case "getchat":
				epf = chatapiGetchatFlags

			case "postchat":
				epf = chatapiPostchatFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "chatapi":
			c := chatapic.NewClient(cc, opts...)
			switch epn {
			case "getchat":
				endpoint = c.Getchat()
				data, err = chatapic.BuildGetchatPayload(*chatapiGetchatMessageFlag, *chatapiGetchatKeyFlag)
			case "postchat":
				endpoint = c.Postchat()
				data, err = chatapic.BuildPostchatPayload(*chatapiPostchatMessageFlag, *chatapiPostchatKeyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// chatapiUsage displays the usage of the chatapi command and its subcommands.
func chatapiUsage() {
	fmt.Fprintf(os.Stderr, `The service performs get chat.
Usage:
    %s [globalflags] chatapi COMMAND [flags]

COMMAND:
    getchat: Getchat implements getchat.
    postchat: Postchat implements postchat.

Additional help:
    %s chatapi COMMAND --help
`, os.Args[0], os.Args[0])
}
func chatapiGetchatUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] chatapi getchat -message JSON -key STRING

Getchat implements getchat.
    -message JSON: 
    -key STRING: 

Example:
    `+os.Args[0]+` chatapi getchat --message '{
      "id": 3009151431332760301
   }' --key "Aspernatur et nihil ab sint voluptas."
`, os.Args[0])
}

func chatapiPostchatUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] chatapi postchat -message JSON -key STRING

Postchat implements postchat.
    -message JSON: 
    -key STRING: 

Example:
    `+os.Args[0]+` chatapi postchat --message '{
      "Chat": "Aliquam dolor necessitatibus.",
      "Cookie": "Quaerat est est aut.",
      "Id": "Neque est.",
      "Member": "Reiciendis quae reiciendis aperiam atque minus.",
      "PostDt": "2000-03-10T00:55:40Z",
      "RoomName": "Fugiat excepturi quo accusamus sed asperiores nihil.",
      "UserId": "Saepe voluptas possimus voluptas."
   }' --key "Ut quia est."
`, os.Args[0])
}
