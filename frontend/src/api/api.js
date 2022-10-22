import axios from "axios";


const instance = axios.create({
  withCredentials: true,
  baseURL: 'http://192.168.137.15:4444/api/'
});

export const authAPI = {
  register(username, password) {
    return instance.post(`user`, {username, password});
  },
  login(username, password, max_age = 60 * 30) {
    return instance.post(`session`, {username, password,  max_age});
  },
  checkAuth() {
    return instance.get(`user`)
  },
  createGame(gamename, password, settings) {
    return instance.post('game/create', {gamename, password, settings})
  },
  loginGame(gamename, password) {
    return instance.post('game', {gamename, password})
  }
}

export const gameAPI = {
  getGame(gamename) {
    return instance.get(`game?gamename=${gamename}`)
  },
  move(gamename, from, to) {
    return instance.post('game/move', {gamename, from, to})
  }
}

window.authAPI = authAPI
window.gameAPI = gameAPI