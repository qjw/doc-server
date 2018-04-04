import { descriptionApi,currentApi } from '../services/index'
import {routerRedux} from 'dva/router'
import queryString from 'query-string'
import axios from 'axios'

export default {

  namespace: 'app',

  state: {
	description: null,
	pathname: "",
	current: null
  },

  subscriptions: {	
    setup ({ dispatch, history }) {		
      axios.interceptors.response.use(function (response) {
        // Do something with response data
        return response;
      }, function (error) {
        const {response} = error
        let msg;
        let statusCode;
        if (response && response instanceof Object) {
          statusCode = response.status
          if (statusCode === 401) {
            // Do something with response error
            dispatch({
              type: 'unauthorized',
              payload: error,
            })
          }
        }

        return Promise.reject(error);
      })
	
		history.listen(({ pathname }) => {
			dispatch({
			    type: 'querySuccess',
			    data: {
					pathname: pathname,
				},
			})
			dispatch({ type: 'query' })
			if(pathname !== "/login" && pathname !== "/login_cb"){
				dispatch({ type: 'queryCurrent' })
			} else {
				dispatch({
				    type: 'querySuccess',
				    data: {
						current: null,
					},
				})
			}
	    })
    },
  },

  effects: {
    * query ({}, { call, put, select }) {
      const { app } = yield select(_ => _)
	  if(!app.description) {
		  const {data} = yield call(descriptionApi)
	      yield put({
	        type: 'querySuccess',
	        data: {
				description: data,
			},
	      })
	  }
    },
	
    * queryCurrent ({}, { call, put, select }) {
      const { app } = yield select(_ => _)
	  if(!app.current) {
		  const {data} = yield call(currentApi)
	      yield put({
	        type: 'querySuccess',
	        data: {
				current: data,
			},
	      })
	  }
    },
	
	* unauthorized(action, {put, select}) {
      const {pathname} = yield select(_ => _.app)
      if (pathname != "" && pathname !== "/login") {
        yield put(routerRedux.push({
          pathname: "/login",
          search: "?" + queryString.stringify({
            url: window.location.href
          }),
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