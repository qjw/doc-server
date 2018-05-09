import { descriptionApi } from '../services/index'
import {routerRedux} from 'dva/router'
import queryString from 'query-string'
import axios from 'axios'

export default {

  namespace: 'app',

  state: {
	description: null,
	pathname: "",
  },

  subscriptions: {	
    setup ({ dispatch, history }) {			
		history.listen(({ pathname }) => {
			dispatch({
			    type: 'querySuccess',
			    data: {
					pathname: pathname,
				},
			})
			dispatch({ type: 'query' })
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