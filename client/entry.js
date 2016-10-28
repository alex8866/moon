import React from 'react'
import ReactDOM from 'react-dom'
import { Router, Route, browserHistory } from 'react-router'
/* Import components for use with router */
import App from './components/app.js'
import About from './components/about.js'
import Contact from './components/contact.js'
/* Import stylesheet for main app */
import './style.scss'

const routes = {
	path: '/',
	component: App,
	childRoutes: [
		{ path: 'about', component: About },
		{ path: 'contact', component: Contact },
	]
}

ReactDOM.render((
	<Router
		routes={routes}
		history={browserHistory}
	/>),
	document.getElementById('app'))

if(module.hot) {
	module.hot.accept();
}
