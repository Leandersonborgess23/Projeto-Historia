package main

import (
	"fmt"
	"os"
)

func main() {
	// Cria o arquivo HTML
	file, err := os.Create("formacao_territorial_brasil.html")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	// Define o conteúdo HTML dos slides
	content := `
<!DOCTYPE html>
<html lang="pt-BR">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Formação Territorial do Brasil</title>
    <style>
        body { font-family: 'Arial', sans-serif; margin: 0; padding: 0; }
        .container { width: 80%; margin: auto; }
        .slide { display: none; padding: 20px; border: 1px solid #ddd; margin-top: 20px; border-radius: 10px; }
        .active { display: block; }
        .navigation { text-align: center; margin-top: 20px; }
        .navigation button { padding: 10px 20px; margin: 0 10px; border: none; border-radius: 5px; background-color: #007bff; color: white; font-size: 16px; cursor: pointer; }
        .navigation button:hover { background-color: #0056b3; }
        h1 { text-align: center; color: #333; }
        h2 { color: #007bff; }
        p { line-height: 1.6; }
        a { color: #007bff; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
    <script>
        let currentSlide = 0;
        function showSlide(index) {
            const slides = document.querySelectorAll('.slide');
            slides.forEach((slide, i) => {
                slide.classList.toggle('active', i === index);
            });
            currentSlide = index;
        }
        function nextSlide() {
            const slides = document.querySelectorAll('.slide');
            if (currentSlide < slides.length - 1) {
                showSlide(currentSlide + 1);
            }
        }
        function prevSlide() {
            if (currentSlide > 0) {
                showSlide(currentSlide - 1);
            }
        }
        document.addEventListener('DOMContentLoaded', () => {
            showSlide(currentSlide);
        });
    </script>
</head>
<body>
    <div class="container">
        <h1>Formação Territorial do Brasil</h1>
        
        <div class="slide active">
            <h2>Introdução</h2>
            <p>A formação territorial do Brasil é um processo complexo que abrange desde a época colonial até a expansão territorial no século XX. Vamos explorar os principais eventos e tratados que moldaram as fronteiras do Brasil.</p>
        </div>
        
        <div class="slide">
            <h2>Período Colonial</h2>
            <p>No período colonial, o território brasileiro foi delimitado por tratados com potências europeias e pela expansão interna. O Tratado de Tordesilhas (1494) estabeleceu uma linha de divisão entre as terras portuguesas e espanholas.</p>
            <p><strong>Fonte:</strong> <a href="https://www.britannica.com/event/Treaty-of-Tordesillas" target="_blank">Encyclopaedia Britannica</a></p>
        </div>
        
        <div class="slide">
            <h2>Independência e Consolidação</h2>
            <p>Após a independência em 1822, o Brasil consolidou seu território através de tratados e conflitos. A Guerra de Cisplatina (1825-1828) resultou na perda da província de Cisplatina para o Uruguai, e os Tratados de Limitação com os países vizinhos ajudaram a definir as fronteiras.</p>
            <p><strong>Fonte:</strong> <a href="https://www.historiabrasileira.com.br/1822-independencia/" target="_blank">História Brasileira</a></p>
        </div>
        
        <div class="slide">
            <h2>Expansão para o Oeste</h2>
            <p>No final do século XIX e início do século XX, o Brasil expandiu seu território para o oeste com a ocupação e colonização da região amazônica. O Tratado de Petrópolis (1903) com a Bolívia foi fundamental para a incorporação do Acre.</p>
            <p><strong>Fonte:</strong> <a href="https://www.culturabrasil.org.br/tratado-de-petropolis/" target="_blank">Cultura Brasil</a></p>
        </div>
        
        <div class="navigation">
            <button onclick="prevSlide()">Anterior</button>
            <button onclick="nextSlide()">Próximo</button>
        </div>
    </div>
</body>
</html>
    `

	// Escreve o conteúdo no arquivo
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Arquivo HTML gerado com sucesso: formacao_territorial_brasil.html")
}
