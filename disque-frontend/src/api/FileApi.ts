import {handlerResult, httpClient} from "./axios";

export const fileList = () => handlerResult<{name: string, age: number}>(() => httpClient.get("/data"))
