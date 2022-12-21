/* @refresh reload */
import { render } from 'solid-js/web';
import {Router, hashIntegration} from "@solidjs/router";
import App from './App';
import "./index.scss"

render(() => (<Router source={hashIntegration()}><App /></Router>), document.getElementById('root') as HTMLElement);
