# Vetor responsavel por guardar os candidatos a prefeitura de Gotham.
candidatos:str = ["coringa","mulhergato","heravenenosa","charada"]  
# Vetor responsavel por guardar o numero/quantidade de eleitores em Gotham.
eleitores:str = ['1','2','3','4','5','6','7','8','9','10','11','12','13','14','15','16','17','18','19','20']

# Variaveis que irao receber informaçoes do usuario: 
# Variavel 'titulo' para receber um valor especifico que esta guardado em eleitores.
titulo:str = ""
# Variavel 'voto' para receber o codigo do candidato, P = Coringa, V = Hera Venenosa, G = Mulher Gato, C = Charada
voto:str = ""
# Variaveis que irao contabilizar o numero de votos que cada candidato recebe e guardar esse valor:
p:int = 0
v:int = 0
g:int = 0
c:int = 0

# Laço de repetiçao for para verificar a variavel 'i' esta dentro dos eleitores e pedir ao usuario o seu titulo de eleitor de 1 a 20.
for i in eleitores:
    titulo:str = str (input("Indique seu titulo de eleitor (1-20): "))
    print("Eleitor",titulo ,"\n")
# if para verificar se a informaçao dada pelo usuario esta dentro de eleitores, se sim perguntar que candidato esta recebendo o voto e guardar esse voto na variavel 'voto'. 
    if titulo in eleitores:
        
        voto:str = input("Quem será seu candidato para acabar com o CRIME em Gotham? Coringa(P), Hera Venenosa(V), Mulher Gato(G), Charada(C): ")
# Uma sequencia de 'se' para cada candidato (p, v, g ou c), caso o usuario selecione p por exemplo, sera contabilizado um voto para o 
# Coringa, e a variavel 'p', ira receber o valor dela + o voto e contabilizara a favor do Coringa.
        if voto == "p":
            print("Eleitor",titulo,"votou no Coringa.\n")
            
            p += 1
        if voto == "v":
            print("Eleitor",titulo,"votou na Hera Venenosa.\n")
            
            v += 1
        if voto == "g":
            print("Eleitor",titulo,"votou na Mulher Gato.\n")
            
            g += 1
        if voto == "c":
            print("Eleitor",titulo,"votou no Charada.\n")
            
            c += 1
# Sequencia de print retornando o valor nas variaveis p, v, g e c indicando quantos votos cada candidato recebeu.
print("Coringa tem",p,"votos.")
print("Hera Venenosa tem",v,"votos.")
print("Mulher Gato tem",g,"votos.")
print("Charada tem",c,"votos.")


    
        