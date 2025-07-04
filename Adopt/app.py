from flask import Flask, request, jsonify
from flask_cors import CORS

app = Flask(__name__)
CORS(app)

TOTAL_VAGAS = 1000
limites = {
    "equino": 0.10,
    "felino": 0.25,
    "canino": 0.25,
    "ave": 0.20,
    "aquatico": 0.20
}

estado = {
    "equino": 0,
    "felino": 0,
    "canino": 0,
    "ave": 0,
    "aquatico": 0
}

historico = {}

@app.route("/atualizar", methods=["POST"])
def atualizar():
    data = request.get_json()
    tipo = data["tipo"]
    quantidade = int(data["quantidade"])
    data_str = data["data"]
    operacao = data["operacao"]

    mes = data_str[:7]

    if mes not in historico:
        historico[mes] = {"entrada": 0, "saida": 0}

    if operacao == "entrada":
        estado[tipo] += quantidade
        historico[mes]["entrada"] += quantidade
    else:
        estado[tipo] = max(0, estado[tipo] - quantidade)
        historico[mes]["saida"] += quantidade

    total_ocupado = sum(estado.values())
    porcentagem_total = total_ocupado / TOTAL_VAGAS * 100
    porcentagens = {k: (v / TOTAL_VAGAS * 100) for k, v in estado.items()}

    excedeu_limite = estado[tipo] > (TOTAL_VAGAS * limites[tipo])
    tem_espaco = total_ocupado <= TOTAL_VAGAS

    return jsonify({
        "excedeu_limite": excedeu_limite,
        "tem_espaco": tem_espaco,
        "porcentagem_total": porcentagem_total,
        "porcentagens": porcentagens,
        "estado": estado,
        "historico": historico
    })

if __name__ == "__main__":
    app.run(debug=True)
