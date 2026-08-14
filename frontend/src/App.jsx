import { useState } from 'react'

function App() {
  const [clienteID, setClienteID] = useState('')
  const [barbeiroID, setBarbeiroID] = useState('')
  const [data, setData] = useState('')
  const [horarios, setHorarios] = useState([])

  function buscarHorarios() {
    fetch(`http://localhost:8080/barbeiros/${barbeiroID}/horarios?data=${data}`)
      .then((resposta) => resposta.json())
      .then((dados) => setHorarios(dados))
  }

  function agendar(hora) {
    const dataHora = `${data}T${hora}:00Z`
    fetch('http://localhost:8080/agendamentos', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        cliente_id: Number(clienteID),
        barbeiro_id: Number(barbeiroID),
        data_hora: dataHora,
      }),
    }).then((resposta) => {
      if (resposta.ok) {
        buscarHorarios()
      } else {
        alert('Não foi possível agendar esse horário.')
      }
    })
  }

  return (
    <div>
      <h1>Barbearia Dourado</h1>
      <h2>Agendar um horário</h2>

      <input type="number" placeholder="ID do cliente"
        value={clienteID} onChange={(e) => setClienteID(e.target.value)} />
      <input type="number" placeholder="ID do barbeiro"
        value={barbeiroID} onChange={(e) => setBarbeiroID(e.target.value)} />
      <input type="date"
        value={data} onChange={(e) => setData(e.target.value)} />
      <button onClick={buscarHorarios}>Ver horários</button>

      <h3>Horários disponíveis (clique para agendar):</h3>
      <div>
        {horarios.map((hora) => (
          <button key={hora} onClick={() => agendar(hora)}>
            {hora}
          </button>
        ))}
      </div>
    </div>
  )
}

export default App