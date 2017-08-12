import React, { Component } from 'react'
import { HashRouter, Link, Route } from 'react-router-dom'
import { Box } from 'react-polymer-layout'

class Nav extends Component {
  render() {
    return <Box vertical>
      <Link to="/">home</Link>
      <Link to="/content">content</Link>
      <Link to="/about">about</Link>
      {this.props.children}
    </Box>
  }
}

class NameComponent extends Component {
  render() {
    return <div>{this.props.name}</div>
  }
}

class App extends Component {
  render() {
    return <HashRouter>
      <Nav>
        <Route exact path="/" render={() => <NameComponent name="Home" />} />
        <Route exact path="/content" render={() => <NameComponent name="Content" />} />
        <Route exact path="/about" render={() => <NameComponent name="About" />} />
      </Nav>
    </HashRouter>
  }
}

export default App;
