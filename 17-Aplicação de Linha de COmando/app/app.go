package app

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet"

	flags := []cli.Flag{
				cli.StringFlag {
					Name: "host",
					Value: "devbook.com.br",
				}
			}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores"
			Usage: "Busca o nome do servidores na internet"
			Flags: flags,
			Action: buscarServidores
		}
	}

	return app
}

func buscarIps(c *cli.Context) { 
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	
	for _, servidor := range servidores{
		fmt.Println(servidor.host)
	}
}

// COMANDOS:

// go run main.go ip --parametro1 valor1 --parametro2 valor2"log"
// go run main.go ip --host amazon.com.br
