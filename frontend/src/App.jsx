import { useState } from 'react'

function App() {
  const [clienteID, setClienteID] = useState('')
  const [barbeiroID, setBarbeiroID] = useState('')
  const [data, setData] = useState('')
  const [horarios, setHorarios] = useState([])
  const [buscou, setBuscou] = useState(false)

  function buscarHorarios() {
    fetch(`http://localhost:8080/barbeiros/${barbeiroID}/horarios?data=${data}`)
      .then((r) => r.json())
      .then((dados) => { setHorarios(dados); setBuscou(true) })
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
    }).then((r) => {
      if (r.ok) buscarHorarios()
      else alert('Não foi possível agendar esse horário.')
    })
  }

  return (
    <div className="page">
      <div className="card">
        <header className="header">
          <span className="logo">💈</span>
          <h1>Barbearia do Matheus</h1>
          <p className="subtitle">Agende seu horário</p>
        </header>

        <div className="form">
          <label>Cliente
            <input type="number" placeholder="ID do cliente"
              value={clienteID} onChange={(e) => setClienteID(e.target.value)} />
          </label>
          <label>Barbeiro
            <input type="number" placeholder="ID do barbeiro"
              value={barbeiroID} onChange={(e) => setBarbeiroID(e.target.value)} />
          </label>
          <label>Data
            <input type="date"
              value={data} onChange={(e) => setData(e.target.value)} />
          </label>
          <button className="btn-primary" onClick={buscarHorarios}>Ver horários</button>
        </div>

        {buscou && (
          <div className="horarios">
            <h3>Horários disponíveis</h3>
            {horarios.length === 0 ? (
              <p className="vazio">Nenhum horário livre nesse dia. 😕</p>
            ) : (
              <div className="slots">
                {horarios.map((hora) => (
                  <button key={hora} className="slot" onClick={() => agendar(hora)}>
                    {hora}
                  </button>
                ))}
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  )
}

export default App