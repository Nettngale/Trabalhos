senha: str = ""     #Variavel da senha
num1: float = 0     #Variavel do primeiro numero
num2: float = 0     #Variavel do segundo numero
s: str = ""         #Variavel do sexo
a: float = 0        #Variavel da altura
m: float = 72.7     #Variavel com o valor para a multiplicaçao caso for Masculino
f: float = 62.1     #Variavel com o valor para a multiplicaçao caso for Feminino
M: str = "Manhã"         #Variavel para o turno da manha
T: str = "Tarde"         #Variavel para o turno da tarde
N: str = "Noite"         #Variavel para o turno da noite
turno: str = ""
nota1: float = 0         #Variavel da nota 1
nota2: float = 0         #Variavel da nota 2
nota3: float = 0         #Variavel da nota 3
nota4: float = 0         #Variavel da nota 4
notatotal: float = 0     #Variavel da soma das notas
media: float = 0         #Media das 4 notas


senha = str(input("Digite a senha: "))                                              #Entrada da senha pelo usuario
if senha == "dev24":                                                                #Inicio do teste da senha, se correto, irá dar continuidade a sequencia, proximo testes do numeros
    print ("Acesso Permitido.")
    
    turno = str(input("Você estuda em que turno? M(Manhã), T(Tarde), N(Noite):"))    #Entrada que recebe o horario que o aluno estuda 
    if turno == "M":                                                                 #M se estuda de manha  
        print(" ")                                                                   #Quebra de linha
        print("Bom dia!")                                                            #Caso estude de manha recebera esta mensagem
        print(" ")                                                                   #Quebra de linha
    if turno == "T":                                                                 #T se estuda de tarde
        print(" ")                                                                   #Quebra de linha
        print("Boa Tarde!")                                                          #Caso estude a tarde recebera esta mensagem
        print(" ")                                                                   #Quebra de linha
    if turno == "N":                                                                 #N se estuda de noite
        print(" ")                                                                   #Quebra de linha
        print("Boa Noite!")                                                          #Caso estude a noite recebera esta mensagem
        print(" ")                                                                   #Quebra de linha
    if turno != "M" and turno != "T" and turno != "N":                               #Caso o usuário digite um valor diferente de "M" para manha, "T" para tarde, "N" para noite                               
        print("Horário Inexistente")                                                 #Exibira esta mensagem caso o horario nao for compativel com os disponiveis

    nota1 = float(input("Digite a primeira nota: "))                                 #Recebe a primeira nota
    nota2 = float(input("Digite a segunda nota: "))                                  #Recebe a segunda nota
    nota3 = float(input("Digite a terceira nota: "))                                 #Recebe a terceira nota
    nota4 = float(input("Digite a quarta nota: "))                                   #Recebe a quarta nota
    notatotal = float(nota1 + nota2 + nota3 + nota4)                                 #Soma das 4 notas
    media = float((nota1 + nota2 + nota3 + nota4) / 2)                               #Media das 4 notas
    print (" ")                                                                      #Quebra de linha
    print ("Total: ",notatotal)                                                      #Saida com o valor da soma das notas
    print ("Sua média é: ",media)                                                    #Resultado da media
    
    if media >= 7:                                                                   #Teste se a media for maior ou igual a 7
        print("Aprovado! Parabéns!")                                                 #Se for maior que 7 a mensagem de aprovado ira aparecer
        print(" ")                                                                   #Quebra de linha
    else:                                                                            #Teste se a media nao for maior ou igual a 7
        print("Reprovado :(")                                                        #Aparecera a mensagem de reprovado
        print(" ")                                                                   #Quebra de linha


    num1 = float(input("Digite o primeiro Número: "))                           #Entrada do valor do primeiro numero pelo usuario
    num2 = float(input("Digite o segundo Número: "))                            #Entrada do valor do segundo numero pelo usuario
    if num1 > num2:                                                             #Teste para verificar caso o primeiro numero é > que o segundo numero
       print("O primeiro número:",num1," é maior que o segundo número:",num2)
    if num2 > num1:                                                             #Teste para verificar caso o segundo numero é > que o primeiro numero
        print("O segundo número:",num2,"é maior que o primeiro número:",num1)
    if num1 == num2:                                                            #Teste se forem iguais, exibira a mensagem que sao iguais
        print("Os números tem o mesmo valor...")
    
    
    s = str(input("Digite o sexo M (Masculino) e F (Feminino): "))              #Entrada do sexo
    if s == "M":                                                                #Teste se do sexo masculino
        a = float(input("Digite a altura: "))                                   #Entrada da altura
        print(" ")                                                              #Quebra de linha
        print("Altura:",a)                                                      #Saida da altura                                                                                                                 
        print("Sexo:Masculino")                                                 #Saida do sexo
        print("Resultado:",m * a)                                               #Saida do calculo de 72.7(variavel m) * altura                                        
    
    if s == "F":                                                                #Teste se do sexo feminino
        a = float(input("Digite a altura: "))                                   #Entrada da altura
        print(" ")                                                              #Quebra de linha
        print("Altura:",a)                                                      #Saida da altura
        print("Sexo:Feminino")                                                  #Saida do sexo
        print("Resultado:",f * a)                                               #Saida do calculo de 62.1(variavel f) * altura 
    if s != "F" and s != "M":                                                   #Teste se s for diferente da entrada "F" e "M" 
        print("Erro, reiniciar processo.")                                      #Mensagem de erro caso for diferente de "F" e "M"             
else:                                                                           #Se a senha nao estiver correta um comando de saida negando acesso e finalizando o processo será lançado
    print("Acesso Negado! Senha Invalida.")


