export interface FetchResponseType<T> {
    status: number,
    message: string,
    data?: T
}