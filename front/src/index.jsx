import React, { useState, useEffect } from 'react'
import { render } from 'react-dom'
import axios from 'axios'
import moment from 'moment'

function App() {
  const [weight, setWeight] = useState(0)
  const [yesterdayWeight, setYesterdayWeight] = useState(0)
  const [value, setValue] = useState('')
  const date = new Date()

  useEffect(() => {
    const fetchData = async () => {
      const result = await axios.get(`http://localhost:8080/`)
      setWeight(result.data.weight.weight)
      setYesterdayWeight(result.data.yesterday_weight.weight)
    }

    fetchData()
  }, [])

  const handleChange = (event) => {
    switch (event.target.name) {
      case 'weight':
        setValue(event.target.value)
        break
    }
  }

  const handleSubmit = (event) => {
    event.preventDefault()

    const newWeight = parseInt(value)
    const params = JSON.stringify({ weight: newWeight })
    axios.post(`http://localhost:8080/`, params).then(() => {
      setWeight(newWeight)
      setValue('')
    })
  }

  return (
    <React.Fragment>
      <h1>わさ体重記録</h1>
      <p>{moment().format('YYYY-MM-DD')}</p>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="weight"
          value={value}
          onChange={handleChange}
        />
        <input type="submit" value="保存" />
        <p>今日の体重: {weight} g</p>
        <p>昨日の体重: {yesterdayWeight} g</p>
      </form>
      <a href="/weights/all/">記録</a>
    </React.Fragment>
  )
}

render(<App />, document.getElementById('app'))
