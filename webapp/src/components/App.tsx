import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './styles/App.css'

import BoardRow from './ttt/BoardRow'
import {BoardRow as ProtoBoardRow} from '@proto/games/tictactoe/board_pb'

function App() {
  const [count, setCount] = useState(0)

  const boardRow = ProtoBoardRow.fromJson({cells:[1,2,3]})

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>

      <BoardRow {...boardRow.cells}/>
    </>
  )
}

export default App
