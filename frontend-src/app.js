import { render } from 'react-dom'
import React from 'react'
import h1 from './styles/main.scss'

class App extends React.Component {
  render() {
    return (
      <h1> was this worth it???? </h1>
    )
  }
}

render(<App />, document.body)
