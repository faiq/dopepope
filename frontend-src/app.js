import RD from 'react-dom'
import react from 'react'

class App extends react.Component {
  render() {
    return (
      <h1> Hello, bundlers! </h1>
    )
  }
}

RD.render(<App />, document.body)
