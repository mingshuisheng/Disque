import axios, {AxiosResponse} from "axios";

export const httpClient = axios.create({
  baseURL: process.env.NODE_ENV === "development" ? "/api" : "/"
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
