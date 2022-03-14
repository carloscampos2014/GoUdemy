package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Dados Web"
	app.Usage = "Pegar dados de um domínio web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Buscar IPS de endereços da internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Buscar os nomes dos servidores de endereços da internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatalln(erro.Error())
	}

	fmt.Printf("IPs do Host %v:\n", host)

	for _, ip := range ips {
		fmt.Println("  ", ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatalln(erro.Error())
	}

	fmt.Printf("Servidores do Host %v:\n", host)

	for _, server := range servers {
		fmt.Println("  ", server.Host)
	}
}


