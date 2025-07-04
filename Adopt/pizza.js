const registros = JSON.parse(localStorage.getItem("registros") || "[]");

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

const labels = Object.keys(totalPorTipo);
const data = labels.map(t => totalPorTipo[t]);
const backgroundColors = ['#ff6384', '#36a2eb', '#ffce56', '#8e44ad', '#2ecc71'];

new Chart(document.getElementById("graficoPizza"), {
  type: "pie",
  data: {
    labels,
    datasets: [{
      data,
      backgroundColor: backgroundColors
    }]
  }
});
