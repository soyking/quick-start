import React from 'react'
import ReactDOM from 'react-dom'
import { hashHistory, Router, Route, IndexRedirect } from 'react-router'
import App from './components/App'
import Content from './components/Content'
import './css/my.css'

ReactDOM.render((
    <Router history={hashHistory}>
        <Route path="/" component={App}>
            <IndexRedirect to="content" />
            <Route path="content" component={Content} />
        </Route>
    </Router>
), document.querySelector(".root"))