async function processar() {
  const tipo = document.getElementById("tipo").value;
  const quantidade = document.getElementById("quantidade").value;
  const data = document.getElementById("data").value;
  const operacao = document.getElementById("operacao").value;

  const res = await fetch("http://127.0.0.1:5000/atualizar", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ tipo, quantidade, data, operacao })
  });
  const dados = await res.json();

  sessionStorage.setItem("porcentagens", JSON.stringify(dados.porcentagens));
  sessionStorage.setItem("historico", JSON.stringify(dados.historico));

  if (!dados.tem_espaco || dados.excedeu_limite) {
    alert("Excedeu o limite! Já entramos em contato com outros abrigos.");
  } else {
    location.href = "grafico-pizza.html";
  }
}
