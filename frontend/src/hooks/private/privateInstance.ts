import Axios, { AxiosError, AxiosRequestConfig } from 'axios'

const baseURL = process.env.NEXT_PUBLIC_API_URL
export const AXIOS_INSTANCE = Axios.create({
  baseURL: baseURL,
  withCredentials: true, // enable cookies/credentials
})

// add a second `options` argument here if you want to pass extra options to each generated query
export const privateInstance = <T>(
  config: AxiosRequestConfig,
  options?: AxiosRequestConfig,
): Promise<T> => {
  const source = Axios.CancelToken.source()
  const promise = AXIOS_INSTANCE({
    ...config,
    ...options,
    cancelToken: source.token,
  }).then(({ data }) => data)

  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-expect-error
  promise.cancel = () => {
    source.cancel('Query was cancelled')
  }

  return promise
}

// In some case with react-query and swr you want to be able to override the return error type so you can also do it here like this
export type ErrorType<Error> = AxiosError<Error>
