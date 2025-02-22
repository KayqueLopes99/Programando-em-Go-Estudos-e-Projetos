package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"errors"
)

// Estrutura que define uma pergunta do quiz
type Question struct {
	Text    string   
	Options []string 
	Answer  int      
}

// Estrutura que define o estado do jogo
type GameState struct {
	Name      string     
	Points    int        
	Questions []Question 
}

// Função que inicializa o jogo, pedindo o nome do jogador
func (g *GameState) Init() {
	fmt.Println("------------------------------------")
	fmt.Println("\033[1;34mSeja Bem-Vindo ao Game-Quiz de Pokémon!\033[0m") 
	fmt.Println("------------------------------------")
	fmt.Print("Escreva seu Nome: ") 

	
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n') 


	if err != nil {
		panic("Erro ao ler string!") 
	}

	g.Name = strings.TrimSpace(name)
	fmt.Printf("Vamos Comecar o Game_Quiz, %s!\n", g.Name)
}

// Função que processa o arquivo CSV com as perguntas do quiz
func (g *GameState) ProcessCSV() {

	f, err := os.Open("text.csv")
	
	
	if err != nil {
		panic("Erro ao ler arquivo!")
	}
	defer f.Close() 

	
	reader := csv.NewReader(f)
	reader.Comma = ','


	records, err := reader.ReadAll() 
	if err != nil {
		panic("Erro ao ler CSV!")
	}

	
	for index, record := range records {
		
		if index > 0 {
			
			correctAnswer, _ := Toint(record[5])

			
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}
			g.Questions = append(g.Questions, question)
		}
	}
}


func Toint(s string) (int, error) {
	
	i, err := strconv.Atoi(s)
	if err != nil {
	
		return 0, errors.New("por favor, insira um número e não caracteres")
	}


	return i, nil
}

// Função que executa o quiz, fazendo as perguntas ao jogador e verificando as respostas
func (g *GameState) Run() {
	
	for index, question := range g.Questions {	
		fmt.Printf("\033[32m %d. %s\033[m\n", index+1, question.Text)
		for j, option := range question.Options {
			fmt.Printf("\033[36m[%d] %s\033[m\n", j+1, option)
		}
		var answer int
		var err error
		for {
			fmt.Println("Digite uma Alternativa: ")
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')
			read = strings.TrimSpace(read)
			answer, err = Toint(read)

			if err != nil || answer < 1 || answer > 4 {
				fmt.Println("Por favor, insira uma alternativa válida (1 a 4).")
				continue
			}

			if answer == question.Answer {
				fmt.Println("Resposta correta!")
				g.Points += 10 
			} else {
				
				fmt.Printf("Resposta errada! A resposta correta era: %d\n", question.Answer)
				fmt.Println("------------------------------------")
			}
			break
		}
	}
}

// Função principal que inicializa e executa o jogo
func main() {
	game_01 := &GameState{Points: 0}
	game_01.ProcessCSV()
	game_01.Init()
	game_01.Run()

	fmt.Printf("\033[32mFim do Game-Quiz\nParabéns, %s! Você obteve %d pontos!\033[m\n", game_01.Name, game_01.Points)
}
