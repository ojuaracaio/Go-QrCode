package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	var opcao int
	println("Bem-vindo ao gerador de QRCode!")
	//cria um loop infinito com o gerador
	for k := 0; k == 0; {
		println("\nQual tipo de código deseja gerar?")
		println("\n1-Texto Simples")
		println("2-Rede Wifi")
		println("3-sair\n")
		println("opção: ")
		fmt.Scanln(&opcao)
		if opcao == 1 {
			println("Digite o texto que deseja codificar: ")
			in := bufio.NewReader(os.Stdin)
			input, _ := in.ReadString('\n')
			input = input[:len(input)-2]
			imprimir(gerarTextoSimples(input))
			criarArquivoSimples(input)
		} else if opcao == 2 {
			println("Digite o nome da rede wifi: ")
			in := bufio.NewReader(os.Stdin)
			inputRede, _ := in.ReadString('\n')
			println("Digite a senha: ")
			inputSenha, _ := in.ReadString('\n')
			inputRede = inputRede[:len(inputRede)-2]
			inputSenha = inputSenha[:len(inputSenha)-2]
			imprimir(gerarTextoSimples(formataWifi(inputRede, inputSenha)))
			criarArquivoWifi(formataWifi(inputRede, inputSenha), inputRede)
		} else if opcao == 3 {
			break
		} else {
			println("opção inválida.")
			fmt.Scanln(opcao)
		}
		//fmt.Scanln()
		//cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		//cmd.Stdout = os.Stdout
		//cmd.Run()
		//fmt.Scanln()
	}

}

//gera o qrcode como uma lista de listas de valores booleanos (preto=1 e branco=0)
func gerarTextoSimples(texto string) [][]bool {
	img, _ := qrcode.New(texto, qrcode.Low)
	//println(img)
	return img.Bitmap()
}

//formata a rede e senha numa string padrão para o qrcode
func formataWifi(rede string, senha string) string {
	wifi := "WIFI:S:" + rede + ";T:WPA;P:" + senha + ";;"
	return wifi
}

//imprime a lista de listas usando emojis quadrados para representar os 0s e 1s
func imprimir(png [][]bool) {
	for i := 0; i < len(png); i++ {
		for j := 0; j < len(png[i]); j++ {
			if png[i][j] {
				//print("⬛")
				//print(" ")
				print("▓▓")
			} else {
				//print("⬜")
				//print("■")
				print("  ")
			}
		}
		print("\n")
	}
}

//cria um arquivo png com o qrcode gerado
func criarArquivoSimples(texto string) {
	err := qrcode.WriteFile(texto, qrcode.Low, 256, texto+".png")
	if err != nil {
		print("\nHouve um erro ao criar o arquivo png.\n")
	} else {
		println("Um arquivo png foi criado com seu código.\n")
	}
}

//cria um arquivo png com o nome da rede (assim a senha fica oculta)
func criarArquivoWifi(texto string, titulo string) {
	err := qrcode.WriteFile(texto, qrcode.Low, 256, titulo+".png")
	if err != nil {
		print("\nHouve um erro ao criar o arquivo png.\n")
	} else {
		println("Um arquivo png foi criado com seu código.\n")
	}
}
