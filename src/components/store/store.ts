import { Commit, createStore } from 'vuex';

export default createStore({
  state: {
      authenticated: false,
      login: '',
    firstSection: {
      attr_Name: '',
      attr_Player: '',
      attr_Chronicle: '',
      attr_Nature: '',
      attr_Mask: '',
      attr_Energy: '',
      attr_Fraction: '',
      attr_Sect: '',
      attr_Concept: ''
    }
  },
  mutations: {
    UPDATE_FIELD(state, { fieldName, value }) {
      state.firstSection[fieldName] = value
    },
    UPDATE_FIRST_SECTION(state, newFirstSection) {
    Object.assign(state.firstSection, newFirstSection);
  },
    SET_AUTH: (state, auth) => state.authenticated = auth,
    SET_LOGIN: (state, login) => state.login = login,
  },
  actions: {
    setAuth: ({ commit }: { commit: Commit }, auth: boolean) => commit('SET_AUTH', auth),
    setLogin: ({ commit }, login) => commit('SET_LOGIN', login)
  },
  modules: {},  
});