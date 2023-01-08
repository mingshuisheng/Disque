import axios from "axios";
import type {AxiosResponse} from "axios";


export const baseURL =  process.env.NODE_ENV === "development" ? "/api" : "/"

export const coverURL = (url: string) => baseURL + url.startsWith("/")? baseURL + url: `${baseURL}/${url}`

export const httpClient = axios.create({
  baseURL
})

export const handlerResult = async <T>(request: () => Promise<AxiosResponse<T>>, showLoading: boolean = true): Promise<T> => {
  let response
  if (showLoading) {
    response = await handlerLoading(request)
  } else {
    response = await request();
  }
  if (response.status != 200) {
    throw new Error("server error");
  }
  return response.data
}

const handlerLoading = async <T>(request: () => Promise<AxiosResponse<T>>): Promise<AxiosResponse<T>> => {
  showLoading()
  let result = await request()
  hideLoading()
  return result
}

const showLoading = () => {

}

const hideLoading = () => {

}
