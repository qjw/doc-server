import { loginApi } from '../services/index'
import {routerRedux} from 'dva/router'
import * as qs from 'query-string';
import URL from 'url-parse'

export default {

  namespace: 'login',

  state: {
  },

  subscriptions: {
    setup ({ dispatch, history }) {
    },
  },

  effects: {
    * login ({payload}, { call, put }) {
      	const {data} = yield call(loginApi)
      	yield put({
        		type: 'querySuccess',
        		data: data,
      	})
	
		let url = payload
		if(!url){
		    yield put(routerRedux.push({
	          pathname: url,
	        }))
		} else {
	  		const query = new URL(url);
		    yield put(routerRedux.push({
	          pathname: query.pathname,
			  search: query.query,
	        }))
		}
    },
  },

  reducers: {
    querySuccess (state, { data }) {
      return {
        ...state,
        ...data,
      }
    },
  },
}