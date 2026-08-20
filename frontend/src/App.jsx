import { useState } from 'react'

const API = import.meta.env.VITE_API_URL || 'http://localhost:8080'

function App() {
  const [clienteID, setClienteID] = useState('')
  const [barbeiroID, setBarbeiroID] = useState('')
  const [data, setData] = useState('')
  const [horarios, setHorarios] = useState([])
  const [buscou, setBuscou] = useState(false)
  const [toast, setToast] = useState(null)
  const [agendamentos, setAgendamentos] = useState([])

  function mostrarToast(texto, tipo) {
    setToast({ texto, tipo })
    setTimeout(() => setToast(null), 3000)
  }

  function buscarAgendamentos() {
    fetch(`${API}/agendamentos?cliente_id=${clienteID}`)
      .then((r) => r.json())
      .then((dados) => setAgendamentos(dados))
  }

  function buscarHorarios() {
    fetch(`${API}/barbeiros/${barbeiroID}/horarios?data=${data}`)
      .then((r) => r.json())
      .then((dados) => { setHorarios(dados); setBuscou(true) })
    buscarAgendamentos()
  }

  function agendar(hora) {
    const dataHora = `${data}T${hora}:00Z`
    fetch(`${API}/agendamentos`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        cliente_id: Number(clienteID),
        barbeiro_id: Number(barbeiroID),
        data_hora: dataHora,
      }),
    }).then((r) => {
      if (r.ok) {
        mostrarToast('Horário agendado com sucesso! ✅', 'sucesso')
        buscarHorarios()
        buscarAgendamentos()
      } else {
        mostrarToast('Não foi possível agendar esse horário. 😕', 'erro')
      }
    })
  }

  function cancelar(id) {
    fetch(`${API}/agendamentos/${id}/cancelar`, { method: 'PATCH' })
      .then((r) => {
        if (r.ok) {
          mostrarToast('Agendamento cancelado.', 'sucesso')
          buscarAgendamentos()
          buscarHorarios()
        } else {
          mostrarToast('Não foi possível cancelar.', 'erro')
        }
      })
  }

  return (
    <div className="page">
      {toast && <div className={`toast toast-${toast.tipo}`}>{toast.texto}</div>}

      <div className="card">
        <header className="header">
          <span className="logo">💈</span>
          <h1>Barbearia Dourado</h1>
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

        <div className="lista">
          <h3>Agendamentos</h3>
          {agendamentos.length === 0 ? (
            <p className="vazio">Nenhum agendamento ainda.</p>
          ) : (
            agendamentos.map((ag) => (
              <div key={ag.id} className="agendamento">
                <span>
                  #{ag.id} · barbeiro {ag.barbeiro_id} · {ag.data_hora.slice(0, 16).replace('T', ' ')}
                  <span className={`tag tag-${ag.status}`}>{ag.status}</span>
                </span>
                {ag.status === 'agendado' && (
                  <button className="btn-cancelar" onClick={() => cancelar(ag.id)}>Cancelar</button>
                )}
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  )
}

export default App