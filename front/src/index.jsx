import React, { useState, useEffect } from 'react'
import {render} from 'react-dom'
import axios from 'axios'

function App() {
  const [weight, setWeight] = useState(0)
  const [yesterdayWeight, setYesterdayWeight] = useState(0)

  useEffect(() => {
    const fetchData = async ()  => {
      const result = await axios.get(`http://localhost:8080/`)
      setWeight(result.data.weight.weight)
    }

    fetchData()
  }, [])

  return (
    <React.Fragment>
      <form action="/weights/" method="post">
        <p>2021-01-01</p>
        <input type="text" name="weight" />
        <input type="submit" value="保存" />
        <p>今日の体重: { weight } g</p>
        <p>昨日の体重: { yesterdayWeight } g</p>
      </form>
      <a href="/weights/all/">記録</a>
    </React.Fragment>
  )
}

render(<App/>, document.getElementById('app'))
