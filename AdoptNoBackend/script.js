const TOTAL_CAPACIDADE = 1000;
const limites = {
  equino: 0.10,
  felino: 0.25,
  canino: 0.25,
  ave: 0.20,
  aquatico: 0.20
};

function processar() {
  const tipo = document.getElementById("tipo").value;
  const quantidade = parseInt(document.getElementById("quantidade").value);
  const data = document.getElementById("data").value;
  const operacao = document.getElementById("operacao").value;

  if (!data || isNaN(quantidade) || quantidade <= 0) {
    alert("Por favor, preencha todos os campos corretamente.");
    return;
  }

  const registros = JSON.parse(localStorage.getItem("registros") || "[]");

  // Calcular a ocupação atual
  const totalPorTipo = {
    equino: 0,
    felino: 0,
    canino: 0,
    ave: 0,
    aquatico: 0
  };

  let totalAtual = 0;

  registros.forEach(({ tipo, quantidade, operacao }) => {
    if (operacao === "entrada") {
      totalPorTipo[tipo] += quantidade;
      totalAtual += quantidade;
    } else {
      totalPorTipo[tipo] -= quantidade;
      totalAtual -= quantidade;
    }
  });

  if (operacao === "entrada") {
    const novoTotalTipo = totalPorTipo[tipo] + quantidade;
    const novoTotalGeral = totalAtual + quantidade;
    const limiteTipo = TOTAL_CAPACIDADE * limites[tipo];

    if (novoTotalTipo > limiteTipo) {
      alert(`O limite de ${tipo}s foi atingido, uma mensagem automática já foi enviada para outros abrigos verificando se há vagas.`);
    } else if (novoTotalGeral > TOTAL_CAPACIDADE) {
      alert("A capacidade total do abrigo foi ultrapassada.");
    } else {
      registros.push({ tipo, quantidade, data, operacao });
      localStorage.setItem("registros", JSON.stringify(registros));
      location.href = "grafico-pizza.html";
    }
  } else {
    registros.push({ tipo, quantidade, data, operacao });
    localStorage.setItem("registros", JSON.stringify(registros));
    location.href = "grafico-pizza.html";
  }
}
