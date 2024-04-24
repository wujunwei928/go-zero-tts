import axios from "axios";

const service = axios.create({
  baseURL: `${import.meta.env.VITE_BASE_PATH}${import.meta.env.VITE_BASE_API}`,
  withCredentials: true,  // 后端设置 Access-Control-Allow-Origin 为 * 时，前端不能设置 withCredentials
  timeout: 99999
})

export function request(config) {
  return new Promise((resolve, reject) => {
    service(config)
      .then((response) => {
        resolve(response)
      })
      .catch((error) => {
        reject(error)
      })
  })
}

export function get(url, params = {}, config = {}) {
  return request({
    url,
    method: 'get',
    params,
    ...config
  })
}

export function post(url, data = {}, config = {}) {
  return request({
    url,
    method: 'post',
    data,
    ...config
  })
}

export function postJson(url, data = {}, config = {}) {
  return post(url, data, {
    headers: {
      'Content-Type': 'application/json'
    },
    ...config
  })
}

export function postForm(url, data = {}, config = {}) {
  return post(url, data, {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    ...config
  })
}
