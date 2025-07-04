const registros = JSON.parse(localStorage.getItem("registros") || "[]");

const meses = ["01","02","03","04","05","06","07","08","09","10","11","12"];
const entradas = Array(12).fill(0);
const saidas = Array(12).fill(0);

registros.forEach(({ quantidade, data, operacao }) => {
  const mes = new Date(data).getMonth();
  if (operacao === "entrada") entradas[mes] += quantidade;
  else saidas[mes] += quantidade;
});

new Chart(document.getElementById("graficoBarra"), {
  type: "bar",
  data: {
    labels: meses.map(m => `Mês ${m}`),
    datasets: [
      {
        label: "Entradas",
        backgroundColor: "#2ecc71",
        data: entradas
      },
      {
        label: "Saídas",
        backgroundColor: "#e74c3c",
        data: saidas
      }
    ]
  },
  options: {
    responsive: true,
    plugins: { legend: { position: 'top' } }
  }
});
