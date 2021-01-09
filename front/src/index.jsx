import React, { useState } from 'react';
import {render} from 'react-dom';

function App() {
  const [weight, setWeight] = useState(0)
  const [yesterdayWeight, setYesterdayWeight] = useState(0)

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

render(<App/>, document.getElementById('app'));
