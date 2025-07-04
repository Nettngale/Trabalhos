const porcentagens = JSON.parse(sessionStorage.getItem("porcentagens"));

const ctx = document.getElementById("graficoPizza").getContext("2d");
const labels = Object.keys(porcentagens);
const valores = Object.values(porcentagens);

new Chart(ctx, {
  type: "pie",
  data: {
    labels,
    datasets: [{
      data: valores,
      backgroundColor: ['#f39c12', '#2ecc71', '#3498db', '#9b59b6', '#e74c3c']
    }]
  }
});
