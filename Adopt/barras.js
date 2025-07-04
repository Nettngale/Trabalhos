const historico = JSON.parse(sessionStorage.getItem("historico"));
const ctx = document.getElementById("graficoBarra").getContext("2d");

const labels = Object.keys(historico);
const entradas = labels.map(m => historico[m].entrada);
const saidas = labels.map(m => historico[m].saida);

new Chart(ctx, {
  type: "bar",
  data: {
    labels,
    datasets: [
      { label: "Entradas", data: entradas, backgroundColor: "#2ecc71" },
      { label: "Saídas", data: saidas, backgroundColor: "#e74c3c" }
    ]
  }
});
