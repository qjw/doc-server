import pathToRegexp from 'path-to-regexp'
import queryString from 'query-string'
import {routerRedux} from 'dva/router'
import { loginQyApi } from '../services/index'
import URL from 'url-parse'

export default {
  namespace: 'loginCB',

  state: {
    data: {},
  },

  subscriptions: {
    setup({dispatch, history}) {
      history.listen((location) => {
		if(location.pathname == "/login_cb") {
		  console.log("fuckshit",location.search)
          location.query = queryString.parse(location.search)
		  console.log("fuckshit",location.query)
          dispatch({type: 'login', payload: location.query})
        }
      })
    },
  },

  effects: {
    * login({
              payload,
            }, {call, put}) {
      const {data} = yield call(loginQyApi, payload)
	  console.log("fuckshit",payload)
      if (payload["url"]) {
		let query = new URL(payload["url"]);
		query = query.hash
		if(query.length > 0) {
			query = query.substr(1)
			query = new URL(query);
		    yield put(routerRedux.push({
	          pathname: query.pathname,
			  search: query.query,
	        }))
		} else {
	        yield put(routerRedux.push({
	          pathname: "/",
	        }))
		}
      }else {
        yield put(routerRedux.push({
          pathname: "/",
        }))
      }
    },
  },
}