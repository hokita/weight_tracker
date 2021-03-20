import React, { useState, useEffect } from 'react'
import axios from 'axios'
import moment from 'moment'

const apiURL = `http://${process.env.API_DOMAIN}:8081/`

const App = () => {
  const [weight, setWeight] = useState(0)
  const [yesterdayWeight, setYesterdayWeight] = useState(0)
  const [weights, setWeights] = useState([])
  const [value, setValue] = useState('')
  const [listToggle, setListToggle] = useState(false)

  useEffect(() => {
    const fetchData = async () => {
      const result = await axios.get(apiURL)
      setWeight(result.data.weight.weight)
      setYesterdayWeight(result.data.yesterday_weight.weight)
    }

    fetchData()
  }, [])

  const handleChange = (event: any) => {
    switch (event.target.name) {
      case 'weight':
        setValue(event.target.value)
        break
    }
  }

  const handleSubmit = (event: any) => {
    event.preventDefault()

    const newWeight = parseInt(value)
    const params = JSON.stringify({ weight: newWeight })
    axios.post(apiURL, params).then(() => {
      setWeight(newWeight)
      setValue('')
    })
  }

  const handleGetWeights = (event: any) => {
    event.preventDefault()

    if (weights.length === 0) {
      axios.get(apiURL + 'weights/all/').then((result: any) => {
        setWeights(result.data)
      })
    }

    setListToggle(!listToggle)
  }

  const calcDifference = () => {
    if (weight === 0) return 0
    return weight - yesterdayWeight
  }

  return (
    <React.Fragment>
      <h1>わさ体重記録</h1>
      <p>{moment().format('YYYY-MM-DD')}</p>
      <form onSubmit={handleSubmit}>
        <input
          type="number"
          name="weight"
          value={value}
          onChange={handleChange}
        />
        <input type="submit" value="保存" />
        <p>
          今日の体重: {weight} g ({calcDifference()} g)
        </p>
        <p>昨日の体重: {yesterdayWeight} g</p>
      </form>
      <button onClick={handleGetWeights}>履歴</button>
      <List weights={weights} display={listToggle} />
    </React.Fragment>
  )
}

const List = ({ weights, display }: { weights: any; display: any }) => {
  if (!display) return null

  return (
    <table width="200" style={{ borderCollapse: 'collapse' }}>
      <tbody>
        <tr>
          <th>日付</th>
          <th>体重</th>
        </tr>
        {weights.map((weight: any, index: any) => {
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

export default App
