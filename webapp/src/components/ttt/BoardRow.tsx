// import { useState } from 'react'
// import reactLogo from '../assets/react.svg'
// import viteLogo from '/vite.svg'
// import '../styles/App.css'

import { useState } from 'react'
import "./styles/BoardRow.css"

function BoardRow(cells: number[]) {

  const playerMapping = new Map<number, string>([
    [1, "X"],
    [2, "O"],
    [3, "\u9650"],
    [4, "\u9632"]
  ])

  const [rowData, setRowData] = useState(cells)

  return (
    <>
      {
      rowData.map((num) => {
        return (
          <div className='cell' >
            {playerMapping.get(num)}
          </div>
        )
      })
      }
    </>
  )
}

export default BoardRow
