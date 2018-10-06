package translator

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"
)

type parameters struct {
	source string
	dest   string
	text   string
}

/// entry point for console app
func RunTranslateApp(args []string) {
	app := cli.NewApp()

	app.Name = "Translator"
	app.Usage = "Simple translator"
	app.Flags = generateFlags()

	app.Action = func(c *cli.Context) error {
		params, err := parseParameters(c)
		if err != nil {
			return err
		}
		return run(params)
	}

	err := app.Run(args)
	if err != nil {
		fmt.Errorf("An error occured during translate %v", err)
	}
}

func run(params parameters) error {
	translator := DefaultTranslator()
	results, err := translator.Translate(params.source, params.dest, params.text)
	if err != nil {
		return err
	}
	for _, result := range results {
		fmt.Println(result.word)
	}

	return nil
}

// TODO: move to separate class
func parseParameters(c *cli.Context) (params parameters, err error) {
	if !c.Args().Present() {
		err = errors.New("Expecting text to translate")
		return
	}

	source := c.String("source")
	text := c.Args().First()

	if source == EN {
		params = parameters{source, PL, text}

	} else if source == PL {
		params = parameters{source, EN, text}

	} else {
		err = errors.New("Invalid source, only {pl, en} supported")
	}

	return
}

func generateFlags() []cli.Flag {

	return []cli.Flag{
		cli.StringFlag{
			Name:  "source, s",
			Value: "en",
			Usage: "source word lang",
		},
	}
}
