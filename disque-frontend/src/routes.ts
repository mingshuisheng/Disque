import type {RouteDefinition} from "@solidjs/router";
import Home from "./pages/home/Home";
import {Navigator} from "@solidjs/router"

export const routes: RouteDefinition[] = [
  {
    path: "/",
    component: Home
  },
  {
    path: '/folder/:fileID',
    component: Home,
  }
]

export let globalNavigator: Navigator | undefined

export const setGlobalNavigator = (navigator: Navigator) => {
  globalNavigator = navigator
}


export const gotoFolder = (ID: number) => {
  globalNavigator?.(`/folder/${ID}`)
}
