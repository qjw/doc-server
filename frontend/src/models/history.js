import { historyApi } from '../services/index'
import queryString from 'query-string'

export default {

  namespace: 'history',

  state: {
	file: "",
	name: "",
	production: ""
  },

  subscriptions: {
    setup ({ dispatch, history }) {
      history.listen((location) => {
        if (location.pathname == "/history") {
          location.query = queryString.parse(location.search);
          dispatch({ type: 'query', payload: location.query})
        }
      })
    },
  },

  effects: {
    * query ({payload}, { call, put }) {
      const {data} = yield call(historyApi, payload)
      yield put({
        type: 'querySuccess',
        data: {
			data,
			file: payload["file"],
			name: payload["name"],
			production: payload["production"]
		},
      })
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