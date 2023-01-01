/* @refresh reload */
import {render} from 'solid-js/web';

import App from './App';
import {HopeProvider} from '@hope-ui/solid'
import "./index.scss"

render(() => (
  <HopeProvider>
      <App/>
  </HopeProvider>
), document.getElementById('root') as HTMLElement);
