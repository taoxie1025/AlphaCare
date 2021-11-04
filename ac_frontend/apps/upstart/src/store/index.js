import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import { default as API_ENDPOINTS} from '../api'
import createPersistedState from 'vuex-persistedstate'
import Cookies from 'js-cookie'
import {UserRegistryClient} from "@/client/user_registry_grpc_web_pb";

Vue.use(Vuex)

let userRegistryClient = new UserRegistryClient(process.env.VUE_APP_NYX_BACKEND_ENDPOINT+process.env.VUE_APP_NYX_BACKEND_PORT, null, null);

export default new Vuex.Store({
    plugins: [createPersistedState({
        storage: window.sessionStorage,
    })],
    state: {
        status: '',
        token: Cookies.get('token') || '',
        claims:{},
        user:{}
  },
  mutations: {
      auth_request(state){
          state.status = 'loading'
      },
      auth_success(state, payload){
          state.status = 'success'
          state.token = payload.token
          state.claims = payload.claims
      },
      auth_error(state){
          state.status = 'error'
      },
      logout(state){
          state.status = ''
          state.token = ''
          state.user = {}
          state.claims = {}
          sessionStorage.clear()
          localStorage.clear()
      },
      read_user_success(state, payload){
          state.user = payload
      }
  },
  actions: {
    login({commit}, claims){
      return new Promise((resolve, reject) => {
        commit('auth_request')
        userRegistryClient.loginUser(claims, {}, (err, response) => {
            if (!err) {
                let resp = response.toObject()
                const token = resp.jwttoken
                const claims = resp.claims
                localStorage.setItem('token', token)
                Cookies.set('token', token, 'value', {expires: 2}) // 2 days
                window.localStorage.setItem('logged_in', true) // for triggering storage change event
                Cookies.set('claims', JSON.stringify(claims), 'value', {expires: 2})
                axios.defaults.headers.common['Authorization'] = token
                commit('auth_success', {token, claims})
                resolve(resp)
            } else {
                commit('auth_error')
                localStorage.removeItem('token')
                reject(err)
            }
        })
    })},
    register({commit}, claims){
      return new Promise((resolve, reject) => {
          commit('auth_request')
          userRegistryClient.createUser(claims, {}, (err, response) => {
              if (!err) {
                  let resp = response.toObject()
                  const token = resp.jwttoken
                  const claims = resp.claims
                  localStorage.setItem('token', token)
                  Cookies.set('token', token, 'value', {expires: 2}) // 2 days
                  Cookies.set('claims', JSON.stringify(claims), 'value', {expires: 2})
                  window.localStorage.setItem('logged_in', true)
                  axios.defaults.headers.common['Authorization'] = token
                  commit('auth_success', token, claims)
                  resolve(resp)
              } else {
                  commit('auth_error', err)
                  localStorage.removeItem('token')
                  reject(err)
              }
        })
    })},
    logout({commit}){
      return new Promise((resolve) => {
          commit('logout')
          localStorage.removeItem('token')
          window.localStorage.setItem('logged_in', false)
          Cookies.remove('token')
          Cookies.remove('claims')
          commit('logout')
          delete axios.defaults.headers.common['Authorization']
          resolve()
      })
    },
    readUser({commit}, email){
      const data = {requesterEmail: email}
      return new Promise((resolve, reject) => {
          axios({url: API_ENDPOINTS.READ_USER(email), params: data, method: 'GET' })
              .then(resp => {
                  const user = resp.data
                  commit('read_user_success', user)
                  resolve(resp)
              })
              .catch(err => {
                  reject(err)
              })
      })
    },
  },
  getters: {
    isLoggedIn: state => !!state.token && !!Cookies.get('token') && !!Cookies.get('claims'),
    authStatus: state => state.status,
    claims: state => state.claims || JSON.parse(Cookies.get('claims')),
    token: state => state.claims?.token || JSON.parse(Cookies.get('claims'))?.token,
    email: state => state.claims?.email || JSON.parse(Cookies.get('claims'))?.email,
    userScope: state => state.claims?.scope || JSON.parse(Cookies.get('claims'))?.userScope,
    user: state => state.user
  },
  modules: {
  }
})
