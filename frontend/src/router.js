import React from 'react';
import { Router, Route, IndexRoute } from 'dva/router';
import IndexPage from './routes/IndexPage';
import HistoryPage from './routes/HistoryPage';
import LoginPage from './routes/LoginPage';
import LoginCbPage from './routes/LoginCbPage'
import App from './routes/App'
import Index from './models/index';
import History from './models/history';
import Login from './models/login';
import LoginCb from './models/login_cb';
import * as modelApp from './models/app';

function RouterConfig({ history }) {
  return (
    <Router history={history}>
	  <Route path="/" component={App} model={modelApp.App}>
    		<IndexRoute component={IndexPage} model={Index}/>
    		<Route path="/history" component={HistoryPage} model={History}/>
    		<Route path="/login" component={LoginPage} model={Login}/>
    		<Route path="/login_cb" component={LoginCbPage} model={LoginCb}/>
	  </Route>
    </Router>
  );
}

export default RouterConfig;
