import React, { useState, useEffect } from 'react'
import { render } from 'react-dom'
import axios from 'axios'
import moment from 'moment'

const App = () => {
  const [weight, setWeight] = useState(0)
  const [yesterdayWeight, setYesterdayWeight] = useState(0)
  const [weights, setWeights] = useState([])
  const [value, setValue] = useState('')
  const date = new Date()

  useEffect(() => {
    const fetchData = async () => {
      const result = await axios.get(`http://localhost:8081/`)
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
    axios.post(`http://localhost:8081/`, params).then(() => {
      setWeight(newWeight)
      setValue('')
    })
  }

  const handleGetWeights = (event) => {
    event.preventDefault()

    axios.get(`http://localhost:8081/weights/all/`).then((result) => {
      setWeights(result.data)
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
      <a onClick={handleGetWeights}>記録</a>
      <List weights={weights} />
    </React.Fragment>
  )
}

const List = (props) => {
  if (props.weights.length == 0) return null

  return (
    <table width="200" border="1" style={{ borderCollapse: 'collapse' }}>
      <tbody>
        <tr>
          <th>日付</th>
          <th>体重</th>
        </tr>
        {props.weights.map((weight, index) => {
          return (
            <tr key={index}>
              <td align="center">{moment(weight.date).format('YYYY-MM-DD')}</td>
              <td align="center">{weight.weight} g</td>
            </tr>
          )
        })}
      </tbody>
    </table>
  )
}

render(<App />, document.getElementById('app'))
